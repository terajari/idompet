// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/terajari/idompet/database/sqlc (interfaces: Store)

// Package mockdatabase is a generated GoMock package.
package mockdatabase

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	database "github.com/terajari/idompet/database/sqlc"
)

// MockStore is a mock of Store interface.
type MockStore struct {
	ctrl     *gomock.Controller
	recorder *MockStoreMockRecorder
}

// MockStoreMockRecorder is the mock recorder for MockStore.
type MockStoreMockRecorder struct {
	mock *MockStore
}

// NewMockStore creates a new mock instance.
func NewMockStore(ctrl *gomock.Controller) *MockStore {
	mock := &MockStore{ctrl: ctrl}
	mock.recorder = &MockStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStore) EXPECT() *MockStoreMockRecorder {
	return m.recorder
}

// ChangeSaldo mocks base method.
func (m *MockStore) ChangeSaldo(arg0 context.Context, arg1 database.ChangeSaldoParams) (database.Akun, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChangeSaldo", arg0, arg1)
	ret0, _ := ret[0].(database.Akun)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ChangeSaldo indicates an expected call of ChangeSaldo.
func (mr *MockStoreMockRecorder) ChangeSaldo(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChangeSaldo", reflect.TypeOf((*MockStore)(nil).ChangeSaldo), arg0, arg1)
}

// CreateAkun mocks base method.
func (m *MockStore) CreateAkun(arg0 context.Context, arg1 database.CreateAkunParams) (database.Akun, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateAkun", arg0, arg1)
	ret0, _ := ret[0].(database.Akun)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateAkun indicates an expected call of CreateAkun.
func (mr *MockStoreMockRecorder) CreateAkun(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAkun", reflect.TypeOf((*MockStore)(nil).CreateAkun), arg0, arg1)
}

// CreateCatatan mocks base method.
func (m *MockStore) CreateCatatan(arg0 context.Context, arg1 database.CreateCatatanParams) (database.Catatan, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateCatatan", arg0, arg1)
	ret0, _ := ret[0].(database.Catatan)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateCatatan indicates an expected call of CreateCatatan.
func (mr *MockStoreMockRecorder) CreateCatatan(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCatatan", reflect.TypeOf((*MockStore)(nil).CreateCatatan), arg0, arg1)
}

// CreateTransfer mocks base method.
func (m *MockStore) CreateTransfer(arg0 context.Context, arg1 database.CreateTransferParams) (database.Transfer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTransfer", arg0, arg1)
	ret0, _ := ret[0].(database.Transfer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateTransfer indicates an expected call of CreateTransfer.
func (mr *MockStoreMockRecorder) CreateTransfer(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTransfer", reflect.TypeOf((*MockStore)(nil).CreateTransfer), arg0, arg1)
}

// DeleteAkun mocks base method.
func (m *MockStore) DeleteAkun(arg0 context.Context, arg1 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAkun", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteAkun indicates an expected call of DeleteAkun.
func (mr *MockStoreMockRecorder) DeleteAkun(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAkun", reflect.TypeOf((*MockStore)(nil).DeleteAkun), arg0, arg1)
}

// GetAkun mocks base method.
func (m *MockStore) GetAkun(arg0 context.Context, arg1 int64) (database.Akun, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAkun", arg0, arg1)
	ret0, _ := ret[0].(database.Akun)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAkun indicates an expected call of GetAkun.
func (mr *MockStoreMockRecorder) GetAkun(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAkun", reflect.TypeOf((*MockStore)(nil).GetAkun), arg0, arg1)
}

// GetAkunForUpdate mocks base method.
func (m *MockStore) GetAkunForUpdate(arg0 context.Context, arg1 int64) (database.Akun, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAkunForUpdate", arg0, arg1)
	ret0, _ := ret[0].(database.Akun)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAkunForUpdate indicates an expected call of GetAkunForUpdate.
func (mr *MockStoreMockRecorder) GetAkunForUpdate(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAkunForUpdate", reflect.TypeOf((*MockStore)(nil).GetAkunForUpdate), arg0, arg1)
}

// GetCatatan mocks base method.
func (m *MockStore) GetCatatan(arg0 context.Context, arg1 int64) (database.Catatan, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCatatan", arg0, arg1)
	ret0, _ := ret[0].(database.Catatan)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCatatan indicates an expected call of GetCatatan.
func (mr *MockStoreMockRecorder) GetCatatan(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCatatan", reflect.TypeOf((*MockStore)(nil).GetCatatan), arg0, arg1)
}

// GetTransfer mocks base method.
func (m *MockStore) GetTransfer(arg0 context.Context, arg1 int64) (database.Transfer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTransfer", arg0, arg1)
	ret0, _ := ret[0].(database.Transfer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTransfer indicates an expected call of GetTransfer.
func (mr *MockStoreMockRecorder) GetTransfer(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTransfer", reflect.TypeOf((*MockStore)(nil).GetTransfer), arg0, arg1)
}

// ListAkun mocks base method.
func (m *MockStore) ListAkun(arg0 context.Context, arg1 database.ListAkunParams) ([]database.Akun, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAkun", arg0, arg1)
	ret0, _ := ret[0].([]database.Akun)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAkun indicates an expected call of ListAkun.
func (mr *MockStoreMockRecorder) ListAkun(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAkun", reflect.TypeOf((*MockStore)(nil).ListAkun), arg0, arg1)
}

// ListCatatan mocks base method.
func (m *MockStore) ListCatatan(arg0 context.Context, arg1 database.ListCatatanParams) ([]database.Catatan, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListCatatan", arg0, arg1)
	ret0, _ := ret[0].([]database.Catatan)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListCatatan indicates an expected call of ListCatatan.
func (mr *MockStoreMockRecorder) ListCatatan(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListCatatan", reflect.TypeOf((*MockStore)(nil).ListCatatan), arg0, arg1)
}

// PerformTransfer mocks base method.
func (m *MockStore) PerformTransfer(arg0 context.Context, arg1 database.TransferTxParams) (*database.TransferTxResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PerformTransfer", arg0, arg1)
	ret0, _ := ret[0].(*database.TransferTxResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PerformTransfer indicates an expected call of PerformTransfer.
func (mr *MockStoreMockRecorder) PerformTransfer(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PerformTransfer", reflect.TypeOf((*MockStore)(nil).PerformTransfer), arg0, arg1)
}

// UpdateAkun mocks base method.
func (m *MockStore) UpdateAkun(arg0 context.Context, arg1 database.UpdateAkunParams) (database.Akun, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAkun", arg0, arg1)
	ret0, _ := ret[0].(database.Akun)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateAkun indicates an expected call of UpdateAkun.
func (mr *MockStoreMockRecorder) UpdateAkun(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAkun", reflect.TypeOf((*MockStore)(nil).UpdateAkun), arg0, arg1)
}
