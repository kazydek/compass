// Code generated by mockery v1.0.0. DO NOT EDIT.

package automock

import context "context"
import mock "github.com/stretchr/testify/mock"
import model "github.com/kyma-incubator/compass/components/director/internal2/model"

// Service is an autogenerated mock type for the Service type
type Service struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, packageID, in, defaultAuth, requestInputSchema
func (_m *Service) Create(ctx context.Context, packageID string, in model.PackageInstanceAuthRequestInput, defaultAuth *model.Auth, requestInputSchema *string) (string, error) {
	ret := _m.Called(ctx, packageID, in, defaultAuth, requestInputSchema)

	var r0 string
	if rf, ok := ret.Get(0).(func(context.Context, string, model.PackageInstanceAuthRequestInput, *model.Auth, *string) string); ok {
		r0 = rf(ctx, packageID, in, defaultAuth, requestInputSchema)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, model.PackageInstanceAuthRequestInput, *model.Auth, *string) error); ok {
		r1 = rf(ctx, packageID, in, defaultAuth, requestInputSchema)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: ctx, id
func (_m *Service) Delete(ctx context.Context, id string) error {
	ret := _m.Called(ctx, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Get provides a mock function with given fields: ctx, id
func (_m *Service) Get(ctx context.Context, id string) (*model.PackageInstanceAuth, error) {
	ret := _m.Called(ctx, id)

	var r0 *model.PackageInstanceAuth
	if rf, ok := ret.Get(0).(func(context.Context, string) *model.PackageInstanceAuth); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.PackageInstanceAuth)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RequestDeletion provides a mock function with given fields: ctx, instanceAuth, defaultPackageInstanceAuth
func (_m *Service) RequestDeletion(ctx context.Context, instanceAuth *model.PackageInstanceAuth, defaultPackageInstanceAuth *model.Auth) (bool, error) {
	ret := _m.Called(ctx, instanceAuth, defaultPackageInstanceAuth)

	var r0 bool
	if rf, ok := ret.Get(0).(func(context.Context, *model.PackageInstanceAuth, *model.Auth) bool); ok {
		r0 = rf(ctx, instanceAuth, defaultPackageInstanceAuth)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *model.PackageInstanceAuth, *model.Auth) error); ok {
		r1 = rf(ctx, instanceAuth, defaultPackageInstanceAuth)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SetAuth provides a mock function with given fields: ctx, id, in
func (_m *Service) SetAuth(ctx context.Context, id string, in model.PackageInstanceAuthSetInput) error {
	ret := _m.Called(ctx, id, in)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, model.PackageInstanceAuthSetInput) error); ok {
		r0 = rf(ctx, id, in)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
