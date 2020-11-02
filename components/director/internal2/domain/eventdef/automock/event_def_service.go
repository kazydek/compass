// Code generated by mockery v1.0.0. DO NOT EDIT.

package automock

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	model "github.com/kyma-incubator/compass/components/director/internal2/model"
)

// EventDefService is an autogenerated mock type for the EventDefService type
type EventDefService struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, applicationID, in
func (_m *EventDefService) Create(ctx context.Context, applicationID string, in model.EventDefinitionInput) (string, error) {
	ret := _m.Called(ctx, applicationID, in)

	var r0 string
	if rf, ok := ret.Get(0).(func(context.Context, string, model.EventDefinitionInput) string); ok {
		r0 = rf(ctx, applicationID, in)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, model.EventDefinitionInput) error); ok {
		r1 = rf(ctx, applicationID, in)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateInPackage provides a mock function with given fields: ctx, packageID, in
func (_m *EventDefService) CreateInPackage(ctx context.Context, packageID string, in model.EventDefinitionInput) (string, error) {
	ret := _m.Called(ctx, packageID, in)

	var r0 string
	if rf, ok := ret.Get(0).(func(context.Context, string, model.EventDefinitionInput) string); ok {
		r0 = rf(ctx, packageID, in)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, model.EventDefinitionInput) error); ok {
		r1 = rf(ctx, packageID, in)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: ctx, id
func (_m *EventDefService) Delete(ctx context.Context, id string) error {
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
func (_m *EventDefService) Get(ctx context.Context, id string) (*model.EventDefinition, error) {
	ret := _m.Called(ctx, id)

	var r0 *model.EventDefinition
	if rf, ok := ret.Get(0).(func(context.Context, string) *model.EventDefinition); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.EventDefinition)
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

// GetFetchRequest provides a mock function with given fields: ctx, eventAPIDefID
func (_m *EventDefService) GetFetchRequest(ctx context.Context, eventAPIDefID string) (*model.FetchRequest, error) {
	ret := _m.Called(ctx, eventAPIDefID)

	var r0 *model.FetchRequest
	if rf, ok := ret.Get(0).(func(context.Context, string) *model.FetchRequest); ok {
		r0 = rf(ctx, eventAPIDefID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.FetchRequest)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, eventAPIDefID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RefetchAPISpec provides a mock function with given fields: ctx, id
func (_m *EventDefService) RefetchAPISpec(ctx context.Context, id string) (*model.EventSpec, error) {
	ret := _m.Called(ctx, id)

	var r0 *model.EventSpec
	if rf, ok := ret.Get(0).(func(context.Context, string) *model.EventSpec); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.EventSpec)
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

// Update provides a mock function with given fields: ctx, id, in
func (_m *EventDefService) Update(ctx context.Context, id string, in model.EventDefinitionInput) error {
	ret := _m.Called(ctx, id, in)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, model.EventDefinitionInput) error); ok {
		r0 = rf(ctx, id, in)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
