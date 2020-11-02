// Code generated by mockery v1.0.0. DO NOT EDIT.

package automock

import integrationsystem "github.com/kyma-incubator/compass/components/director/internal2/domain/integrationsystem"
import mock "github.com/stretchr/testify/mock"
import model "github.com/kyma-incubator/compass/components/director/internal2/model"

// Converter is an autogenerated mock type for the Converter type
type Converter struct {
	mock.Mock
}

// FromEntity provides a mock function with given fields: in
func (_m *Converter) FromEntity(in *integrationsystem.Entity) *model.IntegrationSystem {
	ret := _m.Called(in)

	var r0 *model.IntegrationSystem
	if rf, ok := ret.Get(0).(func(*integrationsystem.Entity) *model.IntegrationSystem); ok {
		r0 = rf(in)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.IntegrationSystem)
		}
	}

	return r0
}

// ToEntity provides a mock function with given fields: in
func (_m *Converter) ToEntity(in *model.IntegrationSystem) *integrationsystem.Entity {
	ret := _m.Called(in)

	var r0 *integrationsystem.Entity
	if rf, ok := ret.Get(0).(func(*model.IntegrationSystem) *integrationsystem.Entity); ok {
		r0 = rf(in)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*integrationsystem.Entity)
		}
	}

	return r0
}
