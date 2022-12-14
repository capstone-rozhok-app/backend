// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	pengambilanrosok "rozhok/features/pengambilan_rosok"

	mock "github.com/stretchr/testify/mock"
)

// PengambilanRosokStruct is an autogenerated mock type for the PengambilanRosokData type
type PengambilanRosokStruct struct {
	mock.Mock
}

// CreatePengambilanRosok provides a mock function with given fields: TransaksiCore
func (_m *PengambilanRosokStruct) CreatePengambilanRosok(TransaksiCore pengambilanrosok.Core) (int, error) {
	ret := _m.Called(TransaksiCore)

	var r0 int
	if rf, ok := ret.Get(0).(func(pengambilanrosok.Core) int); ok {
		r0 = rf(TransaksiCore)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(pengambilanrosok.Core) error); ok {
		r1 = rf(TransaksiCore)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Get provides a mock function with given fields: TransaksiCore
func (_m *PengambilanRosokStruct) Get(TransaksiCore pengambilanrosok.Core) (pengambilanrosok.Core, error) {
	ret := _m.Called(TransaksiCore)

	var r0 pengambilanrosok.Core
	if rf, ok := ret.Get(0).(func(pengambilanrosok.Core) pengambilanrosok.Core); ok {
		r0 = rf(TransaksiCore)
	} else {
		r0 = ret.Get(0).(pengambilanrosok.Core)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(pengambilanrosok.Core) error); ok {
		r1 = rf(TransaksiCore)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAll provides a mock function with given fields: TransaksiCore
func (_m *PengambilanRosokStruct) GetAll(TransaksiCore pengambilanrosok.Core) ([]pengambilanrosok.Core, error) {
	ret := _m.Called(TransaksiCore)

	var r0 []pengambilanrosok.Core
	if rf, ok := ret.Get(0).(func(pengambilanrosok.Core) []pengambilanrosok.Core); ok {
		r0 = rf(TransaksiCore)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]pengambilanrosok.Core)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(pengambilanrosok.Core) error); ok {
		r1 = rf(TransaksiCore)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewPengambilanRosokStruct interface {
	mock.TestingT
	Cleanup(func())
}

// NewPengambilanRosokStruct creates a new instance of PengambilanRosokStruct. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewPengambilanRosokStruct(t mockConstructorTestingTNewPengambilanRosokStruct) *PengambilanRosokStruct {
	mock := &PengambilanRosokStruct{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
