// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/ProtonMail/proton-bridge/v3/internal/user (interfaces: MessageDownloader)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	io "io"
	reflect "reflect"

	proton "github.com/ProtonMail/go-proton-api"
	gomock "github.com/golang/mock/gomock"
)

// MockMessageDownloader is a mock of MessageDownloader interface.
type MockMessageDownloader struct {
	ctrl     *gomock.Controller
	recorder *MockMessageDownloaderMockRecorder
}

// MockMessageDownloaderMockRecorder is the mock recorder for MockMessageDownloader.
type MockMessageDownloaderMockRecorder struct {
	mock *MockMessageDownloader
}

// NewMockMessageDownloader creates a new mock instance.
func NewMockMessageDownloader(ctrl *gomock.Controller) *MockMessageDownloader {
	mock := &MockMessageDownloader{ctrl: ctrl}
	mock.recorder = &MockMessageDownloaderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMessageDownloader) EXPECT() *MockMessageDownloaderMockRecorder {
	return m.recorder
}

// GetAttachmentInto mocks base method.
func (m *MockMessageDownloader) GetAttachmentInto(arg0 context.Context, arg1 string, arg2 io.ReaderFrom) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAttachmentInto", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// GetAttachmentInto indicates an expected call of GetAttachmentInto.
func (mr *MockMessageDownloaderMockRecorder) GetAttachmentInto(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAttachmentInto", reflect.TypeOf((*MockMessageDownloader)(nil).GetAttachmentInto), arg0, arg1, arg2)
}

// GetMessage mocks base method.
func (m *MockMessageDownloader) GetMessage(arg0 context.Context, arg1 string) (proton.Message, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMessage", arg0, arg1)
	ret0, _ := ret[0].(proton.Message)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMessage indicates an expected call of GetMessage.
func (mr *MockMessageDownloaderMockRecorder) GetMessage(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMessage", reflect.TypeOf((*MockMessageDownloader)(nil).GetMessage), arg0, arg1)
}
