// Code generated by mockery v2.42.2. DO NOT EDIT.

package mocks

import (
	context "context"
	domain "order/internal/application/core/domain"

	mock "github.com/stretchr/testify/mock"
)

// APIPort is an autogenerated mock type for the APIPort type
type APIPort struct {
	mock.Mock
}

// CreateOrder provides a mock function with given fields: ctx, order
func (_m *APIPort) CreateOrder(ctx context.Context, order domain.Order) error {
	ret := _m.Called(ctx, order)

	if len(ret) == 0 {
		panic("no return value specified for CreateOrder")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, domain.Order) error); ok {
		r0 = rf(ctx, order)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetOrder provides a mock function with given fields: ctx, orderID
func (_m *APIPort) GetOrder(ctx context.Context, orderID int32) (*domain.Order, error) {
	ret := _m.Called(ctx, orderID)

	if len(ret) == 0 {
		panic("no return value specified for GetOrder")
	}

	var r0 *domain.Order
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int32) (*domain.Order, error)); ok {
		return rf(ctx, orderID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int32) *domain.Order); ok {
		r0 = rf(ctx, orderID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Order)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int32) error); ok {
		r1 = rf(ctx, orderID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetOrders provides a mock function with given fields: ctx, limit, offset
func (_m *APIPort) GetOrders(ctx context.Context, limit int32, offset int32) ([]domain.Order, error) {
	ret := _m.Called(ctx, limit, offset)

	if len(ret) == 0 {
		panic("no return value specified for GetOrders")
	}

	var r0 []domain.Order
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int32, int32) ([]domain.Order, error)); ok {
		return rf(ctx, limit, offset)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int32, int32) []domain.Order); ok {
		r0 = rf(ctx, limit, offset)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.Order)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int32, int32) error); ok {
		r1 = rf(ctx, limit, offset)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewAPIPort creates a new instance of APIPort. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewAPIPort(t interface {
	mock.TestingT
	Cleanup(func())
}) *APIPort {
	mock := &APIPort{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
