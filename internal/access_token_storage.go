// Copyright © 2025 Ory Corp
// SPDX-License-Identifier: Apache-2.0

// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/ory/fosite/handler/oauth2 (interfaces: AccessTokenStorage)
//
// Generated by this command:
//
//	mockgen -package internal -destination internal/access_token_storage.go github.com/ory/fosite/handler/oauth2 AccessTokenStorage
//

// Package internal is a generated GoMock package.
package internal

import (
	context "context"
	reflect "reflect"

	fosite "github.com/ory/fosite"
	gomock "go.uber.org/mock/gomock"
)

// MockAccessTokenStorage is a mock of AccessTokenStorage interface.
type MockAccessTokenStorage struct {
	ctrl     *gomock.Controller
	recorder *MockAccessTokenStorageMockRecorder
	isgomock struct{}
}

// MockAccessTokenStorageMockRecorder is the mock recorder for MockAccessTokenStorage.
type MockAccessTokenStorageMockRecorder struct {
	mock *MockAccessTokenStorage
}

// NewMockAccessTokenStorage creates a new mock instance.
func NewMockAccessTokenStorage(ctrl *gomock.Controller) *MockAccessTokenStorage {
	mock := &MockAccessTokenStorage{ctrl: ctrl}
	mock.recorder = &MockAccessTokenStorageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAccessTokenStorage) EXPECT() *MockAccessTokenStorageMockRecorder {
	return m.recorder
}

// CreateAccessTokenSession mocks base method.
func (m *MockAccessTokenStorage) CreateAccessTokenSession(ctx context.Context, signature string, request fosite.Requester) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateAccessTokenSession", ctx, signature, request)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateAccessTokenSession indicates an expected call of CreateAccessTokenSession.
func (mr *MockAccessTokenStorageMockRecorder) CreateAccessTokenSession(ctx, signature, request any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAccessTokenSession", reflect.TypeOf((*MockAccessTokenStorage)(nil).CreateAccessTokenSession), ctx, signature, request)
}

// DeleteAccessTokenSession mocks base method.
func (m *MockAccessTokenStorage) DeleteAccessTokenSession(ctx context.Context, signature string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAccessTokenSession", ctx, signature)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteAccessTokenSession indicates an expected call of DeleteAccessTokenSession.
func (mr *MockAccessTokenStorageMockRecorder) DeleteAccessTokenSession(ctx, signature any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAccessTokenSession", reflect.TypeOf((*MockAccessTokenStorage)(nil).DeleteAccessTokenSession), ctx, signature)
}

// GetAccessTokenSession mocks base method.
func (m *MockAccessTokenStorage) GetAccessTokenSession(ctx context.Context, signature string, session fosite.Session) (fosite.Requester, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAccessTokenSession", ctx, signature, session)
	ret0, _ := ret[0].(fosite.Requester)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAccessTokenSession indicates an expected call of GetAccessTokenSession.
func (mr *MockAccessTokenStorageMockRecorder) GetAccessTokenSession(ctx, signature, session any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccessTokenSession", reflect.TypeOf((*MockAccessTokenStorage)(nil).GetAccessTokenSession), ctx, signature, session)
}
