package mocks

import (
	"api/internal/core/domain"
	"github.com/golang/mock/gomock"
	"reflect"
)

type MockUserMysqlRepository struct {
	ctrl     *gomock.Controller
	recorder *MockUserMysqlRepositoryRecorder
}

type MockUserMysqlRepositoryRecorder struct {
	mock *MockUserMysqlRepository
}

func NewMockUserMysqlRepository(ctrl *gomock.Controller) *MockUserMysqlRepository {
	mock := &MockUserMysqlRepository{ctrl: ctrl}
	mock.recorder = &MockUserMysqlRepositoryRecorder{mock: mock}
	return mock
}

func (m *MockUserMysqlRepository) EXPECT() *MockUserMysqlRepositoryRecorder {
	return m.recorder
}

func (m *MockUserMysqlRepository) Create(user *domain.User) (*domain.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", user)
	ret0, _ := ret[0].(*domain.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockUserMysqlRepositoryRecorder) Create(user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockUserMysqlRepository)(nil).Create), user)
}

func (m *MockUserMysqlRepository) List(queryParameter string) ([]domain.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", queryParameter)
	ret0, _ := ret[0].([]domain.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockUserMysqlRepositoryRecorder) List(queryParameter interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockUserMysqlRepository)(nil).List), queryParameter)
}

func (m *MockUserMysqlRepository) Get(id int) (*domain.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", id)
	ret0, _ := ret[0].(*domain.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockUserMysqlRepositoryRecorder) Get(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockUserMysqlRepository)(nil).Get), id)
}

func (m *MockUserMysqlRepository) Update(user *domain.User) (*domain.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", user)
	ret0, _ := ret[0].(*domain.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockUserMysqlRepositoryRecorder) Update(user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockUserMysqlRepository)(nil).Update), user)
}
