// Code generated by MockGen. DO NOT EDIT.
// Source: user.go

// Package mock_repository is a generated GoMock package.
package mock_repository

import (
	gomock "github.com/golang/mock/gomock"
	entity "github.com/kons16/team7-backend/domain/entity"
	reflect "reflect"
)

// MockUser is a mock of User interface
type MockUser struct {
	ctrl     *gomock.Controller
	recorder *MockUserMockRecorder
}

// MockUserMockRecorder is the mock recorder for MockUser
type MockUserMockRecorder struct {
	mock *MockUser
}

// NewMockUser creates a new mock instance
func NewMockUser(ctrl *gomock.Controller) *MockUser {
	mock := &MockUser{ctrl: ctrl}
	mock.recorder = &MockUserMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockUser) EXPECT() *MockUserMockRecorder {
	return m.recorder
}

// FindByID mocks base method
func (m *MockUser) FindByID(id string) (*entity.User, error) {
	//m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByID", id)
	ret0, _ := ret[0].(*entity.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByID indicates an expected call of FindByID
func (mr *MockUserMockRecorder) FindByID(id interface{}) *gomock.Call {
	//mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByID", reflect.TypeOf((*MockUser)(nil).FindByID), id)
}

// Create mocks base method
func (m *MockUser) Create(user *entity.User) (int, error) {
	//m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", user)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create
func (mr *MockUserMockRecorder) Create(user interface{}) *gomock.Call {
	//mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockUser)(nil).Create), user)
}

// FindUserBySubmitID mocks base method
func (m *MockUser) FindUserBySubmitID(submitID string) (*entity.LoginGetUser, error) {
	//m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindUserBySubmitID", submitID)
	ret0, _ := ret[0].(*entity.LoginGetUser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindUserBySubmitID indicates an expected call of FindUserBySubmitID
func (mr *MockUserMockRecorder) FindUserBySubmitID(submitID interface{}) *gomock.Call {
	//mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindUserBySubmitID", reflect.TypeOf((*MockUser)(nil).FindUserBySubmitID), submitID)
}