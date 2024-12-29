// Code generated by MockGen. DO NOT EDIT.
// Source: generator.go

// Package staticnetworkconfig is a generated GoMock package.
package staticnetworkconfig

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	models "github.com/openshift/assisted-service/models"
)

// MockStaticNetworkConfig is a mock of StaticNetworkConfig interface.
type MockStaticNetworkConfig struct {
	ctrl     *gomock.Controller
	recorder *MockStaticNetworkConfigMockRecorder
}

// MockStaticNetworkConfigMockRecorder is the mock recorder for MockStaticNetworkConfig.
type MockStaticNetworkConfigMockRecorder struct {
	mock *MockStaticNetworkConfig
}

// NewMockStaticNetworkConfig creates a new mock instance.
func NewMockStaticNetworkConfig(ctrl *gomock.Controller) *MockStaticNetworkConfig {
	mock := &MockStaticNetworkConfig{ctrl: ctrl}
	mock.recorder = &MockStaticNetworkConfigMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStaticNetworkConfig) EXPECT() *MockStaticNetworkConfigMockRecorder {
	return m.recorder
}

// FormatStaticNetworkConfigForDB mocks base method.
func (m *MockStaticNetworkConfig) FormatStaticNetworkConfigForDB(staticNetworkConfig []*models.HostStaticNetworkConfig) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FormatStaticNetworkConfigForDB", staticNetworkConfig)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FormatStaticNetworkConfigForDB indicates an expected call of FormatStaticNetworkConfigForDB.
func (mr *MockStaticNetworkConfigMockRecorder) FormatStaticNetworkConfigForDB(staticNetworkConfig interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FormatStaticNetworkConfigForDB", reflect.TypeOf((*MockStaticNetworkConfig)(nil).FormatStaticNetworkConfigForDB), staticNetworkConfig)
}

// GenerateStaticNetworkConfigData mocks base method.
func (m *MockStaticNetworkConfig) GenerateStaticNetworkConfigData(ctx context.Context, hostsYAMLS string) ([]StaticNetworkConfigData, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateStaticNetworkConfigData", ctx, hostsYAMLS)
	ret0, _ := ret[0].([]StaticNetworkConfigData)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GenerateStaticNetworkConfigData indicates an expected call of GenerateStaticNetworkConfigData.
func (mr *MockStaticNetworkConfigMockRecorder) GenerateStaticNetworkConfigData(ctx, hostsYAMLS interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateStaticNetworkConfigData", reflect.TypeOf((*MockStaticNetworkConfig)(nil).GenerateStaticNetworkConfigData), ctx, hostsYAMLS)
}

// GenerateStaticNetworkConfigDataYAML mocks base method.
func (m *MockStaticNetworkConfig) GenerateStaticNetworkConfigDataYAML(staticNetworkConfigStr string) ([]StaticNetworkConfigData, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateStaticNetworkConfigDataYAML", staticNetworkConfigStr)
	ret0, _ := ret[0].([]StaticNetworkConfigData)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GenerateStaticNetworkConfigDataYAML indicates an expected call of GenerateStaticNetworkConfigDataYAML.
func (mr *MockStaticNetworkConfigMockRecorder) GenerateStaticNetworkConfigDataYAML(staticNetworkConfigStr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateStaticNetworkConfigDataYAML", reflect.TypeOf((*MockStaticNetworkConfig)(nil).GenerateStaticNetworkConfigDataYAML), staticNetworkConfigStr)
}

// ShouldUseNmstateService mocks base method.
func (m *MockStaticNetworkConfig) ShouldUseNmstateService(staticNetworkConfigStr, openshiftVersion string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ShouldUseNmstateService", staticNetworkConfigStr, openshiftVersion)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ShouldUseNmstateService indicates an expected call of ShouldUseNmstateService.
func (mr *MockStaticNetworkConfigMockRecorder) ShouldUseNmstateService(staticNetworkConfigStr, openshiftVersion interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ShouldUseNmstateService", reflect.TypeOf((*MockStaticNetworkConfig)(nil).ShouldUseNmstateService), staticNetworkConfigStr, openshiftVersion)
}

// ValidateStaticConfigParamsYAML mocks base method.
func (m *MockStaticNetworkConfig) ValidateStaticConfigParamsYAML(staticNetworkConfig []*models.HostStaticNetworkConfig) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValidateStaticConfigParamsYAML", staticNetworkConfig)
	ret0, _ := ret[0].(error)
	return ret0
}

// ValidateStaticConfigParamsYAML indicates an expected call of ValidateStaticConfigParamsYAML.
func (mr *MockStaticNetworkConfigMockRecorder) ValidateStaticConfigParamsYAML(staticNetworkConfig interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidateStaticConfigParamsYAML", reflect.TypeOf((*MockStaticNetworkConfig)(nil).ValidateStaticConfigParamsYAML), staticNetworkConfig)
}
