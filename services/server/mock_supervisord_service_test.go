// Code generated by mockery v1.0.0. DO NOT EDIT.

package server

import mock "github.com/stretchr/testify/mock"

// mockSupervisordService is an autogenerated mock type for the supervisordService type
type mockSupervisordService struct {
	mock.Mock
}

// StartPMMUpdate provides a mock function with given fields:
func (_m *mockSupervisordService) StartPMMUpdate() (uint32, error) {
	ret := _m.Called()

	var r0 uint32
	if rf, ok := ret.Get(0).(func() uint32); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(uint32)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
