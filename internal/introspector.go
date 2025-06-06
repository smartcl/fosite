// Copyright © 2025 Ory Corp
// SPDX-License-Identifier: Apache-2.0

// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/ory/fosite (interfaces: TokenIntrospector)
//
// Generated by this command:
//
//	mockgen -package internal -destination internal/introspector.go github.com/ory/fosite TokenIntrospector
//

// Package internal is a generated GoMock package.
package internal

import (
	context "context"
	reflect "reflect"

	fosite "github.com/ory/fosite"
	gomock "go.uber.org/mock/gomock"
)

// MockTokenIntrospector is a mock of TokenIntrospector interface.
type MockTokenIntrospector struct {
	ctrl     *gomock.Controller
	recorder *MockTokenIntrospectorMockRecorder
	isgomock struct{}
}

// MockTokenIntrospectorMockRecorder is the mock recorder for MockTokenIntrospector.
type MockTokenIntrospectorMockRecorder struct {
	mock *MockTokenIntrospector
}

// NewMockTokenIntrospector creates a new mock instance.
func NewMockTokenIntrospector(ctrl *gomock.Controller) *MockTokenIntrospector {
	mock := &MockTokenIntrospector{ctrl: ctrl}
	mock.recorder = &MockTokenIntrospectorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTokenIntrospector) EXPECT() *MockTokenIntrospectorMockRecorder {
	return m.recorder
}

// IntrospectToken mocks base method.
func (m *MockTokenIntrospector) IntrospectToken(ctx context.Context, token string, tokenUse fosite.TokenType, accessRequest fosite.AccessRequester, scopes []string) (fosite.TokenType, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IntrospectToken", ctx, token, tokenUse, accessRequest, scopes)
	ret0, _ := ret[0].(fosite.TokenType)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IntrospectToken indicates an expected call of IntrospectToken.
func (mr *MockTokenIntrospectorMockRecorder) IntrospectToken(ctx, token, tokenUse, accessRequest, scopes any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IntrospectToken", reflect.TypeOf((*MockTokenIntrospector)(nil).IntrospectToken), ctx, token, tokenUse, accessRequest, scopes)
}
