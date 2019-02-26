// Code generated by mockery v1.0.0. DO NOT EDIT.
package mocks

import certificates "github.com/kyma-project/kyma/components/connectivity-certs-controller/internal/certificates"

import mock "github.com/stretchr/testify/mock"

// Client is an autogenerated mock type for the Client type
type Client struct {
	mock.Mock
}

// RequestCertificates provides a mock function with given fields: csrInfoURL
func (_m *Client) RequestCertificates(csrInfoURL string) (certificates.Certificates, error) {
	ret := _m.Called(csrInfoURL)

	var r0 certificates.Certificates
	if rf, ok := ret.Get(0).(func(string) certificates.Certificates); ok {
		r0 = rf(csrInfoURL)
	} else {
		r0 = ret.Get(0).(certificates.Certificates)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(csrInfoURL)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}