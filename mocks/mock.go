// Code generated by MockGen. DO NOT EDIT.
// Source: main.go
//
// Generated by this command:
//
//	mockgen -source=main.go -destination=mocks/mock.go
//
// Package mock_main is a generated GoMock package.
package mock_main

import (
	image "image"
	jpeg "image/jpeg"
	io "io"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockCoder is a mock of Coder interface.
type MockCoder struct {
	ctrl     *gomock.Controller
	recorder *MockCoderMockRecorder
}

// MockCoderMockRecorder is the mock recorder for MockCoder.
type MockCoderMockRecorder struct {
	mock *MockCoder
}

// NewMockCoder creates a new mock instance.
func NewMockCoder(ctrl *gomock.Controller) *MockCoder {
	mock := &MockCoder{ctrl: ctrl}
	mock.recorder = &MockCoderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCoder) EXPECT() *MockCoderMockRecorder {
	return m.recorder
}

// Decode mocks base method.
func (m *MockCoder) Decode(r io.Reader) (image.Image, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Decode", r)
	ret0, _ := ret[0].(image.Image)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Decode indicates an expected call of Decode.
func (mr *MockCoderMockRecorder) Decode(r any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Decode", reflect.TypeOf((*MockCoder)(nil).Decode), r)
}

// Encode mocks base method.
func (m_2 *MockCoder) Encode(w io.Writer, m image.Image, o *jpeg.Options) error {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "Encode", w, m, o)
	ret0, _ := ret[0].(error)
	return ret0
}

// Encode indicates an expected call of Encode.
func (mr *MockCoderMockRecorder) Encode(w, m, o any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Encode", reflect.TypeOf((*MockCoder)(nil).Encode), w, m, o)
}