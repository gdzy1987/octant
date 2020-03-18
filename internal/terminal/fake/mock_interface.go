// Code generated by MockGen. DO NOT EDIT.
// Source: manager.go

// Package fake is a generated GoMock package.
package fake

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	terminal "github.com/vmware-tanzu/octant/internal/terminal"
	log "github.com/vmware-tanzu/octant/pkg/log"
	store "github.com/vmware-tanzu/octant/pkg/store"
	reflect "reflect"
)

// MockManager is a mock of Manager interface
type MockManager struct {
	ctrl     *gomock.Controller
	recorder *MockManagerMockRecorder
}

// MockManagerMockRecorder is the mock recorder for MockManager
type MockManagerMockRecorder struct {
	mock *MockManager
}

// NewMockManager creates a new mock instance
func NewMockManager(ctrl *gomock.Controller) *MockManager {
	mock := &MockManager{ctrl: ctrl}
	mock.recorder = &MockManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockManager) EXPECT() *MockManagerMockRecorder {
	return m.recorder
}

// List mocks base method
func (m *MockManager) List(namespace string) []terminal.Instance {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", namespace)
	ret0, _ := ret[0].([]terminal.Instance)
	return ret0
}

// List indicates an expected call of List
func (mr *MockManagerMockRecorder) List(namespace interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockManager)(nil).List), namespace)
}

// Get mocks base method
func (m *MockManager) Get(id string) (terminal.Instance, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", id)
	ret0, _ := ret[0].(terminal.Instance)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockManagerMockRecorder) Get(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockManager)(nil).Get), id)
}

// Delete mocks base method
func (m *MockManager) Delete(id string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Delete", id)
}

// Delete indicates an expected call of Delete
func (mr *MockManagerMockRecorder) Delete(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockManager)(nil).Delete), id)
}

// Create mocks base method
func (m *MockManager) Create(ctx context.Context, logger log.Logger, key store.Key, container, command, sessionID string) (terminal.Instance, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, logger, key, container, command, sessionID)
	ret0, _ := ret[0].(terminal.Instance)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create
func (mr *MockManagerMockRecorder) Create(ctx, logger, key, container, command, sessionID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockManager)(nil).Create), ctx, logger, key, container, command, sessionID)
}

// Select mocks base method
func (m *MockManager) Select(ctx context.Context) chan terminal.Instance {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Select", ctx)
	ret0, _ := ret[0].(chan terminal.Instance)
	return ret0
}

// Select indicates an expected call of Select
func (mr *MockManagerMockRecorder) Select(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Select", reflect.TypeOf((*MockManager)(nil).Select), ctx)
}

// ForceUpdate mocks base method
func (m *MockManager) ForceUpdate(id string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ForceUpdate", id)
}

// ForceUpdate indicates an expected call of ForceUpdate
func (mr *MockManagerMockRecorder) ForceUpdate(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ForceUpdate", reflect.TypeOf((*MockManager)(nil).ForceUpdate), id)
}

// SendScrollback mocks base method
func (m *MockManager) SendScrollback(id string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendScrollback", id)
	ret0, _ := ret[0].(bool)
	return ret0
}

// SendScrollback indicates an expected call of SendScrollback
func (mr *MockManagerMockRecorder) SendScrollback(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendScrollback", reflect.TypeOf((*MockManager)(nil).SendScrollback), id)
}

// SetScrollback mocks base method
func (m *MockManager) SetScrollback(id string, send bool) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetScrollback", id, send)
}

// SetScrollback indicates an expected call of SetScrollback
func (mr *MockManagerMockRecorder) SetScrollback(id, send interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetScrollback", reflect.TypeOf((*MockManager)(nil).SetScrollback), id, send)
}

// StopAll mocks base method
func (m *MockManager) StopAll() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StopAll")
	ret0, _ := ret[0].(error)
	return ret0
}

// StopAll indicates an expected call of StopAll
func (mr *MockManagerMockRecorder) StopAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StopAll", reflect.TypeOf((*MockManager)(nil).StopAll))
}
