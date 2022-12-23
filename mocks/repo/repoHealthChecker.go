// Code generated by MockGen. DO NOT EDIT.
// Source: repo/repoHealthChecker.go

// Package mock_repo is a generated GoMock package.
package mock_repo

import (
	reflect "reflect"
	domain "tashilkar_health_checker/domain"

	gomock "github.com/golang/mock/gomock"
	primitive "go.mongodb.org/mongo-driver/bson/primitive"
)

// MockHealthChecker is a mock of HealthChecker interface.
type MockHealthChecker struct {
	ctrl     *gomock.Controller
	recorder *MockHealthCheckerMockRecorder
}

// MockHealthCheckerMockRecorder is the mock recorder for MockHealthChecker.
type MockHealthCheckerMockRecorder struct {
	mock *MockHealthChecker
}

// NewMockHealthChecker creates a new mock instance.
func NewMockHealthChecker(ctrl *gomock.Controller) *MockHealthChecker {
	mock := &MockHealthChecker{ctrl: ctrl}
	mock.recorder = &MockHealthCheckerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockHealthChecker) EXPECT() *MockHealthCheckerMockRecorder {
	return m.recorder
}

// DeleteApi mocks base method.
func (m *MockHealthChecker) DeleteApi(id primitive.ObjectID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteApi", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteApi indicates an expected call of DeleteApi.
func (mr *MockHealthCheckerMockRecorder) DeleteApi(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteApi", reflect.TypeOf((*MockHealthChecker)(nil).DeleteApi), id)
}

// GetApiLists mocks base method.
func (m *MockHealthChecker) GetApiLists() ([]domain.Api, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetApiLists")
	ret0, _ := ret[0].([]domain.Api)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetApiLists indicates an expected call of GetApiLists.
func (mr *MockHealthCheckerMockRecorder) GetApiLists() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetApiLists", reflect.TypeOf((*MockHealthChecker)(nil).GetApiLists))
}

// GetStatus mocks base method.
func (m *MockHealthChecker) GetStatus() (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStatus")
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetStatus indicates an expected call of GetStatus.
func (mr *MockHealthCheckerMockRecorder) GetStatus() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStatus", reflect.TypeOf((*MockHealthChecker)(nil).GetStatus))
}

// InsertCheckedEndPoint mocks base method.
func (m *MockHealthChecker) InsertCheckedEndPoint(request domain.CheckedApi) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "InsertCheckedEndPoint", request)
}

// InsertCheckedEndPoint indicates an expected call of InsertCheckedEndPoint.
func (mr *MockHealthCheckerMockRecorder) InsertCheckedEndPoint(request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertCheckedEndPoint", reflect.TypeOf((*MockHealthChecker)(nil).InsertCheckedEndPoint), request)
}

// InsertNewEndPoint mocks base method.
func (m *MockHealthChecker) InsertNewEndPoint(request domain.RegisterApiReq) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertNewEndPoint", request)
	ret0, _ := ret[0].(error)
	return ret0
}

// InsertNewEndPoint indicates an expected call of InsertNewEndPoint.
func (mr *MockHealthCheckerMockRecorder) InsertNewEndPoint(request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertNewEndPoint", reflect.TypeOf((*MockHealthChecker)(nil).InsertNewEndPoint), request)
}

// SetStatus mocks base method.
func (m *MockHealthChecker) SetStatus(availability domain.HealthCheckerAvailability) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetStatus", availability)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetStatus indicates an expected call of SetStatus.
func (mr *MockHealthCheckerMockRecorder) SetStatus(availability interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetStatus", reflect.TypeOf((*MockHealthChecker)(nil).SetStatus), availability)
}
