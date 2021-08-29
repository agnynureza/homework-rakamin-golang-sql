// Code generated by MockGen. DO NOT EDIT.
// Source: ./services/movies.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	models "github.com/agnynureza/homework-rakamin-golang-sql/models"
	gomock "github.com/golang/mock/gomock"
)

// MockMovieServiceInterface is a mock of MovieServiceInterface interface.
type MockMovieServiceInterface struct {
	ctrl     *gomock.Controller
	recorder *MockMovieServiceInterfaceMockRecorder
}

// MockMovieServiceInterfaceMockRecorder is the mock recorder for MockMovieServiceInterface.
type MockMovieServiceInterfaceMockRecorder struct {
	mock *MockMovieServiceInterface
}

// NewMockMovieServiceInterface creates a new mock instance.
func NewMockMovieServiceInterface(ctrl *gomock.Controller) *MockMovieServiceInterface {
	mock := &MockMovieServiceInterface{ctrl: ctrl}
	mock.recorder = &MockMovieServiceInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMovieServiceInterface) EXPECT() *MockMovieServiceInterfaceMockRecorder {
	return m.recorder
}

// CreateNewMovie mocks base method.
func (m *MockMovieServiceInterface) CreateNewMovie(movie *models.Movies) (*models.Movies, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateNewMovie", movie)
	ret0, _ := ret[0].(*models.Movies)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateNewMovie indicates an expected call of CreateNewMovie.
func (mr *MockMovieServiceInterfaceMockRecorder) CreateNewMovie(movie interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateNewMovie", reflect.TypeOf((*MockMovieServiceInterface)(nil).CreateNewMovie), movie)
}

// DeleteMovie mocks base method.
func (m *MockMovieServiceInterface) DeleteMovie(slug string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteMovie", slug)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteMovie indicates an expected call of DeleteMovie.
func (mr *MockMovieServiceInterfaceMockRecorder) DeleteMovie(slug interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteMovie", reflect.TypeOf((*MockMovieServiceInterface)(nil).DeleteMovie), slug)
}

// GetMovie mocks base method.
func (m *MockMovieServiceInterface) GetMovie(slug string) (models.Movies, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMovie", slug)
	ret0, _ := ret[0].(models.Movies)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMovie indicates an expected call of GetMovie.
func (mr *MockMovieServiceInterfaceMockRecorder) GetMovie(slug interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMovie", reflect.TypeOf((*MockMovieServiceInterface)(nil).GetMovie), slug)
}

// UpdateMovie mocks base method.
func (m *MockMovieServiceInterface) UpdateMovie(movie *models.Movies, slug string) (models.Movies, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateMovie", movie, slug)
	ret0, _ := ret[0].(models.Movies)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateMovie indicates an expected call of UpdateMovie.
func (mr *MockMovieServiceInterfaceMockRecorder) UpdateMovie(movie, slug interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateMovie", reflect.TypeOf((*MockMovieServiceInterface)(nil).UpdateMovie), movie, slug)
}
