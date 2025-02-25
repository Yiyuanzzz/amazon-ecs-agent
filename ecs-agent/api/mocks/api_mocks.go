// Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.
//

// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/aws/amazon-ecs-agent/ecs-agent/api (interfaces: ECSDiscoverEndpointSDK)

// Package mock_api is a generated GoMock package.
package mock_api

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockECSDiscoverEndpointSDK is a mock of ECSDiscoverEndpointSDK interface.
type MockECSDiscoverEndpointSDK struct {
	ctrl     *gomock.Controller
	recorder *MockECSDiscoverEndpointSDKMockRecorder
}

// MockECSDiscoverEndpointSDKMockRecorder is the mock recorder for MockECSDiscoverEndpointSDK.
type MockECSDiscoverEndpointSDKMockRecorder struct {
	mock *MockECSDiscoverEndpointSDK
}

// NewMockECSDiscoverEndpointSDK creates a new mock instance.
func NewMockECSDiscoverEndpointSDK(ctrl *gomock.Controller) *MockECSDiscoverEndpointSDK {
	mock := &MockECSDiscoverEndpointSDK{ctrl: ctrl}
	mock.recorder = &MockECSDiscoverEndpointSDKMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockECSDiscoverEndpointSDK) EXPECT() *MockECSDiscoverEndpointSDKMockRecorder {
	return m.recorder
}

// DiscoverPollEndpoint mocks base method.
func (m *MockECSDiscoverEndpointSDK) DiscoverPollEndpoint(arg0 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DiscoverPollEndpoint", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DiscoverPollEndpoint indicates an expected call of DiscoverPollEndpoint.
func (mr *MockECSDiscoverEndpointSDKMockRecorder) DiscoverPollEndpoint(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DiscoverPollEndpoint", reflect.TypeOf((*MockECSDiscoverEndpointSDK)(nil).DiscoverPollEndpoint), arg0)
}

// DiscoverServiceConnectEndpoint mocks base method.
func (m *MockECSDiscoverEndpointSDK) DiscoverServiceConnectEndpoint(arg0 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DiscoverServiceConnectEndpoint", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DiscoverServiceConnectEndpoint indicates an expected call of DiscoverServiceConnectEndpoint.
func (mr *MockECSDiscoverEndpointSDKMockRecorder) DiscoverServiceConnectEndpoint(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DiscoverServiceConnectEndpoint", reflect.TypeOf((*MockECSDiscoverEndpointSDK)(nil).DiscoverServiceConnectEndpoint), arg0)
}

// DiscoverTelemetryEndpoint mocks base method.
func (m *MockECSDiscoverEndpointSDK) DiscoverTelemetryEndpoint(arg0 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DiscoverTelemetryEndpoint", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DiscoverTelemetryEndpoint indicates an expected call of DiscoverTelemetryEndpoint.
func (mr *MockECSDiscoverEndpointSDKMockRecorder) DiscoverTelemetryEndpoint(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DiscoverTelemetryEndpoint", reflect.TypeOf((*MockECSDiscoverEndpointSDK)(nil).DiscoverTelemetryEndpoint), arg0)
}
