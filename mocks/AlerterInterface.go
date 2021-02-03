// Code generated by mockery v2.3.0. DO NOT EDIT.

package mocks

import (
	alerter "github.com/cloudskiff/driftctl/pkg/alerter"
	mock "github.com/stretchr/testify/mock"
)

// AlerterInterface is an autogenerated mock type for the AlerterInterface type
type AlerterInterface struct {
	mock.Mock
}

// SendAlert provides a mock function with given fields: key, alert
func (_m *AlerterInterface) SendAlert(key string, alert alerter.Alert) {
	_m.Called(key, alert)
}
