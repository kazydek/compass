// Code generated by mockery v1.0.0. DO NOT EDIT.

package automock

import (
	context "context"

	persistence "github.com/kyma-incubator/compass/components/director/pkg/persistence"
	mock "github.com/stretchr/testify/mock"

	sql "database/sql"
)

// Transactioner is an autogenerated mock type for the Transactioner type
type Transactioner struct {
	mock.Mock
}

// Begin provides a mock function with given fields:
func (_m *Transactioner) Begin() (persistence.PersistenceTx, error) {
	ret := _m.Called()

	var r0 persistence.PersistenceTx
	if rf, ok := ret.Get(0).(func() persistence.PersistenceTx); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(persistence.PersistenceTx)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PingContext provides a mock function with given fields: ctx
func (_m *Transactioner) PingContext(ctx context.Context) error {
	ret := _m.Called(ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RollbackUnlessCommitted provides a mock function with given fields: tx
func (_m *Transactioner) RollbackUnlessCommitted(tx persistence.PersistenceTx) {
	_m.Called(tx)
}

// Stats provides a mock function with given fields:
func (_m *Transactioner) Stats() sql.DBStats {
	ret := _m.Called()

	var r0 sql.DBStats
	if rf, ok := ret.Get(0).(func() sql.DBStats); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(sql.DBStats)
	}

	return r0
}
