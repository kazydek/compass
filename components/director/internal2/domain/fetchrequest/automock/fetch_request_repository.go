// Code generated by mockery v1.0.0. DO NOT EDIT.

package automock

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	model "github.com/kyma-incubator/compass/components/director/internal2/model"
)

// FetchRequestRepository is an autogenerated mock type for the FetchRequestRepository type
type FetchRequestRepository struct {
	mock.Mock
}

// Update provides a mock function with given fields: ctx, item
func (_m *FetchRequestRepository) Update(ctx context.Context, item *model.FetchRequest) error {
	ret := _m.Called(ctx, item)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *model.FetchRequest) error); ok {
		r0 = rf(ctx, item)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
