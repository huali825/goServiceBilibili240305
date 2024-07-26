// Code generated by MockGen. DO NOT EDIT.
// Source: 02webook/internal/repository/article/article_author.go
//
// Generated by this command:
//
//	mockgen -source=02webook/internal/repository/article/article_author.go -package=artrepomocks -destination=02webook/internal/repository/article/mocks/article_author.mock.go
//

// Package artrepomocks is a generated GoMock package.
package artrepomocks

import (
	context "context"
	domain "go20240218/02webook/internal/domain"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockArticleAuthorRepository is a mock of ArticleAuthorRepository interface.
type MockArticleAuthorRepository struct {
	ctrl     *gomock.Controller
	recorder *MockArticleAuthorRepositoryMockRecorder
}

// MockArticleAuthorRepositoryMockRecorder is the mock recorder for MockArticleAuthorRepository.
type MockArticleAuthorRepositoryMockRecorder struct {
	mock *MockArticleAuthorRepository
}

// NewMockArticleAuthorRepository creates a new mock instance.
func NewMockArticleAuthorRepository(ctrl *gomock.Controller) *MockArticleAuthorRepository {
	mock := &MockArticleAuthorRepository{ctrl: ctrl}
	mock.recorder = &MockArticleAuthorRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockArticleAuthorRepository) EXPECT() *MockArticleAuthorRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockArticleAuthorRepository) Create(ctx context.Context, art domain.Article) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, art)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockArticleAuthorRepositoryMockRecorder) Create(ctx, art any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockArticleAuthorRepository)(nil).Create), ctx, art)
}

// Update mocks base method.
func (m *MockArticleAuthorRepository) Update(ctx context.Context, art domain.Article) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, art)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockArticleAuthorRepositoryMockRecorder) Update(ctx, art any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockArticleAuthorRepository)(nil).Update), ctx, art)
}
