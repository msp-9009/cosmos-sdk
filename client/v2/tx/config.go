package tx

import (
	apitxsigning "cosmossdk.io/api/cosmos/tx/signing/v1beta1"
	"cosmossdk.io/client/v2/offchain"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	signingtypes "github.com/cosmos/cosmos-sdk/types/tx/signing"
)

// TxConfig defines an interface a client can utilize to generate an
// application-defined concrete transaction type. The type returned must
// implement TxBuilder.
type TxConfig interface {
	TxEncodingConfig
	TxSigningConfig
	TxBuilderProvider
}

// TxEncodingConfig defines an interface that contains transaction
// encoders and decoders
type TxEncodingConfig interface {
	TxEncoder() sdk.TxApiEncoder
	TxDecoder() sdk.TxApiDecoder
	TxJSONEncoder() sdk.TxApiEncoder
	TxJSONDecoder() sdk.TxApiDecoder
}

type TxSigningConfig interface {
	SignModeHandler() *offchain.HandlerMap
	SigningContext() *offchain.SignContext
	MarshalSignatureJSON([]signingtypes.SignatureV2) ([]byte, error)
	UnmarshalSignatureJSON([]byte) ([]signingtypes.SignatureV2, error)
}

type TxParameters struct {
	timeoutHeight uint64
	chainID       string
	memo          string
	signMode      apitxsigning.SignMode

	AccountConfig
	GasConfig
	FeeConfig
	ExecutionOptions
	ExtensionOptions
}

// AccountConfig defines the 'account' related fields in a transaction.
type AccountConfig struct {
	accountNumber uint64
	sequence      uint64
	fromName      string
	fromAddress   sdk.AccAddress
}

// GasConfig defines the 'gas' related fields in a transaction.
type GasConfig struct {
	gas           uint64
	gasAdjustment float64
	gasPrices     sdk.DecCoins
}

// FeeConfig defines the 'fee' related fields in a transaction.
type FeeConfig struct {
	fees       sdk.Coins
	feeGranter sdk.AccAddress
	feePayer   sdk.AccAddress
}

// ExecutionOptions defines the transaction execution options ran by the client
type ExecutionOptions struct {
	unordered          bool
	offline            bool
	generateOnly       bool
	simulateAndExecute bool
	preprocessTxHook   PreprocessTxFn
}

type ExtensionOptions struct {
	ExtOptions []*codectypes.Any
}
