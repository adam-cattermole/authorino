// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/common/auth.go

// Package mock_common is a generated GoMock package.
package mock_common

import (
	reflect "reflect"

	common "github.com/3scale-labs/authorino/pkg/common"
	envoy_service_auth_v3 "github.com/envoyproxy/go-control-plane/envoy/service/auth/v3"
	gomock "github.com/golang/mock/gomock"
	context "golang.org/x/net/context"
)

// MockAuthContext is a mock of AuthContext interface.
type MockAuthContext struct {
	ctrl     *gomock.Controller
	recorder *MockAuthContextMockRecorder
}

// MockAuthContextMockRecorder is the mock recorder for MockAuthContext.
type MockAuthContextMockRecorder struct {
	mock *MockAuthContext
}

// NewMockAuthContext creates a new mock instance.
func NewMockAuthContext(ctrl *gomock.Controller) *MockAuthContext {
	mock := &MockAuthContext{ctrl: ctrl}
	mock.recorder = &MockAuthContextMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAuthContext) EXPECT() *MockAuthContextMockRecorder {
	return m.recorder
}

// GetAPI mocks base method.
func (m *MockAuthContext) GetAPI() interface{} {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAPI")
	ret0, _ := ret[0].(interface{})
	return ret0
}

// GetAPI indicates an expected call of GetAPI.
func (mr *MockAuthContextMockRecorder) GetAPI() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAPI", reflect.TypeOf((*MockAuthContext)(nil).GetAPI))
}

// GetHttp mocks base method.
func (m *MockAuthContext) GetHttp() *envoy_service_auth_v3.AttributeContext_HttpRequest {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetHttp")
	ret0, _ := ret[0].(*envoy_service_auth_v3.AttributeContext_HttpRequest)
	return ret0
}

// GetHttp indicates an expected call of GetHttp.
func (mr *MockAuthContextMockRecorder) GetHttp() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHttp", reflect.TypeOf((*MockAuthContext)(nil).GetHttp))
}

// GetParentContext mocks base method.
func (m *MockAuthContext) GetParentContext() *context.Context {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetParentContext")
	ret0, _ := ret[0].(*context.Context)
	return ret0
}

// GetParentContext indicates an expected call of GetParentContext.
func (mr *MockAuthContextMockRecorder) GetParentContext() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetParentContext", reflect.TypeOf((*MockAuthContext)(nil).GetParentContext))
}

// GetRequest mocks base method.
func (m *MockAuthContext) GetRequest() *envoy_service_auth_v3.CheckRequest {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRequest")
	ret0, _ := ret[0].(*envoy_service_auth_v3.CheckRequest)
	return ret0
}

// GetRequest indicates an expected call of GetRequest.
func (mr *MockAuthContextMockRecorder) GetRequest() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRequest", reflect.TypeOf((*MockAuthContext)(nil).GetRequest))
}

// GetResolvedIdentity mocks base method.
func (m *MockAuthContext) GetResolvedIdentity() (interface{}, interface{}) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetResolvedIdentity")
	ret0, _ := ret[0].(interface{})
	ret1, _ := ret[1].(interface{})
	return ret0, ret1
}

// GetResolvedIdentity indicates an expected call of GetResolvedIdentity.
func (mr *MockAuthContextMockRecorder) GetResolvedIdentity() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetResolvedIdentity", reflect.TypeOf((*MockAuthContext)(nil).GetResolvedIdentity))
}

// GetResolvedMetadata mocks base method.
func (m *MockAuthContext) GetResolvedMetadata() map[interface{}]interface{} {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetResolvedMetadata")
	ret0, _ := ret[0].(map[interface{}]interface{})
	return ret0
}

// GetResolvedMetadata indicates an expected call of GetResolvedMetadata.
func (mr *MockAuthContextMockRecorder) GetResolvedMetadata() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetResolvedMetadata", reflect.TypeOf((*MockAuthContext)(nil).GetResolvedMetadata))
}

// MockAuthConfigEvaluator is a mock of AuthConfigEvaluator interface.
type MockAuthConfigEvaluator struct {
	ctrl     *gomock.Controller
	recorder *MockAuthConfigEvaluatorMockRecorder
}

// MockAuthConfigEvaluatorMockRecorder is the mock recorder for MockAuthConfigEvaluator.
type MockAuthConfigEvaluatorMockRecorder struct {
	mock *MockAuthConfigEvaluator
}

// NewMockAuthConfigEvaluator creates a new mock instance.
func NewMockAuthConfigEvaluator(ctrl *gomock.Controller) *MockAuthConfigEvaluator {
	mock := &MockAuthConfigEvaluator{ctrl: ctrl}
	mock.recorder = &MockAuthConfigEvaluatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAuthConfigEvaluator) EXPECT() *MockAuthConfigEvaluatorMockRecorder {
	return m.recorder
}

// Call mocks base method.
func (m *MockAuthConfigEvaluator) Call(arg0 common.AuthContext, arg1 context.Context) (interface{}, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Call", arg0, arg1)
	ret0, _ := ret[0].(interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Call indicates an expected call of Call.
func (mr *MockAuthConfigEvaluatorMockRecorder) Call(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Call", reflect.TypeOf((*MockAuthConfigEvaluator)(nil).Call), arg0, arg1)
}

// MockNamedConfigEvaluator is a mock of NamedConfigEvaluator interface.
type MockNamedConfigEvaluator struct {
	ctrl     *gomock.Controller
	recorder *MockNamedConfigEvaluatorMockRecorder
}

// MockNamedConfigEvaluatorMockRecorder is the mock recorder for MockNamedConfigEvaluator.
type MockNamedConfigEvaluatorMockRecorder struct {
	mock *MockNamedConfigEvaluator
}

// NewMockNamedConfigEvaluator creates a new mock instance.
func NewMockNamedConfigEvaluator(ctrl *gomock.Controller) *MockNamedConfigEvaluator {
	mock := &MockNamedConfigEvaluator{ctrl: ctrl}
	mock.recorder = &MockNamedConfigEvaluatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockNamedConfigEvaluator) EXPECT() *MockNamedConfigEvaluatorMockRecorder {
	return m.recorder
}

// GetName mocks base method.
func (m *MockNamedConfigEvaluator) GetName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetName")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetName indicates an expected call of GetName.
func (mr *MockNamedConfigEvaluatorMockRecorder) GetName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetName", reflect.TypeOf((*MockNamedConfigEvaluator)(nil).GetName))
}

// MockIdentityConfigEvaluator is a mock of IdentityConfigEvaluator interface.
type MockIdentityConfigEvaluator struct {
	ctrl     *gomock.Controller
	recorder *MockIdentityConfigEvaluatorMockRecorder
}

// MockIdentityConfigEvaluatorMockRecorder is the mock recorder for MockIdentityConfigEvaluator.
type MockIdentityConfigEvaluatorMockRecorder struct {
	mock *MockIdentityConfigEvaluator
}

// NewMockIdentityConfigEvaluator creates a new mock instance.
func NewMockIdentityConfigEvaluator(ctrl *gomock.Controller) *MockIdentityConfigEvaluator {
	mock := &MockIdentityConfigEvaluator{ctrl: ctrl}
	mock.recorder = &MockIdentityConfigEvaluatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIdentityConfigEvaluator) EXPECT() *MockIdentityConfigEvaluatorMockRecorder {
	return m.recorder
}

// GetOIDC mocks base method.
func (m *MockIdentityConfigEvaluator) GetOIDC() interface{} {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOIDC")
	ret0, _ := ret[0].(interface{})
	return ret0
}

// GetOIDC indicates an expected call of GetOIDC.
func (mr *MockIdentityConfigEvaluatorMockRecorder) GetOIDC() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOIDC", reflect.TypeOf((*MockIdentityConfigEvaluator)(nil).GetOIDC))
}
