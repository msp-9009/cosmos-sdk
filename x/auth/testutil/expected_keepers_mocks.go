// Code generated by MockGen. DO NOT EDIT.
// Source: x/auth/types/expected_keepers.go
//
// Generated by this command:
//
//	mockgen -source=x/auth/types/expected_keepers.go -package testutil -destination x/auth/testutil/expected_keepers_mocks.go
//

// Package testutil is a generated GoMock package.
package testutil

import (
	context "context"
	reflect "reflect"

	transaction "cosmossdk.io/core/transaction"
	types "github.com/cosmos/cosmos-sdk/types"
	gomock "go.uber.org/mock/gomock"
)

// MockBankKeeper is a mock of BankKeeper interface.
type MockBankKeeper struct {
	ctrl     *gomock.Controller
	recorder *MockBankKeeperMockRecorder
	isgomock struct{}
}

// MockBankKeeperMockRecorder is the mock recorder for MockBankKeeper.
type MockBankKeeperMockRecorder struct {
	mock *MockBankKeeper
}

// NewMockBankKeeper creates a new mock instance.
func NewMockBankKeeper(ctrl *gomock.Controller) *MockBankKeeper {
	mock := &MockBankKeeper{ctrl: ctrl}
	mock.recorder = &MockBankKeeperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBankKeeper) EXPECT() *MockBankKeeperMockRecorder {
	return m.recorder
}

// IsSendEnabledCoins mocks base method.
func (m *MockBankKeeper) IsSendEnabledCoins(ctx context.Context, coins ...types.Coin) error {
	m.ctrl.T.Helper()
	varargs := []any{ctx}
	for _, a := range coins {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "IsSendEnabledCoins", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// IsSendEnabledCoins indicates an expected call of IsSendEnabledCoins.
func (mr *MockBankKeeperMockRecorder) IsSendEnabledCoins(ctx any, coins ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx}, coins...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsSendEnabledCoins", reflect.TypeOf((*MockBankKeeper)(nil).IsSendEnabledCoins), varargs...)
}

// MockSendCoinsFromAccountToModule mocks base method.
func (m *MockBankKeeper) MockSendCoinsFromAccountToModule(ctx context.Context, senderAddr types.AccAddress, recipientModule string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MockSendCoinsFromAccountToModule", ctx, senderAddr, recipientModule)
	ret0, _ := ret[0].(error)
	return ret0
}

// MockSendCoinsFromAccountToModule indicates an expected call of MockSendCoinsFromAccountToModule.
func (mr *MockBankKeeperMockRecorder) MockSendCoinsFromAccountToModule(ctx, senderAddr, recipientModule any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MockSendCoinsFromAccountToModule", reflect.TypeOf((*MockBankKeeper)(nil).MockSendCoinsFromAccountToModule), ctx, senderAddr, recipientModule)
}

// SendCoins mocks base method.
func (m *MockBankKeeper) SendCoins(ctx context.Context, from, to types.AccAddress, amt types.Coins) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendCoins", ctx, from, to, amt)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendCoins indicates an expected call of SendCoins.
func (mr *MockBankKeeperMockRecorder) SendCoins(ctx, from, to, amt any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendCoins", reflect.TypeOf((*MockBankKeeper)(nil).SendCoins), ctx, from, to, amt)
}

// SendCoinsFromAccountToModule mocks base method.
func (m *MockBankKeeper) SendCoinsFromAccountToModule(ctx context.Context, senderAddr types.AccAddress, recipientModule string, amt types.Coins) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendCoinsFromAccountToModule", ctx, senderAddr, recipientModule, amt)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendCoinsFromAccountToModule indicates an expected call of SendCoinsFromAccountToModule.
func (mr *MockBankKeeperMockRecorder) SendCoinsFromAccountToModule(ctx, senderAddr, recipientModule, amt any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendCoinsFromAccountToModule", reflect.TypeOf((*MockBankKeeper)(nil).SendCoinsFromAccountToModule), ctx, senderAddr, recipientModule, amt)
}

// MockAccountsModKeeper is a mock of AccountsModKeeper interface.
type MockAccountsModKeeper struct {
	ctrl     *gomock.Controller
	recorder *MockAccountsModKeeperMockRecorder
	isgomock struct{}
}

// MockAccountsModKeeperMockRecorder is the mock recorder for MockAccountsModKeeper.
type MockAccountsModKeeperMockRecorder struct {
	mock *MockAccountsModKeeper
}

// NewMockAccountsModKeeper creates a new mock instance.
func NewMockAccountsModKeeper(ctrl *gomock.Controller) *MockAccountsModKeeper {
	mock := &MockAccountsModKeeper{ctrl: ctrl}
	mock.recorder = &MockAccountsModKeeperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAccountsModKeeper) EXPECT() *MockAccountsModKeeperMockRecorder {
	return m.recorder
}

// InitAccountNumberSeqUnsafe mocks base method.
func (m *MockAccountsModKeeper) InitAccountNumberSeqUnsafe(ctx context.Context, currentAccNum uint64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InitAccountNumberSeqUnsafe", ctx, currentAccNum)
	ret0, _ := ret[0].(error)
	return ret0
}

// InitAccountNumberSeqUnsafe indicates an expected call of InitAccountNumberSeqUnsafe.
func (mr *MockAccountsModKeeperMockRecorder) InitAccountNumberSeqUnsafe(ctx, currentAccNum any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InitAccountNumberSeqUnsafe", reflect.TypeOf((*MockAccountsModKeeper)(nil).InitAccountNumberSeqUnsafe), ctx, currentAccNum)
}

// IsAccountsModuleAccount mocks base method.
func (m *MockAccountsModKeeper) IsAccountsModuleAccount(ctx context.Context, accountAddr []byte) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsAccountsModuleAccount", ctx, accountAddr)
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsAccountsModuleAccount indicates an expected call of IsAccountsModuleAccount.
func (mr *MockAccountsModKeeperMockRecorder) IsAccountsModuleAccount(ctx, accountAddr any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsAccountsModuleAccount", reflect.TypeOf((*MockAccountsModKeeper)(nil).IsAccountsModuleAccount), ctx, accountAddr)
}

// MigrateLegacyAccount mocks base method.
func (m *MockAccountsModKeeper) MigrateLegacyAccount(ctx context.Context, addr []byte, accNum uint64, accType string, msg transaction.Msg) (transaction.Msg, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MigrateLegacyAccount", ctx, addr, accNum, accType, msg)
	ret0, _ := ret[0].(transaction.Msg)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MigrateLegacyAccount indicates an expected call of MigrateLegacyAccount.
func (mr *MockAccountsModKeeperMockRecorder) MigrateLegacyAccount(ctx, addr, accNum, accType, msg any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MigrateLegacyAccount", reflect.TypeOf((*MockAccountsModKeeper)(nil).MigrateLegacyAccount), ctx, addr, accNum, accType, msg)
}

// NextAccountNumber mocks base method.
func (m *MockAccountsModKeeper) NextAccountNumber(ctx context.Context) (uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NextAccountNumber", ctx)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NextAccountNumber indicates an expected call of NextAccountNumber.
func (mr *MockAccountsModKeeperMockRecorder) NextAccountNumber(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NextAccountNumber", reflect.TypeOf((*MockAccountsModKeeper)(nil).NextAccountNumber), ctx)
}

// Query mocks base method.
func (m *MockAccountsModKeeper) Query(ctx context.Context, accountAddr []byte, queryRequest transaction.Msg) (transaction.Msg, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Query", ctx, accountAddr, queryRequest)
	ret0, _ := ret[0].(transaction.Msg)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Query indicates an expected call of Query.
func (mr *MockAccountsModKeeperMockRecorder) Query(ctx, accountAddr, queryRequest any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Query", reflect.TypeOf((*MockAccountsModKeeper)(nil).Query), ctx, accountAddr, queryRequest)
}

// SendModuleMessage mocks base method.
func (m *MockAccountsModKeeper) SendModuleMessage(ctx context.Context, sender []byte, msg transaction.Msg) (transaction.Msg, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendModuleMessage", ctx, sender, msg)
	ret0, _ := ret[0].(transaction.Msg)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SendModuleMessage indicates an expected call of SendModuleMessage.
func (mr *MockAccountsModKeeperMockRecorder) SendModuleMessage(ctx, sender, msg any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendModuleMessage", reflect.TypeOf((*MockAccountsModKeeper)(nil).SendModuleMessage), ctx, sender, msg)
}
