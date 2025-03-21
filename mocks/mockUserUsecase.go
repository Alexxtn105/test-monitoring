// Code generated by MockGen. DO NOT EDIT.
// Source: test-monitoring/domain (interfaces: UserUseCase)

// Package mocks is a generated GoMock package.
package mocks

import (
	domain "test-monitoring/domain"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockUserUseCase is a mock of UserUseCase interface.
type MockUserUseCase struct {
	ctrl     *gomock.Controller
	recorder *MockUserUseCaseMockRecorder
}

// MockUserUseCaseMockRecorder is the mock recorder for MockUserUseCase.
type MockUserUseCaseMockRecorder struct {
	mock *MockUserUseCase
}

// NewMockUserUseCase creates a new mock instance.
func NewMockUserUseCase(ctrl *gomock.Controller) *MockUserUseCase {
	mock := &MockUserUseCase{ctrl: ctrl}
	mock.recorder = &MockUserUseCaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserUseCase) EXPECT() *MockUserUseCaseMockRecorder {
	return m.recorder
}

// CreateUser mocks base method.
func (m *MockUserUseCase) CreateUser(arg0 domain.User) (domain.User, *domain.AppError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", arg0)
	ret0, _ := ret[0].(domain.User)
	ret1, _ := ret[1].(*domain.AppError)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockUserUseCaseMockRecorder) CreateUser(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockUserUseCase)(nil).CreateUser), arg0)
}

// DeleteUserById mocks base method.
func (m *MockUserUseCase) DeleteUserById(arg0 uint) *domain.AppError {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteUserById", arg0)
	ret0, _ := ret[0].(*domain.AppError)
	return ret0
}

// DeleteUserById indicates an expected call of DeleteUserById.
func (mr *MockUserUseCaseMockRecorder) DeleteUserById(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUserById", reflect.TypeOf((*MockUserUseCase)(nil).DeleteUserById), arg0)
}

// GetUserById mocks base method.
func (m *MockUserUseCase) GetUserById(arg0 uint) (domain.User, *domain.AppError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserById", arg0)
	ret0, _ := ret[0].(domain.User)
	ret1, _ := ret[1].(*domain.AppError)
	return ret0, ret1
}

// GetUserById indicates an expected call of GetUserById.
func (mr *MockUserUseCaseMockRecorder) GetUserById(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserById", reflect.TypeOf((*MockUserUseCase)(nil).GetUserById), arg0)
}

// UpdateUser mocks base method.
func (m *MockUserUseCase) UpdateUser(arg0 domain.User) (domain.User, *domain.AppError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUser", arg0)
	ret0, _ := ret[0].(domain.User)
	ret1, _ := ret[1].(*domain.AppError)
	return ret0, ret1
}

// UpdateUser indicates an expected call of UpdateUser.
func (mr *MockUserUseCaseMockRecorder) UpdateUser(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUser", reflect.TypeOf((*MockUserUseCase)(nil).UpdateUser), arg0)
}