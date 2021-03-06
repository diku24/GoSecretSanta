//Code generated by MockGen. Do NOT EDIT.
// Source: service/person-services.go

// Pacakage mocks is a generated GoMock package.
package mocks

import (
	entity "SecretSanta/entity"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

//MockPersonServices is a mock of the PersonRepository interface
type MockPersonServices struct {
	ctrl     *gomock.Controller
	recorder *MockPersonServicesMockRecorder
}

// MockPersonServicesMockRecorder is the mock recorder for MockPersonServices.
type MockPersonServicesMockRecorder struct {
	mock *MockPersonServices
}

// NewMockPersonServices creates a new mock instance.
func NewMockPersonServices(ctrl *gomock.Controller) *MockPersonServices {
	mock := &MockPersonServices{ctrl: ctrl}
	mock.recorder = &MockPersonServicesMockRecorder{mock}
	return mock
}

//EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPersonServices) EXPECT() *MockPersonServicesMockRecorder {
	return m.recorder
}

// AllocateSanta mocks base method.
func (m *MockPersonServices) AllocateSanta() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AllocateSanta")
	ret0, _ := ret[0].(error)
	return ret0
}

//AllocateSanta indicate an expceted call of AllocateSanta.
func (mr *MockPersonServicesMockRecorder) AllocateSanta() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AllocateSanta", reflect.TypeOf((*MockPersonServices)(nil).AllocateSanta))
}

// CreatePersonWish mocks base methods
func (m *MockPersonServices) CreatePersonWish(personWish *entity.PersonWish) (*entity.PersonWish, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreatePersonWish", personWish)
	ret0, _ := ret[0].(*entity.PersonWish)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreatePersonWish indicates an expected call of CreatePersonWish
func (mr *MockPersonServicesMockRecorder) CreatePersonWish(personWish interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatePersonWish", reflect.TypeOf((*MockPersonServices)(nil).CreatePersonWish), personWish)
}

// GetAllWishes mocks base method.
func (m *MockPersonServices) GetAllWishes() ([]entity.PersonWish, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllWishes")
	ret0, _ := ret[0].([]entity.PersonWish)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllWishes indicates an expected call of GetAllWishes.
func (mr *MockPersonServicesMockRecorder) GetAllWishes() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllWishes", reflect.TypeOf((*MockPersonServices)(nil).GetAllWishes))
}
