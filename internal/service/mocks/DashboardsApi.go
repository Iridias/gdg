// Code generated by mockery v2.34.0. DO NOT EDIT.

package mocks

import (
	filters "github.com/esnet/gdg/internal/service/filters"
	mock "github.com/stretchr/testify/mock"

	models "github.com/grafana/grafana-openapi-client-go/models"
)

// DashboardsApi is an autogenerated mock type for the DashboardsApi type
type DashboardsApi struct {
	mock.Mock
}

// DeleteAllDashboards provides a mock function with given fields: filter
func (_m *DashboardsApi) DeleteAllDashboards(filter filters.Filter) []string {
	ret := _m.Called(filter)

	var r0 []string
	if rf, ok := ret.Get(0).(func(filters.Filter) []string); ok {
		r0 = rf(filter)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	return r0
}

// DownloadDashboards provides a mock function with given fields: filter
func (_m *DashboardsApi) DownloadDashboards(filter filters.Filter) []string {
	ret := _m.Called(filter)

	var r0 []string
	if rf, ok := ret.Get(0).(func(filters.Filter) []string); ok {
		r0 = rf(filter)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	return r0
}

// ListDashboards provides a mock function with given fields: filter
func (_m *DashboardsApi) ListDashboards(filter filters.Filter) []*models.Hit {
	ret := _m.Called(filter)

	var r0 []*models.Hit
	if rf, ok := ret.Get(0).(func(filters.Filter) []*models.Hit); ok {
		r0 = rf(filter)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*models.Hit)
		}
	}

	return r0
}

// UploadDashboards provides a mock function with given fields: filter
func (_m *DashboardsApi) UploadDashboards(filter filters.Filter) {
	_m.Called(filter)
}

// NewDashboardsApi creates a new instance of DashboardsApi. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewDashboardsApi(t interface {
	mock.TestingT
	Cleanup(func())
}) *DashboardsApi {
	mock := &DashboardsApi{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
