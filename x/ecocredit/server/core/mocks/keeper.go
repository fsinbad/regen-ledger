// Code generated by MockGen. DO NOT EDIT.
// Source: x/ecocredit/server/core/keeper.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	types "github.com/cosmos/cosmos-sdk/types"
	types0 "github.com/cosmos/cosmos-sdk/x/params/types"
	gomock "github.com/golang/mock/gomock"
)

// MockParamKeeper is a mock of ParamKeeper interface.
type MockParamKeeper struct {
	ctrl     *gomock.Controller
	recorder *MockParamKeeperMockRecorder
}

// MockParamKeeperMockRecorder is the mock recorder for MockParamKeeper.
type MockParamKeeperMockRecorder struct {
	mock *MockParamKeeper
}

// NewMockParamKeeper creates a new mock instance.
func NewMockParamKeeper(ctrl *gomock.Controller) *MockParamKeeper {
	mock := &MockParamKeeper{ctrl: ctrl}
	mock.recorder = &MockParamKeeperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockParamKeeper) EXPECT() *MockParamKeeperMockRecorder {
	return m.recorder
}

// GetParamSet mocks base method.
func (m *MockParamKeeper) GetParamSet(ctx types.Context, ps types0.ParamSet) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "GetParamSet", ctx, ps)
}

// GetParamSet indicates an expected call of GetParamSet.
func (mr *MockParamKeeperMockRecorder) GetParamSet(ctx, ps interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetParamSet", reflect.TypeOf((*MockParamKeeper)(nil).GetParamSet), ctx, ps)
}