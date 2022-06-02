// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/cache/cache.go

// Package mock_cache is a generated GoMock package.
package mock_cache

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	evaluators "github.com/kuadrant/authorino/pkg/evaluators"
)

// MockCache is a mock of Cache interface.
type MockCache struct {
	ctrl     *gomock.Controller
	recorder *MockCacheMockRecorder
}

// MockCacheMockRecorder is the mock recorder for MockCache.
type MockCacheMockRecorder struct {
	mock *MockCache
}

// NewMockCache creates a new mock instance.
func NewMockCache(ctrl *gomock.Controller) *MockCache {
	mock := &MockCache{ctrl: ctrl}
	mock.recorder = &MockCacheMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCache) EXPECT() *MockCacheMockRecorder {
	return m.recorder
}

// Delete mocks base method.
func (m *MockCache) Delete(id string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Delete", id)
}

// Delete indicates an expected call of Delete.
func (mr *MockCacheMockRecorder) Delete(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockCache)(nil).Delete), id)
}

// FindId mocks base method.
func (m *MockCache) FindId(key string) (string, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindId", key)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// FindId indicates an expected call of FindId.
func (mr *MockCacheMockRecorder) FindId(key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindId", reflect.TypeOf((*MockCache)(nil).FindId), key)
}

// FindKeys mocks base method.
func (m *MockCache) FindKeys(id string) []string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindKeys", id)
	ret0, _ := ret[0].([]string)
	return ret0
}

// FindKeys indicates an expected call of FindKeys.
func (mr *MockCacheMockRecorder) FindKeys(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindKeys", reflect.TypeOf((*MockCache)(nil).FindKeys), id)
}

// Get mocks base method.
func (m *MockCache) Get(key string) *evaluators.AuthConfig {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", key)
	ret0, _ := ret[0].(*evaluators.AuthConfig)
	return ret0
}

// Get indicates an expected call of Get.
func (mr *MockCacheMockRecorder) Get(key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockCache)(nil).Get), key)
}

// List mocks base method.
func (m *MockCache) List() []*evaluators.AuthConfig {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List")
	ret0, _ := ret[0].([]*evaluators.AuthConfig)
	return ret0
}

// List indicates an expected call of List.
func (mr *MockCacheMockRecorder) List() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockCache)(nil).List))
}

// Set mocks base method.
func (m *MockCache) Set(id, key string, config evaluators.AuthConfig, override bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Set", id, key, config, override)
	ret0, _ := ret[0].(error)
	return ret0
}

// Set indicates an expected call of Set.
func (mr *MockCacheMockRecorder) Set(id, key, config, override interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Set", reflect.TypeOf((*MockCache)(nil).Set), id, key, config, override)
}
