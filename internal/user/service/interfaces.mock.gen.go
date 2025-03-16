// Code generated by MockGen. DO NOT EDIT.
// Source: internal/user/service/interfaces.go

// Package service is a generated GoMock package.
package service

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockUserServiceInterface is a mock of UserServiceInterface interface.
type MockUserServiceInterface struct {
	ctrl     *gomock.Controller
	recorder *MockUserServiceInterfaceMockRecorder
}

// MockUserServiceInterfaceMockRecorder is the mock recorder for MockUserServiceInterface.
type MockUserServiceInterfaceMockRecorder struct {
	mock *MockUserServiceInterface
}

// NewMockUserServiceInterface creates a new mock instance.
func NewMockUserServiceInterface(ctrl *gomock.Controller) *MockUserServiceInterface {
	mock := &MockUserServiceInterface{ctrl: ctrl}
	mock.recorder = &MockUserServiceInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserServiceInterface) EXPECT() *MockUserServiceInterfaceMockRecorder {
	return m.recorder
}

// CreateEmployee mocks base method.
func (m *MockUserServiceInterface) CreateEmployee(arg0 CreateEmployeeRequest) (CreateEmployeeResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateEmployee", arg0)
	ret0, _ := ret[0].(CreateEmployeeResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateEmployee indicates an expected call of CreateEmployee.
func (mr *MockUserServiceInterfaceMockRecorder) CreateEmployee(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateEmployee", reflect.TypeOf((*MockUserServiceInterface)(nil).CreateEmployee), arg0)
}
