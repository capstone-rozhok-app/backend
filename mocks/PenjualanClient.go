// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	penjualanclient "rozhok/features/penjualan_client"

	mock "github.com/stretchr/testify/mock"
)

// PenjualanClient is an autogenerated mock type for the PenjualanClientData type
type PenjualanClient struct {
	mock.Mock
}

// Delete provides a mock function with given fields: PenjualanClientCore
func (_m *PenjualanClient) Delete(PenjualanClientCore penjualanclient.Core) (int, error) {
	ret := _m.Called(PenjualanClientCore)

	var r0 int
	if rf, ok := ret.Get(0).(func(penjualanclient.Core) int); ok {
		r0 = rf(PenjualanClientCore)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(penjualanclient.Core) error); ok {
		r1 = rf(PenjualanClientCore)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAll provides a mock function with given fields: PenjualanClientCore
func (_m *PenjualanClient) GetAll(PenjualanClientCore penjualanclient.Core) ([]penjualanclient.Core, error) {
	ret := _m.Called(PenjualanClientCore)

	var r0 []penjualanclient.Core
	if rf, ok := ret.Get(0).(func(penjualanclient.Core) []penjualanclient.Core); ok {
		r0 = rf(PenjualanClientCore)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]penjualanclient.Core)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(penjualanclient.Core) error); ok {
		r1 = rf(PenjualanClientCore)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Insert provides a mock function with given fields: PenjualanClientCore
func (_m *PenjualanClient) Insert(PenjualanClientCore penjualanclient.Core) (int, error) {
	ret := _m.Called(PenjualanClientCore)

	var r0 int
	if rf, ok := ret.Get(0).(func(penjualanclient.Core) int); ok {
		r0 = rf(PenjualanClientCore)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(penjualanclient.Core) error); ok {
		r1 = rf(PenjualanClientCore)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: PenjualanClientCore
func (_m *PenjualanClient) Update(PenjualanClientCore penjualanclient.Core) (int, error) {
	ret := _m.Called(PenjualanClientCore)

	var r0 int
	if rf, ok := ret.Get(0).(func(penjualanclient.Core) int); ok {
		r0 = rf(PenjualanClientCore)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(penjualanclient.Core) error); ok {
		r1 = rf(PenjualanClientCore)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewPenjualanClient interface {
	mock.TestingT
	Cleanup(func())
}

// NewPenjualanClient creates a new instance of PenjualanClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewPenjualanClient(t mockConstructorTestingTNewPenjualanClient) *PenjualanClient {
	mock := &PenjualanClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
