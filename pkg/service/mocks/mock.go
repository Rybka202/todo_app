// Code generated by MockGen. DO NOT EDIT.
// Source: service.go
//
// Generated by this command:
//
//	mockgen -source=service.go -destination=mocks/mock.go
//

// Package mock_service is a generated GoMock package.
package mock_service

import (
	reflect "reflect"
	todo_app "todo_app"

	gomock "go.uber.org/mock/gomock"
)

// MockAuthorization is a mock of Authorization interface.
type MockAuthorization struct {
	ctrl     *gomock.Controller
	recorder *MockAuthorizationMockRecorder
}

// MockAuthorizationMockRecorder is the mock recorder for MockAuthorization.
type MockAuthorizationMockRecorder struct {
	mock *MockAuthorization
}

// NewMockAuthorization creates a new mock instance.
func NewMockAuthorization(ctrl *gomock.Controller) *MockAuthorization {
	mock := &MockAuthorization{ctrl: ctrl}
	mock.recorder = &MockAuthorizationMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAuthorization) EXPECT() *MockAuthorizationMockRecorder {
	return m.recorder
}

// CreateUser mocks base method.
func (m *MockAuthorization) CreateUser(user todo_app.User) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", user)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockAuthorizationMockRecorder) CreateUser(user any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockAuthorization)(nil).CreateUser), user)
}

// GenerateToken mocks base method.
func (m *MockAuthorization) GenerateToken(username, password string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateToken", username, password)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GenerateToken indicates an expected call of GenerateToken.
func (mr *MockAuthorizationMockRecorder) GenerateToken(username, password any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateToken", reflect.TypeOf((*MockAuthorization)(nil).GenerateToken), username, password)
}

// ParseToken mocks base method.
func (m *MockAuthorization) ParseToken(token string) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ParseToken", token)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ParseToken indicates an expected call of ParseToken.
func (mr *MockAuthorizationMockRecorder) ParseToken(token any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ParseToken", reflect.TypeOf((*MockAuthorization)(nil).ParseToken), token)
}

// MockTodoList is a mock of TodoList interface.
type MockTodoList struct {
	ctrl     *gomock.Controller
	recorder *MockTodoListMockRecorder
}

// MockTodoListMockRecorder is the mock recorder for MockTodoList.
type MockTodoListMockRecorder struct {
	mock *MockTodoList
}

// NewMockTodoList creates a new mock instance.
func NewMockTodoList(ctrl *gomock.Controller) *MockTodoList {
	mock := &MockTodoList{ctrl: ctrl}
	mock.recorder = &MockTodoListMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTodoList) EXPECT() *MockTodoListMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockTodoList) Create(userID int, list todo_app.TodoList) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", userID, list)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockTodoListMockRecorder) Create(userID, list any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockTodoList)(nil).Create), userID, list)
}

// Delete mocks base method.
func (m *MockTodoList) Delete(userID, listID int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", userID, listID)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockTodoListMockRecorder) Delete(userID, listID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockTodoList)(nil).Delete), userID, listID)
}

// GetAll mocks base method.
func (m *MockTodoList) GetAll(userID int) ([]todo_app.TodoList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll", userID)
	ret0, _ := ret[0].([]todo_app.TodoList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll.
func (mr *MockTodoListMockRecorder) GetAll(userID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockTodoList)(nil).GetAll), userID)
}

// GetById mocks base method.
func (m *MockTodoList) GetById(userID, listID int) (todo_app.TodoList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetById", userID, listID)
	ret0, _ := ret[0].(todo_app.TodoList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetById indicates an expected call of GetById.
func (mr *MockTodoListMockRecorder) GetById(userID, listID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetById", reflect.TypeOf((*MockTodoList)(nil).GetById), userID, listID)
}

// Update mocks base method.
func (m *MockTodoList) Update(userID, listID int, input todo_app.UpdateListInput) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", userID, listID, input)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockTodoListMockRecorder) Update(userID, listID, input any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockTodoList)(nil).Update), userID, listID, input)
}

// MockTodoItem is a mock of TodoItem interface.
type MockTodoItem struct {
	ctrl     *gomock.Controller
	recorder *MockTodoItemMockRecorder
}

// MockTodoItemMockRecorder is the mock recorder for MockTodoItem.
type MockTodoItemMockRecorder struct {
	mock *MockTodoItem
}

// NewMockTodoItem creates a new mock instance.
func NewMockTodoItem(ctrl *gomock.Controller) *MockTodoItem {
	mock := &MockTodoItem{ctrl: ctrl}
	mock.recorder = &MockTodoItemMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTodoItem) EXPECT() *MockTodoItemMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockTodoItem) Create(userID, listID int, item todo_app.TodoItem) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", userID, listID, item)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockTodoItemMockRecorder) Create(userID, listID, item any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockTodoItem)(nil).Create), userID, listID, item)
}

// Delete mocks base method.
func (m *MockTodoItem) Delete(userID, itemID int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", userID, itemID)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockTodoItemMockRecorder) Delete(userID, itemID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockTodoItem)(nil).Delete), userID, itemID)
}

// GetAll mocks base method.
func (m *MockTodoItem) GetAll(userID, listID int) ([]todo_app.TodoItem, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll", userID, listID)
	ret0, _ := ret[0].([]todo_app.TodoItem)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll.
func (mr *MockTodoItemMockRecorder) GetAll(userID, listID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockTodoItem)(nil).GetAll), userID, listID)
}

// GetById mocks base method.
func (m *MockTodoItem) GetById(userID, itemID int) (todo_app.TodoItem, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetById", userID, itemID)
	ret0, _ := ret[0].(todo_app.TodoItem)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetById indicates an expected call of GetById.
func (mr *MockTodoItemMockRecorder) GetById(userID, itemID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetById", reflect.TypeOf((*MockTodoItem)(nil).GetById), userID, itemID)
}

// Update mocks base method.
func (m *MockTodoItem) Update(userID, itemID int, input todo_app.UpdateItemInput) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", userID, itemID, input)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockTodoItemMockRecorder) Update(userID, itemID, input any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockTodoItem)(nil).Update), userID, itemID, input)
}
