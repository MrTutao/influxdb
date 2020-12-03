// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/influxdata/influxdb/v2/v1/authorization (interfaces: AuthFinder)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	influxdb "github.com/influxdata/influxdb/v2"
	reflect "reflect"
)

// MockAuthFinder is a mock of AuthFinder interface
type MockAuthFinder struct {
	ctrl     *gomock.Controller
	recorder *MockAuthFinderMockRecorder
}

// MockAuthFinderMockRecorder is the mock recorder for MockAuthFinder
type MockAuthFinderMockRecorder struct {
	mock *MockAuthFinder
}

// NewMockAuthFinder creates a new mock instance
func NewMockAuthFinder(ctrl *gomock.Controller) *MockAuthFinder {
	mock := &MockAuthFinder{ctrl: ctrl}
	mock.recorder = &MockAuthFinderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAuthFinder) EXPECT() *MockAuthFinderMockRecorder {
	return m.recorder
}

// FindAuthorizationByID mocks base method
func (m *MockAuthFinder) FindAuthorizationByID(arg0 context.Context, arg1 influxdb.ID) (*influxdb.Authorization, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindAuthorizationByID", arg0, arg1)
	ret0, _ := ret[0].(*influxdb.Authorization)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindAuthorizationByID indicates an expected call of FindAuthorizationByID
func (mr *MockAuthFinderMockRecorder) FindAuthorizationByID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAuthorizationByID", reflect.TypeOf((*MockAuthFinder)(nil).FindAuthorizationByID), arg0, arg1)
}
