// Code generated by mockery v1.0.0. DO NOT EDIT.

package server

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// mockVmAlertAlertingRules is an autogenerated mock type for the vmAlertAlertingRules type
type mockVmAlertAlertingRules struct {
	mock.Mock
}

// GetRulesHash provides a mock function with given fields:
func (_m *mockVmAlertAlertingRules) GetRulesHash() ([]byte, error) {
	ret := _m.Called()

	var r0 []byte
	if rf, ok := ret.Get(0).(func() []byte); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ReadRules provides a mock function with given fields:
func (_m *mockVmAlertAlertingRules) ReadRules() (string, error) {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RemoveRulesFile provides a mock function with given fields:
func (_m *mockVmAlertAlertingRules) RemoveRulesFile() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ValidateRules provides a mock function with given fields: ctx, rules
func (_m *mockVmAlertAlertingRules) ValidateRules(ctx context.Context, rules string) error {
	ret := _m.Called(ctx, rules)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, rules)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// WriteRules provides a mock function with given fields: rules
func (_m *mockVmAlertAlertingRules) WriteRules(rules string) error {
	ret := _m.Called(rules)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(rules)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}