package keeper

import (
	"context"
	"fmt"

	"cosmossdk.io/x/evidence/types"

	"github.com/cosmos/cosmos-sdk/telemetry"
	consensusv1 "github.com/cosmos/cosmos-sdk/x/consensus/types"
)

// BeginBlocker iterates through and handles any newly discovered evidence of
// misbehavior submitted by CometBFT. Currently, only equivocation is handled.
func (k Keeper) BeginBlocker(ctx context.Context) error {
	defer telemetry.ModuleMeasureSince(types.ModuleName, telemetry.Now(), telemetry.MetricKeyBeginBlocker)
	if k.HeaderService.HeaderInfo(ctx).Height == 0 {
		return nil
	}

	// If the user is using the legacy CometBFT consensus, we need to convert the
	// evidence to the new types.Misbehavior type and store it in the cache.

	res := consensusv1.QueryCometInfoResponse{}
	if err := k.RouterService.QueryRouterService().InvokeTyped(ctx, &consensusv1.QueryCometInfoRequest{}, &res); err != nil {
		return err
	}

	for _, evidence := range res.CometInfo.Evidence {
		switch evidence.EvidenceType {
		// It's still ongoing discussion how should we treat and slash attacks with
		// premeditation. So for now we agree to treat them in the same way.
		case consensusv1.MisbehaviorType_MISBEHAVIOR_TYPE_LIGHT_CLIENT_ATTACK, consensusv1.MisbehaviorType_MISBEHAVIOR_TYPE_DUPLICATE_VOTE:
			if evidence == nil {
				continue // skip if no evidence
			}
			e := types.FromABCIEvidence(*evidence, k.stakingKeeper.ConsensusAddressCodec())
			err := k.handleEquivocationEvidence(ctx, e)
			if err != nil {
				return err
			}
		default:
			k.Logger.Error(fmt.Sprintf("ignored unknown evidence type: %x", evidence.EvidenceType))
		}
	}
	return nil
}
