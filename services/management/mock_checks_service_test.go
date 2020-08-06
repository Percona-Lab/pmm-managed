// Code generated by mockery v1.0.0. DO NOT EDIT.

package management

import (
	context "context"

	check "github.com/percona-platform/saas/pkg/check"

	mock "github.com/stretchr/testify/mock"
)

// mockChecksService is an autogenerated mock type for the checksService type
type mockChecksService struct {
	mock.Mock
}

// GetSecurityCheckResults provides a mock function with given fields: ctx
func (_m *mockChecksService) GetSecurityCheckResults(ctx context.Context) ([]check.Result, error) {
	ret := _m.Called(ctx)

	var r0 []check.Result
	if rf, ok := ret.Get(0).(func(context.Context) []check.Result); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]check.Result)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// StartChecks provides a mock function with given fields: ctx
func (_m *mockChecksService) StartChecks(ctx context.Context) error {
	ret := _m.Called(ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
