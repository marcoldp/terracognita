// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/cycloidio/terracognita/provider (interfaces: Resource)

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	filter "github.com/cycloidio/terracognita/filter"
	provider "github.com/cycloidio/terracognita/provider"
	writer "github.com/cycloidio/terracognita/writer"
	gomock "github.com/golang/mock/gomock"
	cty "github.com/hashicorp/go-cty/cty"
	schema "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	terraform "github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	states "github.com/hashicorp/terraform/states"
)

// Resource is a mock of Resource interface.
type Resource struct {
	ctrl     *gomock.Controller
	recorder *ResourceMockRecorder
}

// ResourceMockRecorder is the mock recorder for Resource.
type ResourceMockRecorder struct {
	mock *Resource
}

// NewResource creates a new mock instance.
func NewResource(ctrl *gomock.Controller) *Resource {
	mock := &Resource{ctrl: ctrl}
	mock.recorder = &ResourceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *Resource) EXPECT() *ResourceMockRecorder {
	return m.recorder
}

// AttributesReference mocks base method.
func (m *Resource) AttributesReference() ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AttributesReference")
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AttributesReference indicates an expected call of AttributesReference.
func (mr *ResourceMockRecorder) AttributesReference() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AttributesReference", reflect.TypeOf((*Resource)(nil).AttributesReference))
}

// Data mocks base method.
func (m *Resource) Data() *schema.ResourceData {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Data")
	ret0, _ := ret[0].(*schema.ResourceData)
	return ret0
}

// Data indicates an expected call of Data.
func (mr *ResourceMockRecorder) Data() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Data", reflect.TypeOf((*Resource)(nil).Data))
}

// HCL mocks base method.
func (m *Resource) HCL(arg0 writer.Writer) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HCL", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// HCL indicates an expected call of HCL.
func (mr *ResourceMockRecorder) HCL(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HCL", reflect.TypeOf((*Resource)(nil).HCL), arg0)
}

// ID mocks base method.
func (m *Resource) ID() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ID")
	ret0, _ := ret[0].(string)
	return ret0
}

// ID indicates an expected call of ID.
func (mr *ResourceMockRecorder) ID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ID", reflect.TypeOf((*Resource)(nil).ID))
}

// ImpliedType mocks base method.
func (m *Resource) ImpliedType() cty.Type {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ImpliedType")
	ret0, _ := ret[0].(cty.Type)
	return ret0
}

// ImpliedType indicates an expected call of ImpliedType.
func (mr *ResourceMockRecorder) ImpliedType() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ImpliedType", reflect.TypeOf((*Resource)(nil).ImpliedType))
}

// ImportState mocks base method.
func (m *Resource) ImportState() ([]provider.Resource, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ImportState")
	ret0, _ := ret[0].([]provider.Resource)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ImportState indicates an expected call of ImportState.
func (mr *ResourceMockRecorder) ImportState() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ImportState", reflect.TypeOf((*Resource)(nil).ImportState))
}

// InstanceInfo mocks base method.
func (m *Resource) InstanceInfo() *terraform.InstanceInfo {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InstanceInfo")
	ret0, _ := ret[0].(*terraform.InstanceInfo)
	return ret0
}

// InstanceInfo indicates an expected call of InstanceInfo.
func (mr *ResourceMockRecorder) InstanceInfo() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InstanceInfo", reflect.TypeOf((*Resource)(nil).InstanceInfo))
}

// InstanceState mocks base method.
func (m *Resource) InstanceState() *terraform.InstanceState {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InstanceState")
	ret0, _ := ret[0].(*terraform.InstanceState)
	return ret0
}

// InstanceState indicates an expected call of InstanceState.
func (mr *ResourceMockRecorder) InstanceState() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InstanceState", reflect.TypeOf((*Resource)(nil).InstanceState))
}

// Name mocks base method.
func (m *Resource) Name() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Name")
	ret0, _ := ret[0].(string)
	return ret0
}

// Name indicates an expected call of Name.
func (mr *ResourceMockRecorder) Name() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Name", reflect.TypeOf((*Resource)(nil).Name))
}

// Provider mocks base method.
func (m *Resource) Provider() provider.Provider {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Provider")
	ret0, _ := ret[0].(provider.Provider)
	return ret0
}

// Provider indicates an expected call of Provider.
func (mr *ResourceMockRecorder) Provider() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Provider", reflect.TypeOf((*Resource)(nil).Provider))
}

// Read mocks base method.
func (m *Resource) Read(arg0 *filter.Filter) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Read", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Read indicates an expected call of Read.
func (mr *ResourceMockRecorder) Read(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Read", reflect.TypeOf((*Resource)(nil).Read), arg0)
}

// ResourceInstanceObject mocks base method.
func (m *Resource) ResourceInstanceObject() *states.ResourceInstanceObject {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ResourceInstanceObject")
	ret0, _ := ret[0].(*states.ResourceInstanceObject)
	return ret0
}

// ResourceInstanceObject indicates an expected call of ResourceInstanceObject.
func (mr *ResourceMockRecorder) ResourceInstanceObject() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResourceInstanceObject", reflect.TypeOf((*Resource)(nil).ResourceInstanceObject))
}

// SetIgnoreTagFilter mocks base method.
func (m *Resource) SetIgnoreTagFilter(arg0 bool) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetIgnoreTagFilter", arg0)
}

// SetIgnoreTagFilter indicates an expected call of SetIgnoreTagFilter.
func (mr *ResourceMockRecorder) SetIgnoreTagFilter(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetIgnoreTagFilter", reflect.TypeOf((*Resource)(nil).SetIgnoreTagFilter), arg0)
}

// SetImporter mocks base method.
func (m *Resource) SetImporter(arg0 *schema.ResourceImporter) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetImporter", arg0)
}

// SetImporter indicates an expected call of SetImporter.
func (mr *ResourceMockRecorder) SetImporter(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetImporter", reflect.TypeOf((*Resource)(nil).SetImporter), arg0)
}

// State mocks base method.
func (m *Resource) State(arg0 writer.Writer) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "State", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// State indicates an expected call of State.
func (mr *ResourceMockRecorder) State(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "State", reflect.TypeOf((*Resource)(nil).State), arg0)
}

// TFResource mocks base method.
func (m *Resource) TFResource() *schema.Resource {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TFResource")
	ret0, _ := ret[0].(*schema.Resource)
	return ret0
}

// TFResource indicates an expected call of TFResource.
func (mr *ResourceMockRecorder) TFResource() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TFResource", reflect.TypeOf((*Resource)(nil).TFResource))
}

// Type mocks base method.
func (m *Resource) Type() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Type")
	ret0, _ := ret[0].(string)
	return ret0
}

// Type indicates an expected call of Type.
func (mr *ResourceMockRecorder) Type() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Type", reflect.TypeOf((*Resource)(nil).Type))
}
