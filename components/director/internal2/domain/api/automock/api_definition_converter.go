// Code generated by mockery v1.0.0. DO NOT EDIT.

package automock

import (
	api "github.com/kyma-incubator/compass/components/director/internal2/domain/api"
	mock "github.com/stretchr/testify/mock"

	model "github.com/kyma-incubator/compass/components/director/internal2/model"
)

// APIDefinitionConverter is an autogenerated mock type for the APIDefinitionConverter type
type APIDefinitionConverter struct {
	mock.Mock
}

// FromEntity provides a mock function with given fields: entity
func (_m *APIDefinitionConverter) FromEntity(entity api.Entity) model.APIDefinition {
	ret := _m.Called(entity)

	var r0 model.APIDefinition
	if rf, ok := ret.Get(0).(func(api.Entity) model.APIDefinition); ok {
		r0 = rf(entity)
	} else {
		r0 = ret.Get(0).(model.APIDefinition)
	}

	return r0
}

// ToEntity provides a mock function with given fields: apiModel
func (_m *APIDefinitionConverter) ToEntity(apiModel model.APIDefinition) api.Entity {
	ret := _m.Called(apiModel)

	var r0 api.Entity
	if rf, ok := ret.Get(0).(func(model.APIDefinition) api.Entity); ok {
		r0 = rf(apiModel)
	} else {
		r0 = ret.Get(0).(api.Entity)
	}

	return r0
}
