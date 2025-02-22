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
// Source: sigs.k8s.io/cluster-api-provider-aws/pkg/cloud/services (interfaces: EC2MachineInterface)

// Package mock_services is a generated GoMock package.
package mock_services

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1beta1 "sigs.k8s.io/cluster-api-provider-aws/api/v1beta1"
	v1beta10 "sigs.k8s.io/cluster-api-provider-aws/exp/api/v1beta1"
	scope "sigs.k8s.io/cluster-api-provider-aws/pkg/cloud/scope"
)

// MockEC2MachineInterface is a mock of EC2MachineInterface interface.
type MockEC2MachineInterface struct {
	ctrl     *gomock.Controller
	recorder *MockEC2MachineInterfaceMockRecorder
}

// MockEC2MachineInterfaceMockRecorder is the mock recorder for MockEC2MachineInterface.
type MockEC2MachineInterfaceMockRecorder struct {
	mock *MockEC2MachineInterface
}

// NewMockEC2MachineInterface creates a new mock instance.
func NewMockEC2MachineInterface(ctrl *gomock.Controller) *MockEC2MachineInterface {
	mock := &MockEC2MachineInterface{ctrl: ctrl}
	mock.recorder = &MockEC2MachineInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEC2MachineInterface) EXPECT() *MockEC2MachineInterfaceMockRecorder {
	return m.recorder
}

// CreateInstance mocks base method.
func (m *MockEC2MachineInterface) CreateInstance(arg0 *scope.MachineScope, arg1 []byte) (*v1beta1.Instance, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateInstance", arg0, arg1)
	ret0, _ := ret[0].(*v1beta1.Instance)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateInstance indicates an expected call of CreateInstance.
func (mr *MockEC2MachineInterfaceMockRecorder) CreateInstance(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateInstance", reflect.TypeOf((*MockEC2MachineInterface)(nil).CreateInstance), arg0, arg1)
}

// CreateLaunchTemplate mocks base method.
func (m *MockEC2MachineInterface) CreateLaunchTemplate(arg0 *scope.MachinePoolScope, arg1 *string, arg2 []byte) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateLaunchTemplate", arg0, arg1, arg2)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateLaunchTemplate indicates an expected call of CreateLaunchTemplate.
func (mr *MockEC2MachineInterfaceMockRecorder) CreateLaunchTemplate(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateLaunchTemplate", reflect.TypeOf((*MockEC2MachineInterface)(nil).CreateLaunchTemplate), arg0, arg1, arg2)
}

// CreateLaunchTemplateVersion mocks base method.
func (m *MockEC2MachineInterface) CreateLaunchTemplateVersion(arg0 *scope.MachinePoolScope, arg1 *string, arg2 []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateLaunchTemplateVersion", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateLaunchTemplateVersion indicates an expected call of CreateLaunchTemplateVersion.
func (mr *MockEC2MachineInterfaceMockRecorder) CreateLaunchTemplateVersion(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateLaunchTemplateVersion", reflect.TypeOf((*MockEC2MachineInterface)(nil).CreateLaunchTemplateVersion), arg0, arg1, arg2)
}

// DeleteBastion mocks base method.
func (m *MockEC2MachineInterface) DeleteBastion() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteBastion")
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteBastion indicates an expected call of DeleteBastion.
func (mr *MockEC2MachineInterfaceMockRecorder) DeleteBastion() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteBastion", reflect.TypeOf((*MockEC2MachineInterface)(nil).DeleteBastion))
}

// DeleteLaunchTemplate mocks base method.
func (m *MockEC2MachineInterface) DeleteLaunchTemplate(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteLaunchTemplate", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteLaunchTemplate indicates an expected call of DeleteLaunchTemplate.
func (mr *MockEC2MachineInterfaceMockRecorder) DeleteLaunchTemplate(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteLaunchTemplate", reflect.TypeOf((*MockEC2MachineInterface)(nil).DeleteLaunchTemplate), arg0)
}

// DetachSecurityGroupsFromNetworkInterface mocks base method.
func (m *MockEC2MachineInterface) DetachSecurityGroupsFromNetworkInterface(arg0 []string, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DetachSecurityGroupsFromNetworkInterface", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DetachSecurityGroupsFromNetworkInterface indicates an expected call of DetachSecurityGroupsFromNetworkInterface.
func (mr *MockEC2MachineInterfaceMockRecorder) DetachSecurityGroupsFromNetworkInterface(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DetachSecurityGroupsFromNetworkInterface", reflect.TypeOf((*MockEC2MachineInterface)(nil).DetachSecurityGroupsFromNetworkInterface), arg0, arg1)
}

// DiscoverLaunchTemplateAMI mocks base method.
func (m *MockEC2MachineInterface) DiscoverLaunchTemplateAMI(arg0 *scope.MachinePoolScope) (*string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DiscoverLaunchTemplateAMI", arg0)
	ret0, _ := ret[0].(*string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DiscoverLaunchTemplateAMI indicates an expected call of DiscoverLaunchTemplateAMI.
func (mr *MockEC2MachineInterfaceMockRecorder) DiscoverLaunchTemplateAMI(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DiscoverLaunchTemplateAMI", reflect.TypeOf((*MockEC2MachineInterface)(nil).DiscoverLaunchTemplateAMI), arg0)
}

// GetCoreSecurityGroups mocks base method.
func (m *MockEC2MachineInterface) GetCoreSecurityGroups(arg0 *scope.MachineScope) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCoreSecurityGroups", arg0)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCoreSecurityGroups indicates an expected call of GetCoreSecurityGroups.
func (mr *MockEC2MachineInterfaceMockRecorder) GetCoreSecurityGroups(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCoreSecurityGroups", reflect.TypeOf((*MockEC2MachineInterface)(nil).GetCoreSecurityGroups), arg0)
}

// GetFilteredSecurityGroupID mocks base method.
func (m *MockEC2MachineInterface) GetFilteredSecurityGroupID(arg0 v1beta1.AWSResourceReference) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFilteredSecurityGroupID", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFilteredSecurityGroupID indicates an expected call of GetFilteredSecurityGroupID.
func (mr *MockEC2MachineInterfaceMockRecorder) GetFilteredSecurityGroupID(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFilteredSecurityGroupID", reflect.TypeOf((*MockEC2MachineInterface)(nil).GetFilteredSecurityGroupID), arg0)
}

// GetInstanceSecurityGroups mocks base method.
func (m *MockEC2MachineInterface) GetInstanceSecurityGroups(arg0 string) (map[string][]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetInstanceSecurityGroups", arg0)
	ret0, _ := ret[0].(map[string][]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetInstanceSecurityGroups indicates an expected call of GetInstanceSecurityGroups.
func (mr *MockEC2MachineInterfaceMockRecorder) GetInstanceSecurityGroups(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetInstanceSecurityGroups", reflect.TypeOf((*MockEC2MachineInterface)(nil).GetInstanceSecurityGroups), arg0)
}

// GetLaunchTemplate mocks base method.
func (m *MockEC2MachineInterface) GetLaunchTemplate(arg0 string) (*v1beta10.AWSLaunchTemplate, string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLaunchTemplate", arg0)
	ret0, _ := ret[0].(*v1beta10.AWSLaunchTemplate)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetLaunchTemplate indicates an expected call of GetLaunchTemplate.
func (mr *MockEC2MachineInterfaceMockRecorder) GetLaunchTemplate(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLaunchTemplate", reflect.TypeOf((*MockEC2MachineInterface)(nil).GetLaunchTemplate), arg0)
}

// GetLaunchTemplateID mocks base method.
func (m *MockEC2MachineInterface) GetLaunchTemplateID(arg0 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLaunchTemplateID", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLaunchTemplateID indicates an expected call of GetLaunchTemplateID.
func (mr *MockEC2MachineInterfaceMockRecorder) GetLaunchTemplateID(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLaunchTemplateID", reflect.TypeOf((*MockEC2MachineInterface)(nil).GetLaunchTemplateID), arg0)
}

// GetRunningInstanceByTags mocks base method.
func (m *MockEC2MachineInterface) GetRunningInstanceByTags(arg0 *scope.MachineScope) (*v1beta1.Instance, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRunningInstanceByTags", arg0)
	ret0, _ := ret[0].(*v1beta1.Instance)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRunningInstanceByTags indicates an expected call of GetRunningInstanceByTags.
func (mr *MockEC2MachineInterfaceMockRecorder) GetRunningInstanceByTags(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRunningInstanceByTags", reflect.TypeOf((*MockEC2MachineInterface)(nil).GetRunningInstanceByTags), arg0)
}

// InstanceIfExists mocks base method.
func (m *MockEC2MachineInterface) InstanceIfExists(arg0 *string) (*v1beta1.Instance, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InstanceIfExists", arg0)
	ret0, _ := ret[0].(*v1beta1.Instance)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InstanceIfExists indicates an expected call of InstanceIfExists.
func (mr *MockEC2MachineInterfaceMockRecorder) InstanceIfExists(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InstanceIfExists", reflect.TypeOf((*MockEC2MachineInterface)(nil).InstanceIfExists), arg0)
}

// LaunchTemplateNeedsUpdate mocks base method.
func (m *MockEC2MachineInterface) LaunchTemplateNeedsUpdate(arg0 *scope.MachinePoolScope, arg1, arg2 *v1beta10.AWSLaunchTemplate) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LaunchTemplateNeedsUpdate", arg0, arg1, arg2)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LaunchTemplateNeedsUpdate indicates an expected call of LaunchTemplateNeedsUpdate.
func (mr *MockEC2MachineInterfaceMockRecorder) LaunchTemplateNeedsUpdate(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LaunchTemplateNeedsUpdate", reflect.TypeOf((*MockEC2MachineInterface)(nil).LaunchTemplateNeedsUpdate), arg0, arg1, arg2)
}

// PruneLaunchTemplateVersions mocks base method.
func (m *MockEC2MachineInterface) PruneLaunchTemplateVersions(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PruneLaunchTemplateVersions", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// PruneLaunchTemplateVersions indicates an expected call of PruneLaunchTemplateVersions.
func (mr *MockEC2MachineInterfaceMockRecorder) PruneLaunchTemplateVersions(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PruneLaunchTemplateVersions", reflect.TypeOf((*MockEC2MachineInterface)(nil).PruneLaunchTemplateVersions), arg0)
}

// ReconcileBastion mocks base method.
func (m *MockEC2MachineInterface) ReconcileBastion() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileBastion")
	ret0, _ := ret[0].(error)
	return ret0
}

// ReconcileBastion indicates an expected call of ReconcileBastion.
func (mr *MockEC2MachineInterfaceMockRecorder) ReconcileBastion() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileBastion", reflect.TypeOf((*MockEC2MachineInterface)(nil).ReconcileBastion))
}

// TerminateInstance mocks base method.
func (m *MockEC2MachineInterface) TerminateInstance(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TerminateInstance", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// TerminateInstance indicates an expected call of TerminateInstance.
func (mr *MockEC2MachineInterfaceMockRecorder) TerminateInstance(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TerminateInstance", reflect.TypeOf((*MockEC2MachineInterface)(nil).TerminateInstance), arg0)
}

// TerminateInstanceAndWait mocks base method.
func (m *MockEC2MachineInterface) TerminateInstanceAndWait(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TerminateInstanceAndWait", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// TerminateInstanceAndWait indicates an expected call of TerminateInstanceAndWait.
func (mr *MockEC2MachineInterfaceMockRecorder) TerminateInstanceAndWait(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TerminateInstanceAndWait", reflect.TypeOf((*MockEC2MachineInterface)(nil).TerminateInstanceAndWait), arg0)
}

// UpdateInstanceSecurityGroups mocks base method.
func (m *MockEC2MachineInterface) UpdateInstanceSecurityGroups(arg0 string, arg1 []string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateInstanceSecurityGroups", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateInstanceSecurityGroups indicates an expected call of UpdateInstanceSecurityGroups.
func (mr *MockEC2MachineInterfaceMockRecorder) UpdateInstanceSecurityGroups(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateInstanceSecurityGroups", reflect.TypeOf((*MockEC2MachineInterface)(nil).UpdateInstanceSecurityGroups), arg0, arg1)
}

// UpdateResourceTags mocks base method.
func (m *MockEC2MachineInterface) UpdateResourceTags(arg0 *string, arg1, arg2 map[string]string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateResourceTags", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateResourceTags indicates an expected call of UpdateResourceTags.
func (mr *MockEC2MachineInterfaceMockRecorder) UpdateResourceTags(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateResourceTags", reflect.TypeOf((*MockEC2MachineInterface)(nil).UpdateResourceTags), arg0, arg1, arg2)
}
