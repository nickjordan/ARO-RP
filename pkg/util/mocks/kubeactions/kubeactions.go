// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/Azure/ARO-RP/pkg/frontend/kubeactions (interfaces: Interface)

// Package mock_kubeactions is a generated GoMock package.
package mock_kubeactions

import (
	context "context"
	io "io"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	unstructured "k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"

	api "github.com/Azure/ARO-RP/pkg/api"
)

// MockInterface is a mock of Interface interface
type MockInterface struct {
	ctrl     *gomock.Controller
	recorder *MockInterfaceMockRecorder
}

// MockInterfaceMockRecorder is the mock recorder for MockInterface
type MockInterfaceMockRecorder struct {
	mock *MockInterface
}

// NewMockInterface creates a new mock instance
func NewMockInterface(ctrl *gomock.Controller) *MockInterface {
	mock := &MockInterface{ctrl: ctrl}
	mock.recorder = &MockInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockInterface) EXPECT() *MockInterfaceMockRecorder {
	return m.recorder
}

// CreateOrUpdate mocks base method
func (m *MockInterface) CreateOrUpdate(arg0 *api.OpenShiftCluster, arg1 *unstructured.Unstructured) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateOrUpdate", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateOrUpdate indicates an expected call of CreateOrUpdate
func (mr *MockInterfaceMockRecorder) CreateOrUpdate(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOrUpdate", reflect.TypeOf((*MockInterface)(nil).CreateOrUpdate), arg0, arg1)
}

// Delete mocks base method
func (m *MockInterface) Delete(arg0 *api.OpenShiftCluster, arg1, arg2, arg3 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *MockInterfaceMockRecorder) Delete(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockInterface)(nil).Delete), arg0, arg1, arg2, arg3)
}

// Get mocks base method
func (m *MockInterface) Get(arg0 *api.OpenShiftCluster, arg1, arg2, arg3 string) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockInterfaceMockRecorder) Get(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockInterface)(nil).Get), arg0, arg1, arg2, arg3)
}

// List mocks base method
func (m *MockInterface) List(arg0 *api.OpenShiftCluster, arg1, arg2 string) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0, arg1, arg2)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List
func (mr *MockInterfaceMockRecorder) List(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockInterface)(nil).List), arg0, arg1, arg2)
}

// MustGather mocks base method
func (m *MockInterface) MustGather(arg0 context.Context, arg1 *api.OpenShiftCluster, arg2 io.Writer) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MustGather", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// MustGather indicates an expected call of MustGather
func (mr *MockInterfaceMockRecorder) MustGather(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MustGather", reflect.TypeOf((*MockInterface)(nil).MustGather), arg0, arg1, arg2)
}
