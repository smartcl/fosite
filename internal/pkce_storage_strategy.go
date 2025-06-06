// Copyright © 2025 Ory Corp
// SPDX-License-Identifier: Apache-2.0

// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/ory/fosite/handler/pkce (interfaces: PKCERequestStorage)
//
// Generated by this command:
//
//	mockgen -package internal -destination internal/pkce_storage_strategy.go github.com/ory/fosite/handler/pkce PKCERequestStorage
//

// Package internal is a generated GoMock package.
package internal

import (
	context "context"
	reflect "reflect"

	fosite "github.com/ory/fosite"
	gomock "go.uber.org/mock/gomock"
)

// MockPKCERequestStorage is a mock of PKCERequestStorage interface.
type MockPKCERequestStorage struct {
	ctrl     *gomock.Controller
	recorder *MockPKCERequestStorageMockRecorder
	isgomock struct{}
}

// MockPKCERequestStorageMockRecorder is the mock recorder for MockPKCERequestStorage.
type MockPKCERequestStorageMockRecorder struct {
	mock *MockPKCERequestStorage
}

// NewMockPKCERequestStorage creates a new mock instance.
func NewMockPKCERequestStorage(ctrl *gomock.Controller) *MockPKCERequestStorage {
	mock := &MockPKCERequestStorage{ctrl: ctrl}
	mock.recorder = &MockPKCERequestStorageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPKCERequestStorage) EXPECT() *MockPKCERequestStorageMockRecorder {
	return m.recorder
}

// CreatePKCERequestSession mocks base method.
func (m *MockPKCERequestStorage) CreatePKCERequestSession(ctx context.Context, signature string, requester fosite.Requester) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreatePKCERequestSession", ctx, signature, requester)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreatePKCERequestSession indicates an expected call of CreatePKCERequestSession.
func (mr *MockPKCERequestStorageMockRecorder) CreatePKCERequestSession(ctx, signature, requester any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatePKCERequestSession", reflect.TypeOf((*MockPKCERequestStorage)(nil).CreatePKCERequestSession), ctx, signature, requester)
}

// DeletePKCERequestSession mocks base method.
func (m *MockPKCERequestStorage) DeletePKCERequestSession(ctx context.Context, signature string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeletePKCERequestSession", ctx, signature)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeletePKCERequestSession indicates an expected call of DeletePKCERequestSession.
func (mr *MockPKCERequestStorageMockRecorder) DeletePKCERequestSession(ctx, signature any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeletePKCERequestSession", reflect.TypeOf((*MockPKCERequestStorage)(nil).DeletePKCERequestSession), ctx, signature)
}

// GetPKCERequestSession mocks base method.
func (m *MockPKCERequestStorage) GetPKCERequestSession(ctx context.Context, signature string, session fosite.Session) (fosite.Requester, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPKCERequestSession", ctx, signature, session)
	ret0, _ := ret[0].(fosite.Requester)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPKCERequestSession indicates an expected call of GetPKCERequestSession.
func (mr *MockPKCERequestStorageMockRecorder) GetPKCERequestSession(ctx, signature, session any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPKCERequestSession", reflect.TypeOf((*MockPKCERequestStorage)(nil).GetPKCERequestSession), ctx, signature, session)
}
