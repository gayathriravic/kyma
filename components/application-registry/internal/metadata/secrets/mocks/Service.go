// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import apperrors "github.com/kyma-project/kyma/components/application-registry/internal/apperrors"
import applications "github.com/kyma-project/kyma/components/application-registry/internal/metadata/applications"
import mock "github.com/stretchr/testify/mock"
import model "github.com/kyma-project/kyma/components/application-registry/internal/metadata/model"

// Service is an autogenerated mock type for the Service type
type Service struct {
	mock.Mock
}

// Create provides a mock function with given fields: application, serviceID, credentials
func (_m *Service) Create(application string, serviceID string, credentials *model.Credentials) (applications.Credentials, apperrors.AppError) {
	ret := _m.Called(application, serviceID, credentials)

	var r0 applications.Credentials
	if rf, ok := ret.Get(0).(func(string, string, *model.Credentials) applications.Credentials); ok {
		r0 = rf(application, serviceID, credentials)
	} else {
		r0 = ret.Get(0).(applications.Credentials)
	}

	var r1 apperrors.AppError
	if rf, ok := ret.Get(1).(func(string, string, *model.Credentials) apperrors.AppError); ok {
		r1 = rf(application, serviceID, credentials)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(apperrors.AppError)
		}
	}

	return r0, r1
}

// Delete provides a mock function with given fields: name
func (_m *Service) Delete(name string) apperrors.AppError {
	ret := _m.Called(name)

	var r0 apperrors.AppError
	if rf, ok := ret.Get(0).(func(string) apperrors.AppError); ok {
		r0 = rf(name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(apperrors.AppError)
		}
	}

	return r0
}

// Get provides a mock function with given fields: application, credentials
func (_m *Service) Get(application string, credentials applications.Credentials) (model.Credentials, apperrors.AppError) {
	ret := _m.Called(application, credentials)

	var r0 model.Credentials
	if rf, ok := ret.Get(0).(func(string, applications.Credentials) model.Credentials); ok {
		r0 = rf(application, credentials)
	} else {
		r0 = ret.Get(0).(model.Credentials)
	}

	var r1 apperrors.AppError
	if rf, ok := ret.Get(1).(func(string, applications.Credentials) apperrors.AppError); ok {
		r1 = rf(application, credentials)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(apperrors.AppError)
		}
	}

	return r0, r1
}

// Update provides a mock function with given fields: application, serviceID, credentials
func (_m *Service) Update(application string, serviceID string, credentials *model.Credentials) (applications.Credentials, apperrors.AppError) {
	ret := _m.Called(application, serviceID, credentials)

	var r0 applications.Credentials
	if rf, ok := ret.Get(0).(func(string, string, *model.Credentials) applications.Credentials); ok {
		r0 = rf(application, serviceID, credentials)
	} else {
		r0 = ret.Get(0).(applications.Credentials)
	}

	var r1 apperrors.AppError
	if rf, ok := ret.Get(1).(func(string, string, *model.Credentials) apperrors.AppError); ok {
		r1 = rf(application, serviceID, credentials)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(apperrors.AppError)
		}
	}

	return r0, r1
}