// Code generated by MockGen. DO NOT EDIT.
// Source: inmemory_strategy.go

// Package mock_strategy is a generated GoMock package.
package mock_strategy

import (
	blogging "blogging/proto"
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockBlogStrategy is a mock of BlogStrategy interface.
type MockBlogStrategy struct {
	ctrl     *gomock.Controller
	recorder *MockBlogStrategyMockRecorder
}

// MockBlogStrategyMockRecorder is the mock recorder for MockBlogStrategy.
type MockBlogStrategyMockRecorder struct {
	mock *MockBlogStrategy
}

// NewMockBlogStrategy creates a new mock instance.
func NewMockBlogStrategy(ctrl *gomock.Controller) *MockBlogStrategy {
	mock := &MockBlogStrategy{ctrl: ctrl}
	mock.recorder = &MockBlogStrategyMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBlogStrategy) EXPECT() *MockBlogStrategyMockRecorder {
	return m.recorder
}

// CreatePost mocks base method.
func (m *MockBlogStrategy) CreatePost(ctx context.Context, req *blogging.CreatePostRequest) (*blogging.CreatePostResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreatePost", ctx, req)
	ret0, _ := ret[0].(*blogging.CreatePostResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreatePost indicates an expected call of CreatePost.
func (mr *MockBlogStrategyMockRecorder) CreatePost(ctx, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatePost", reflect.TypeOf((*MockBlogStrategy)(nil).CreatePost), ctx, req)
}

// DeletePost mocks base method.
func (m *MockBlogStrategy) DeletePost(ctx context.Context, req *blogging.DeletePostRequest) (*blogging.DeletePostResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeletePost", ctx, req)
	ret0, _ := ret[0].(*blogging.DeletePostResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeletePost indicates an expected call of DeletePost.
func (mr *MockBlogStrategyMockRecorder) DeletePost(ctx, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeletePost", reflect.TypeOf((*MockBlogStrategy)(nil).DeletePost), ctx, req)
}

// ReadPost mocks base method.
func (m *MockBlogStrategy) ReadPost(ctx context.Context, req *blogging.GetPostRequest) (*blogging.GetPostResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadPost", ctx, req)
	ret0, _ := ret[0].(*blogging.GetPostResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadPost indicates an expected call of ReadPost.
func (mr *MockBlogStrategyMockRecorder) ReadPost(ctx, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadPost", reflect.TypeOf((*MockBlogStrategy)(nil).ReadPost), ctx, req)
}

// UpdatePost mocks base method.
func (m *MockBlogStrategy) UpdatePost(ctx context.Context, req *blogging.UpdatePostRequest) (*blogging.UpdatePostResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdatePost", ctx, req)
	ret0, _ := ret[0].(*blogging.UpdatePostResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdatePost indicates an expected call of UpdatePost.
func (mr *MockBlogStrategyMockRecorder) UpdatePost(ctx, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePost", reflect.TypeOf((*MockBlogStrategy)(nil).UpdatePost), ctx, req)
}
