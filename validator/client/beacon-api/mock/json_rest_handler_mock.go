// Code generated by MockGen. DO NOT EDIT.
// Source: validator/client/beacon-api/json_rest_handler.go

// Package mock is a generated GoMock package.
package mock

import (
	bytes "bytes"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	apimiddleware "github.com/prysmaticlabs/prysm/v3/api/gateway/apimiddleware"
)

// MockjsonRestHandler is a mock of jsonRestHandler interface.
type MockjsonRestHandler struct {
	ctrl     *gomock.Controller
	recorder *MockjsonRestHandlerMockRecorder
}

// MockjsonRestHandlerMockRecorder is the mock recorder for MockjsonRestHandler.
type MockjsonRestHandlerMockRecorder struct {
	mock *MockjsonRestHandler
}

// NewMockjsonRestHandler creates a new mock instance.
func NewMockjsonRestHandler(ctrl *gomock.Controller) *MockjsonRestHandler {
	mock := &MockjsonRestHandler{ctrl: ctrl}
	mock.recorder = &MockjsonRestHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockjsonRestHandler) EXPECT() *MockjsonRestHandlerMockRecorder {
	return m.recorder
}

// GetRestJsonResponse mocks base method.
func (m *MockjsonRestHandler) GetRestJsonResponse(query string, responseJson interface{}) (*apimiddleware.DefaultErrorJson, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRestJsonResponse", query, responseJson)
	ret0, _ := ret[0].(*apimiddleware.DefaultErrorJson)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRestJsonResponse indicates an expected call of GetRestJsonResponse.
func (mr *MockjsonRestHandlerMockRecorder) GetRestJsonResponse(query, responseJson interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRestJsonResponse", reflect.TypeOf((*MockjsonRestHandler)(nil).GetRestJsonResponse), query, responseJson)
}

// PostRestJson mocks base method.
func (m *MockjsonRestHandler) PostRestJson(apiEndpoint string, headers map[string]string, data *bytes.Buffer, responseJson interface{}) (*apimiddleware.DefaultErrorJson, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PostRestJson", apiEndpoint, headers, data, responseJson)
	ret0, _ := ret[0].(*apimiddleware.DefaultErrorJson)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PostRestJson indicates an expected call of PostRestJson.
func (mr *MockjsonRestHandlerMockRecorder) PostRestJson(apiEndpoint, headers, data, responseJson interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PostRestJson", reflect.TypeOf((*MockjsonRestHandler)(nil).PostRestJson), apiEndpoint, headers, data, responseJson)
}
