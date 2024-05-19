// Code generated by MockGen. DO NOT EDIT.
// Source: ./domains/rates/service.go

// Package rates is a generated GoMock package.
package rates

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockRateService is a mock of Service interface.
type MockRateService struct {
	ctrl     *gomock.Controller
	recorder *MockRateServiceMockRecorder
}

// MockRateServiceMockRecorder is the mock recorder for MockRateService.
type MockRateServiceMockRecorder struct {
	mock *MockRateService
}

// NewMockRateService creates a new mock instance.
func NewMockRateService(ctrl *gomock.Controller) *MockRateService {
	mock := &MockRateService{ctrl: ctrl}
	mock.recorder = &MockRateServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRateService) EXPECT() *MockRateServiceMockRecorder {
	return m.recorder
}

// Get mocks base method.
func (m *MockRateService) Get() (*float64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get")
	ret0, _ := ret[0].(*float64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockRateServiceMockRecorder) Get() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockRateService)(nil).Get))
}