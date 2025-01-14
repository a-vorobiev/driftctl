// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	resource "github.com/cloudskiff/driftctl/pkg/resource"
	mock "github.com/stretchr/testify/mock"
)

// Filter is an autogenerated mock type for the Filter type
type Filter struct {
	mock.Mock
}

// IsFieldIgnored provides a mock function with given fields: res, path
func (_m *Filter) IsFieldIgnored(res resource.Resource, path []string) bool {
	ret := _m.Called(res, path)

	var r0 bool
	if rf, ok := ret.Get(0).(func(resource.Resource, []string) bool); ok {
		r0 = rf(res, path)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// IsResourceIgnored provides a mock function with given fields: res
func (_m *Filter) IsResourceIgnored(res resource.Resource) bool {
	ret := _m.Called(res)

	var r0 bool
	if rf, ok := ret.Get(0).(func(resource.Resource) bool); ok {
		r0 = rf(res)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}
