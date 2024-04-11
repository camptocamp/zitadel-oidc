// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/zitadel/oidc/v3/pkg/op (interfaces: SigningKey,Key)

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	jose "github.com/go-jose/go-jose/v4"
	gomock "github.com/golang/mock/gomock"
)

// MockSigningKey is a mock of SigningKey interface.
type MockSigningKey struct {
	ctrl     *gomock.Controller
	recorder *MockSigningKeyMockRecorder
}

// MockSigningKeyMockRecorder is the mock recorder for MockSigningKey.
type MockSigningKeyMockRecorder struct {
	mock *MockSigningKey
}

// NewMockSigningKey creates a new mock instance.
func NewMockSigningKey(ctrl *gomock.Controller) *MockSigningKey {
	mock := &MockSigningKey{ctrl: ctrl}
	mock.recorder = &MockSigningKeyMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSigningKey) EXPECT() *MockSigningKeyMockRecorder {
	return m.recorder
}

// ID mocks base method.
func (m *MockSigningKey) ID() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ID")
	ret0, _ := ret[0].(string)
	return ret0
}

// ID indicates an expected call of ID.
func (mr *MockSigningKeyMockRecorder) ID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ID", reflect.TypeOf((*MockSigningKey)(nil).ID))
}

// Key mocks base method.
func (m *MockSigningKey) Key() interface{} {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Key")
	ret0, _ := ret[0].(interface{})
	return ret0
}

// Key indicates an expected call of Key.
func (mr *MockSigningKeyMockRecorder) Key() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Key", reflect.TypeOf((*MockSigningKey)(nil).Key))
}

// SignatureAlgorithm mocks base method.
func (m *MockSigningKey) SignatureAlgorithm() jose.SignatureAlgorithm {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SignatureAlgorithm")
	ret0, _ := ret[0].(jose.SignatureAlgorithm)
	return ret0
}

// SignatureAlgorithm indicates an expected call of SignatureAlgorithm.
func (mr *MockSigningKeyMockRecorder) SignatureAlgorithm() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SignatureAlgorithm", reflect.TypeOf((*MockSigningKey)(nil).SignatureAlgorithm))
}

// MockKey is a mock of Key interface.
type MockKey struct {
	ctrl     *gomock.Controller
	recorder *MockKeyMockRecorder
}

// MockKeyMockRecorder is the mock recorder for MockKey.
type MockKeyMockRecorder struct {
	mock *MockKey
}

// NewMockKey creates a new mock instance.
func NewMockKey(ctrl *gomock.Controller) *MockKey {
	mock := &MockKey{ctrl: ctrl}
	mock.recorder = &MockKeyMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockKey) EXPECT() *MockKeyMockRecorder {
	return m.recorder
}

// Algorithm mocks base method.
func (m *MockKey) Algorithm() jose.SignatureAlgorithm {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Algorithm")
	ret0, _ := ret[0].(jose.SignatureAlgorithm)
	return ret0
}

// Algorithm indicates an expected call of Algorithm.
func (mr *MockKeyMockRecorder) Algorithm() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Algorithm", reflect.TypeOf((*MockKey)(nil).Algorithm))
}

// ID mocks base method.
func (m *MockKey) ID() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ID")
	ret0, _ := ret[0].(string)
	return ret0
}

// ID indicates an expected call of ID.
func (mr *MockKeyMockRecorder) ID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ID", reflect.TypeOf((*MockKey)(nil).ID))
}

// Key mocks base method.
func (m *MockKey) Key() interface{} {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Key")
	ret0, _ := ret[0].(interface{})
	return ret0
}

// Key indicates an expected call of Key.
func (mr *MockKeyMockRecorder) Key() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Key", reflect.TypeOf((*MockKey)(nil).Key))
}

// Use mocks base method.
func (m *MockKey) Use() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Use")
	ret0, _ := ret[0].(string)
	return ret0
}

// Use indicates an expected call of Use.
func (mr *MockKeyMockRecorder) Use() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Use", reflect.TypeOf((*MockKey)(nil).Use))
}
