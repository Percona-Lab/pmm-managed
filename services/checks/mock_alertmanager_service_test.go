// Code generated by mockery v1.0.0. DO NOT EDIT.

package checks

import (
	context "context"

	ammodels "github.com/percona/pmm/api/alertmanager/ammodels"

	mock "github.com/stretchr/testify/mock"
)

// mockAlertmanagerService is an autogenerated mock type for the alertmanagerService type
type mockAlertmanagerService struct {
	mock.Mock
}

// SendAlerts provides a mock function with given fields: ctx, alerts
func (_m *mockAlertmanagerService) SendAlerts(ctx context.Context, alerts ammodels.PostableAlerts) {
	_m.Called(ctx, alerts)
}
