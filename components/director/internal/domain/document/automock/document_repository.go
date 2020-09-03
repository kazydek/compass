// Code generated by mockery v1.0.0. DO NOT EDIT.

package automock

import context "context"

import mock "github.com/stretchr/testify/mock"
import model "github.com/kyma-incubator/compass/components/director/internal/model"

// DocumentRepository is an autogenerated mock type for the DocumentRepository type
type DocumentRepository struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, item
func (_m *DocumentRepository) Create(ctx context.Context, item *model.Document) error {
	ret := _m.Called(ctx, item)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *model.Document) error); ok {
		r0 = rf(ctx, item)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Delete provides a mock function with given fields: ctx, tenant, id
func (_m *DocumentRepository) Delete(ctx context.Context, tenant string, id string) error {
	ret := _m.Called(ctx, tenant, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) error); ok {
		r0 = rf(ctx, tenant, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Exists provides a mock function with given fields: ctx, tenant, id
func (_m *DocumentRepository) Exists(ctx context.Context, tenant string, id string) (bool, error) {
	ret := _m.Called(ctx, tenant, id)

	var r0 bool
	if rf, ok := ret.Get(0).(func(context.Context, string, string) bool); ok {
		r0 = rf(ctx, tenant, id)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, tenant, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByID provides a mock function with given fields: ctx, tenant, id
func (_m *DocumentRepository) GetByID(ctx context.Context, tenant string, id string) (*model.Document, error) {
	ret := _m.Called(ctx, tenant, id)

	var r0 *model.Document
	if rf, ok := ret.Get(0).(func(context.Context, string, string) *model.Document); ok {
		r0 = rf(ctx, tenant, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Document)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, tenant, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetForBundle provides a mock function with given fields: ctx, tenant, id, bundleID
func (_m *DocumentRepository) GetForBundle(ctx context.Context, tenant string, id string, bundleID string) (*model.Document, error) {
	ret := _m.Called(ctx, tenant, id, bundleID)

	var r0 *model.Document
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string) *model.Document); ok {
		r0 = rf(ctx, tenant, id, bundleID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Document)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string, string) error); ok {
		r1 = rf(ctx, tenant, id, bundleID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListForApplication provides a mock function with given fields: ctx, tenant, applicationID, pageSize, cursor
func (_m *DocumentRepository) ListForApplication(ctx context.Context, tenant string, applicationID string, pageSize int, cursor string) (*model.DocumentPage, error) {
	ret := _m.Called(ctx, tenant, applicationID, pageSize, cursor)

	var r0 *model.DocumentPage
	if rf, ok := ret.Get(0).(func(context.Context, string, string, int, string) *model.DocumentPage); ok {
		r0 = rf(ctx, tenant, applicationID, pageSize, cursor)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.DocumentPage)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string, int, string) error); ok {
		r1 = rf(ctx, tenant, applicationID, pageSize, cursor)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListForBundle provides a mock function with given fields: ctx, tenant, bundleID, pageSize, cursor
func (_m *DocumentRepository) ListForBundle(ctx context.Context, tenant string, bundleID string, pageSize int, cursor string) (*model.DocumentPage, error) {
	ret := _m.Called(ctx, tenant, bundleID, pageSize, cursor)

	var r0 *model.DocumentPage
	if rf, ok := ret.Get(0).(func(context.Context, string, string, int, string) *model.DocumentPage); ok {
		r0 = rf(ctx, tenant, bundleID, pageSize, cursor)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.DocumentPage)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string, int, string) error); ok {
		r1 = rf(ctx, tenant, bundleID, pageSize, cursor)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
