package root

import (
	"errors"
	"fmt"

	iavl_v2 "github.com/cosmos/iavl/v2"

	"cosmossdk.io/core/log"
	corestore "cosmossdk.io/core/store"
	"cosmossdk.io/store/v2"
	"cosmossdk.io/store/v2/commitment"
	"cosmossdk.io/store/v2/commitment/iavl"
	"cosmossdk.io/store/v2/commitment/iavlv2"
	"cosmossdk.io/store/v2/commitment/mem"
	"cosmossdk.io/store/v2/db"
	"cosmossdk.io/store/v2/internal"
	"cosmossdk.io/store/v2/metrics"
	"cosmossdk.io/store/v2/migration"
	"cosmossdk.io/store/v2/pruning"
	"cosmossdk.io/store/v2/snapshots"
)

type (
	SCType string
)

const (
	SCTypeIavl   SCType = "iavl"
	SCTypeIavlV2 SCType = "iavl-v2"
)

const storePrefixTpl = "s/k:%s/" // s/k:<storeKey>

// Options are the options for creating a root store.
type Options struct {
	SCType          SCType               `mapstructure:"sc-type" toml:"sc-type" comment:"State commitment database type. Currently we support: \"iavl\" and \"iavl-v2\""`
	SCPruningOption *store.PruningOption `mapstructure:"sc-pruning-option" toml:"sc-pruning-option" comment:"Pruning options for state commitment"`
	IavlConfig      *iavl.Config         `mapstructure:"iavl-config" toml:"iavl-config"`
	IavlV2Config    iavl_v2.TreeOptions
}

// FactoryOptions are the options for creating a root store.
type FactoryOptions struct {
	Logger    log.Logger
	RootDir   string
	Options   Options
	StoreKeys []string
	SCRawDB   corestore.KVStoreWithBatch
}

// DefaultStoreOptions returns the default options for creating a root store.
func DefaultStoreOptions() Options {
	return Options{
		SCType: SCTypeIavl,
		SCPruningOption: &store.PruningOption{
			KeepRecent: 2,
			Interval:   100,
		},
		IavlConfig: &iavl.Config{
			CacheSize:              500_000,
			SkipFastStorageUpgrade: true,
		},
	}
}

// CreateRootStore is a convenience function to create a root store based on the
// provided FactoryOptions. Strictly speaking app developers can create the root
// store directly by calling root.New, so this function is not
// necessary, but demonstrates the required steps and configuration to create a root store.
func CreateRootStore(opts *FactoryOptions) (store.RootStore, error) {
	var (
		sc  *commitment.CommitStore
		err error
	)

	storeOpts := opts.Options

	metadata := commitment.NewMetadataStore(opts.SCRawDB)
	latestVersion, err := metadata.GetLatestVersion()
	if err != nil {
		return nil, err
	}
	if len(opts.StoreKeys) == 0 {
		lastCommitInfo, err := metadata.GetCommitInfo(latestVersion)
		if err != nil {
			return nil, err
		}
		if lastCommitInfo == nil {
			return nil, fmt.Errorf("tried to construct a root store with no store keys specified but no commit info found for version %d", latestVersion)
		}
		for _, si := range lastCommitInfo.StoreInfos {
			opts.StoreKeys = append(opts.StoreKeys, string(si.Name))
		}
	}
	removedStoreKeys, err := metadata.GetRemovedStoreKeys(latestVersion)
	if err != nil {
		return nil, err
	}

	newTreeFn := func(key string, scType SCType) (commitment.Tree, error) {
		if internal.IsMemoryStoreKey(key) {
			return mem.New(), nil
		} else {
			switch scType {
			case SCTypeIavl:
				return iavl.NewIavlTree(db.NewPrefixDB(opts.SCRawDB, []byte(fmt.Sprintf(storePrefixTpl, key))), opts.Logger, storeOpts.IavlConfig), nil
			case SCTypeIavlV2:
				dir := fmt.Sprintf("%s/data/iavl-v2/%s", opts.RootDir, key)
				return iavlv2.NewTree(opts.Options.IavlV2Config, iavl_v2.SqliteDbOptions{Path: dir}, opts.Logger)
			default:
				return nil, errors.New("unsupported commitment store type")
			}
		}
	}

	// check if we need to migrate the store
	isMigrating := false
	scType := storeOpts.SCType

	if scType != SCTypeIavl {
		isMigrating = true  // need to migrate
		scType = SCTypeIavl // only support iavl v1 for migration
	}

	trees := make(map[string]commitment.Tree, len(opts.StoreKeys))
	for _, key := range opts.StoreKeys {
		tree, err := newTreeFn(key, scType)
		if err != nil {
			return nil, err
		}
		if isMigrating {
			v, err := tree.GetLatestVersion()
			if err != nil {
				return nil, err
			}
			if v == 0 && latestVersion > 0 {
				if err := tree.SetInitialVersion(latestVersion + 1); err != nil {
					return nil, err
				}
			}
		}
		trees[key] = tree
	}
	oldTrees := make(map[string]commitment.Tree, len(opts.StoreKeys))
	for _, key := range removedStoreKeys {
		tree, err := newTreeFn(string(key), scType)
		if err != nil {
			return nil, err
		}
		oldTrees[string(key)] = tree
	}

	sc, err = commitment.NewCommitStore(trees, oldTrees, opts.SCRawDB, opts.Logger)
	if err != nil {
		return nil, err
	}

	var mm *migration.Manager
	if isMigrating {
		snapshotDB, err := snapshots.NewStore(fmt.Sprintf("%s/data/snapshots/store.db", opts.RootDir))
		if err != nil {
			return nil, err
		}
		snapshotMgr := snapshots.NewManager(snapshotDB, snapshots.SnapshotOptions{}, sc, nil, opts.Logger)
		var newSC *commitment.CommitStore
		if scType != storeOpts.SCType {
			newTrees := make(map[string]commitment.Tree, len(opts.StoreKeys))
			for _, key := range opts.StoreKeys {
				tree, err := newTreeFn(key, storeOpts.SCType)
				if err != nil {
					return nil, err
				}
				newTrees[key] = tree
			}
			newSC, err = commitment.NewCommitStore(newTrees, nil, opts.SCRawDB, opts.Logger)
			if err != nil {
				return nil, err
			}
		}
		mm = migration.NewManager(opts.SCRawDB, snapshotMgr, newSC, opts.Logger)
	}

	pm := pruning.NewManager(sc, storeOpts.SCPruningOption)
	return New(opts.SCRawDB, opts.Logger, sc, pm, mm, metrics.NoOpMetrics{})
}
