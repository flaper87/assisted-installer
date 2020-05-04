// Code generated by MockGen. DO NOT EDIT.
// Source: k8s_client.go

// Package k8s_client is a generated GoMock package.
package k8s_client

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1 "k8s.io/api/core/v1"
)

// MockK8SClient is a mock of K8SClient interface.
type MockK8SClient struct {
	ctrl     *gomock.Controller
	recorder *MockK8SClientMockRecorder
}

// MockK8SClientMockRecorder is the mock recorder for MockK8SClient.
type MockK8SClientMockRecorder struct {
	mock *MockK8SClient
}

// NewMockK8SClient creates a new mock instance.
func NewMockK8SClient(ctrl *gomock.Controller) *MockK8SClient {
	mock := &MockK8SClient{ctrl: ctrl}
	mock.recorder = &MockK8SClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockK8SClient) EXPECT() *MockK8SClientMockRecorder {
	return m.recorder
}

// ListMasterNodes mocks base method.
func (m *MockK8SClient) ListMasterNodes() (*v1.NodeList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListMasterNodes")
	ret0, _ := ret[0].(*v1.NodeList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListMasterNodes indicates an expected call of ListMasterNodes.
func (mr *MockK8SClientMockRecorder) ListMasterNodes() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListMasterNodes", reflect.TypeOf((*MockK8SClient)(nil).ListMasterNodes))
}

// WaitForMasterNodes mocks base method.
func (m *MockK8SClient) WaitForMasterNodes(minMasterNodes int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WaitForMasterNodes", minMasterNodes)
	ret0, _ := ret[0].(error)
	return ret0
}

// WaitForMasterNodes indicates an expected call of WaitForMasterNodes.
func (mr *MockK8SClientMockRecorder) WaitForMasterNodes(minMasterNodes interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WaitForMasterNodes", reflect.TypeOf((*MockK8SClient)(nil).WaitForMasterNodes), minMasterNodes)
}

// PatchEtcd mocks base method.
func (m *MockK8SClient) PatchEtcd() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PatchEtcd")
	ret0, _ := ret[0].(error)
	return ret0
}

// PatchEtcd indicates an expected call of PatchEtcd.
func (mr *MockK8SClientMockRecorder) PatchEtcd() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PatchEtcd", reflect.TypeOf((*MockK8SClient)(nil).PatchEtcd))
}
