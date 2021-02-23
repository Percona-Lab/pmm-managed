// Code generated by mockery v1.0.0. DO NOT EDIT.

package server

import (
	context "context"

	check "github.com/percona-platform/saas/pkg/check"

	mock "github.com/stretchr/testify/mock"
)

// mockChecksService is an autogenerated mock type for the checksService type
type mockChecksService struct {
	mock.Mock
}

// CleanupAlerts provides a mock function with given fields:
func (_m *mockChecksService) CleanupAlerts() {
	_m.Called()
}

// StartChecks provides a mock function with given fields: ctx, interval
func (_m *mockChecksService) StartChecks(ctx context.Context, interval check.Interval) error {
	ret := _m.Called(ctx, interval)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, check.Interval) error); ok {
		r0 = rf(ctx, interval)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
