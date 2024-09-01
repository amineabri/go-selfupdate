// Code generated by MockGen. DO NOT EDIT.
// Source: updater.go

// Package mock_selfupdate is a generated GoMock package.
package mock

import (
	io "io"
	reflect "reflect"

	selfupdate "github.com/amineabri/go-selfupdate/selfupdate"
	semver "github.com/blang/semver"
	gomock "github.com/golang/mock/gomock"
)

// MockUpdaterIn is a mock of UpdaterIn interface.
type MockUpdaterIn struct {
	ctrl     *gomock.Controller
	recorder *MockUpdaterInMockRecorder
}

// MockUpdaterInMockRecorder is the mock recorder for MockUpdaterIn.
type MockUpdaterInMockRecorder struct {
	mock *MockUpdaterIn
}

// NewMockUpdaterIn creates a new mock instance.
func NewMockUpdaterIn(ctrl *gomock.Controller) *MockUpdaterIn {
	mock := &MockUpdaterIn{ctrl: ctrl}
	mock.recorder = &MockUpdaterInMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUpdaterIn) EXPECT() *MockUpdaterInMockRecorder {
	return m.recorder
}

// DetectLatest mocks base method.
func (m *MockUpdaterIn) DetectLatest(slug string) (*selfupdate.Release, bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DetectLatest", slug)
	ret0, _ := ret[0].(*selfupdate.Release)
	ret1, _ := ret[1].(bool)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// DetectLatest indicates an expected call of DetectLatest.
func (mr *MockUpdaterInMockRecorder) DetectLatest(slug interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DetectLatest", reflect.TypeOf((*MockUpdaterIn)(nil).DetectLatest), slug)
}

// DetectVersion mocks base method.
func (m *MockUpdaterIn) DetectVersion(slug, version string) (*selfupdate.Release, bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DetectVersion", slug, version)
	ret0, _ := ret[0].(*selfupdate.Release)
	ret1, _ := ret[1].(bool)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// DetectVersion indicates an expected call of DetectVersion.
func (mr *MockUpdaterInMockRecorder) DetectVersion(slug, version interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DetectVersion", reflect.TypeOf((*MockUpdaterIn)(nil).DetectVersion), slug, version)
}

// UpdateCommand mocks base method.
func (m *MockUpdaterIn) UpdateCommand(cmdPath string, current semver.Version, slug string) (*selfupdate.Release, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateCommand", cmdPath, current, slug)
	ret0, _ := ret[0].(*selfupdate.Release)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateCommand indicates an expected call of UpdateCommand.
func (mr *MockUpdaterInMockRecorder) UpdateCommand(cmdPath, current, slug interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateCommand", reflect.TypeOf((*MockUpdaterIn)(nil).UpdateCommand), cmdPath, current, slug)
}

// UpdateSelf mocks base method.
func (m *MockUpdaterIn) UpdateSelf(current semver.Version, slug string) (*selfupdate.Release, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateSelf", current, slug)
	ret0, _ := ret[0].(*selfupdate.Release)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateSelf indicates an expected call of UpdateSelf.
func (mr *MockUpdaterInMockRecorder) UpdateSelf(current, slug interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateSelf", reflect.TypeOf((*MockUpdaterIn)(nil).UpdateSelf), current, slug)
}

// UpdateTo mocks base method.
func (m *MockUpdaterIn) UpdateTo(rel *selfupdate.Release, cmdPath string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateTo", rel, cmdPath)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateTo indicates an expected call of UpdateTo.
func (mr *MockUpdaterInMockRecorder) UpdateTo(rel, cmdPath interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateTo", reflect.TypeOf((*MockUpdaterIn)(nil).UpdateTo), rel, cmdPath)
}

// downloadDirectlyFromURL mocks base method.
func (m *MockUpdaterIn) downloadDirectlyFromURL(assetURL string) (io.ReadCloser, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "downloadDirectlyFromURL", assetURL)
	ret0, _ := ret[0].(io.ReadCloser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// downloadDirectlyFromURL indicates an expected call of downloadDirectlyFromURL.
func (mr *MockUpdaterInMockRecorder) downloadDirectlyFromURL(assetURL interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "downloadDirectlyFromURL", reflect.TypeOf((*MockUpdaterIn)(nil).downloadDirectlyFromURL), assetURL)
}
