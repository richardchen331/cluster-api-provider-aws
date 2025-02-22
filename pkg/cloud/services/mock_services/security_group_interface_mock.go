/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by MockGen. DO NOT EDIT.
// Source: sigs.k8s.io/cluster-api-provider-aws/pkg/cloud/services (interfaces: SecurityGroupInterface)

// Package mock_services is a generated GoMock package.
package mock_services

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockSecurityGroupInterface is a mock of SecurityGroupInterface interface.
type MockSecurityGroupInterface struct {
	ctrl     *gomock.Controller
	recorder *MockSecurityGroupInterfaceMockRecorder
}

// MockSecurityGroupInterfaceMockRecorder is the mock recorder for MockSecurityGroupInterface.
type MockSecurityGroupInterfaceMockRecorder struct {
	mock *MockSecurityGroupInterface
}

// NewMockSecurityGroupInterface creates a new mock instance.
func NewMockSecurityGroupInterface(ctrl *gomock.Controller) *MockSecurityGroupInterface {
	mock := &MockSecurityGroupInterface{ctrl: ctrl}
	mock.recorder = &MockSecurityGroupInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSecurityGroupInterface) EXPECT() *MockSecurityGroupInterfaceMockRecorder {
	return m.recorder
}

// DeleteSecurityGroups mocks base method.
func (m *MockSecurityGroupInterface) DeleteSecurityGroups() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteSecurityGroups")
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteSecurityGroups indicates an expected call of DeleteSecurityGroups.
func (mr *MockSecurityGroupInterfaceMockRecorder) DeleteSecurityGroups() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteSecurityGroups", reflect.TypeOf((*MockSecurityGroupInterface)(nil).DeleteSecurityGroups))
}

// ReconcileSecurityGroups mocks base method.
func (m *MockSecurityGroupInterface) ReconcileSecurityGroups() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileSecurityGroups")
	ret0, _ := ret[0].(error)
	return ret0
}

// ReconcileSecurityGroups indicates an expected call of ReconcileSecurityGroups.
func (mr *MockSecurityGroupInterfaceMockRecorder) ReconcileSecurityGroups() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileSecurityGroups", reflect.TypeOf((*MockSecurityGroupInterface)(nil).ReconcileSecurityGroups))
}
