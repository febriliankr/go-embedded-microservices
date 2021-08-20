// Code generated by MockGen. DO NOT EDIT.
// Source: app.go

// Package app is a generated GoMock package.
package app

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	dom "github.com/powerman/go-monolith-example/internal/dom"
)

// MockAppl is a mock of Appl interface.
type MockAppl struct {
	ctrl     *gomock.Controller
	recorder *MockApplMockRecorder
}

// MockApplMockRecorder is the mock recorder for MockAppl.
type MockApplMockRecorder struct {
	mock *MockAppl
}

// NewMockAppl creates a new mock instance.
func NewMockAppl(ctrl *gomock.Controller) *MockAppl {
	mock := &MockAppl{ctrl: ctrl}
	mock.recorder = &MockApplMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAppl) EXPECT() *MockApplMockRecorder {
	return m.recorder
}

// Example mocks base method.
func (m *MockAppl) Example(arg0 Ctx, arg1 dom.Auth, arg2 dom.UserName) (*Example, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Example", arg0, arg1, arg2)
	ret0, _ := ret[0].(*Example)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Example indicates an expected call of Example.
func (mr *MockApplMockRecorder) Example(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Example", reflect.TypeOf((*MockAppl)(nil).Example), arg0, arg1, arg2)
}

// IncExample mocks base method.
func (m *MockAppl) IncExample(arg0 Ctx, arg1 dom.Auth) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IncExample", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// IncExample indicates an expected call of IncExample.
func (mr *MockApplMockRecorder) IncExample(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IncExample", reflect.TypeOf((*MockAppl)(nil).IncExample), arg0, arg1)
}

// MockRepo is a mock of Repo interface.
type MockRepo struct {
	ctrl     *gomock.Controller
	recorder *MockRepoMockRecorder
}

// MockRepoMockRecorder is the mock recorder for MockRepo.
type MockRepoMockRecorder struct {
	mock *MockRepo
}

// NewMockRepo creates a new mock instance.
func NewMockRepo(ctrl *gomock.Controller) *MockRepo {
	mock := &MockRepo{ctrl: ctrl}
	mock.recorder = &MockRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepo) EXPECT() *MockRepoMockRecorder {
	return m.recorder
}

// Example mocks base method.
func (m *MockRepo) Example(arg0 Ctx, arg1 dom.UserName) (*Example, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Example", arg0, arg1)
	ret0, _ := ret[0].(*Example)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Example indicates an expected call of Example.
func (mr *MockRepoMockRecorder) Example(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Example", reflect.TypeOf((*MockRepo)(nil).Example), arg0, arg1)
}

// IncExample mocks base method.
func (m *MockRepo) IncExample(arg0 Ctx, arg1 dom.UserName) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IncExample", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// IncExample indicates an expected call of IncExample.
func (mr *MockRepoMockRecorder) IncExample(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IncExample", reflect.TypeOf((*MockRepo)(nil).IncExample), arg0, arg1)
}
