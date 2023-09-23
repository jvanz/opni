// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/util/notifier/types.go

// Package mock_notifier is a generated GoMock package.
package mock_notifier

import (
	context "context"
	reflect "reflect"

	notifier "github.com/rancher/opni/pkg/util/notifier"
	gomock "go.uber.org/mock/gomock"
)

// MockClonable is a mock of Clonable interface.
type MockClonable[T any] struct {
	ctrl     *gomock.Controller
	recorder *MockClonableMockRecorder[T]
}

// MockClonableMockRecorder is the mock recorder for MockClonable.
type MockClonableMockRecorder[T any] struct {
	mock *MockClonable[T]
}

// NewMockClonable creates a new mock instance.
func NewMockClonable[T any](ctrl *gomock.Controller) *MockClonable[T] {
	mock := &MockClonable[T]{ctrl: ctrl}
	mock.recorder = &MockClonableMockRecorder[T]{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClonable[T]) EXPECT() *MockClonableMockRecorder[T] {
	return m.recorder
}

// Clone mocks base method.
func (m *MockClonable[T]) Clone() T {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Clone")
	ret0, _ := ret[0].(T)
	return ret0
}

// Clone indicates an expected call of Clone.
func (mr *MockClonableMockRecorder[T]) Clone() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Clone", reflect.TypeOf((*MockClonable[T])(nil).Clone))
}

// MockFinder is a mock of Finder interface.
type MockFinder[T notifier.Clonable[T]] struct {
	ctrl     *gomock.Controller
	recorder *MockFinderMockRecorder[T]
}

// MockFinderMockRecorder is the mock recorder for MockFinder.
type MockFinderMockRecorder[T notifier.Clonable[T]] struct {
	mock *MockFinder[T]
}

// NewMockFinder creates a new mock instance.
func NewMockFinder[T notifier.Clonable[T]](ctrl *gomock.Controller) *MockFinder[T] {
	mock := &MockFinder[T]{ctrl: ctrl}
	mock.recorder = &MockFinderMockRecorder[T]{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFinder[T]) EXPECT() *MockFinderMockRecorder[T] {
	return m.recorder
}

// Find mocks base method.
func (m *MockFinder[T]) Find(ctx context.Context) ([]T, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Find", ctx)
	ret0, _ := ret[0].([]T)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Find indicates an expected call of Find.
func (mr *MockFinderMockRecorder[T]) Find(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockFinder[T])(nil).Find), ctx)
}

// MockUpdateNotifier is a mock of UpdateNotifier interface.
type MockUpdateNotifier[T any] struct {
	ctrl     *gomock.Controller
	recorder *MockUpdateNotifierMockRecorder[T]
}

// MockUpdateNotifierMockRecorder is the mock recorder for MockUpdateNotifier.
type MockUpdateNotifierMockRecorder[T any] struct {
	mock *MockUpdateNotifier[T]
}

// NewMockUpdateNotifier creates a new mock instance.
func NewMockUpdateNotifier[T any](ctrl *gomock.Controller) *MockUpdateNotifier[T] {
	mock := &MockUpdateNotifier[T]{ctrl: ctrl}
	mock.recorder = &MockUpdateNotifierMockRecorder[T]{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUpdateNotifier[T]) EXPECT() *MockUpdateNotifierMockRecorder[T] {
	return m.recorder
}

// NotifyC mocks base method.
func (m *MockUpdateNotifier[T]) NotifyC(ctx context.Context) <-chan []T {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NotifyC", ctx)
	ret0, _ := ret[0].(<-chan []T)
	return ret0
}

// NotifyC indicates an expected call of NotifyC.
func (mr *MockUpdateNotifierMockRecorder[T]) NotifyC(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NotifyC", reflect.TypeOf((*MockUpdateNotifier[T])(nil).NotifyC), ctx)
}

// Refresh mocks base method.
func (m *MockUpdateNotifier[T]) Refresh(ctx context.Context) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Refresh", ctx)
}

// Refresh indicates an expected call of Refresh.
func (mr *MockUpdateNotifierMockRecorder[T]) Refresh(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Refresh", reflect.TypeOf((*MockUpdateNotifier[T])(nil).Refresh), ctx)
}
