// Code generated by mockery v2.42.2. DO NOT EDIT.

package mocks

import (
	context "context"
	generated "product/internal/adapters/db/generated"

	mock "github.com/stretchr/testify/mock"
)

// Querier is an autogenerated mock type for the Querier type
type Querier struct {
	mock.Mock
}

// CreateProduct provides a mock function with given fields: ctx, arg
func (_m *Querier) CreateProduct(ctx context.Context, arg generated.CreateProductParams) (generated.Product, error) {
	ret := _m.Called(ctx, arg)

	if len(ret) == 0 {
		panic("no return value specified for CreateProduct")
	}

	var r0 generated.Product
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, generated.CreateProductParams) (generated.Product, error)); ok {
		return rf(ctx, arg)
	}
	if rf, ok := ret.Get(0).(func(context.Context, generated.CreateProductParams) generated.Product); ok {
		r0 = rf(ctx, arg)
	} else {
		r0 = ret.Get(0).(generated.Product)
	}

	if rf, ok := ret.Get(1).(func(context.Context, generated.CreateProductParams) error); ok {
		r1 = rf(ctx, arg)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetProductById provides a mock function with given fields: ctx, id
func (_m *Querier) GetProductById(ctx context.Context, id int32) (generated.Product, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for GetProductById")
	}

	var r0 generated.Product
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int32) (generated.Product, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int32) generated.Product); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(generated.Product)
	}

	if rf, ok := ret.Get(1).(func(context.Context, int32) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListProducts provides a mock function with given fields: ctx, arg
func (_m *Querier) ListProducts(ctx context.Context, arg generated.ListProductsParams) ([]generated.Product, error) {
	ret := _m.Called(ctx, arg)

	if len(ret) == 0 {
		panic("no return value specified for ListProducts")
	}

	var r0 []generated.Product
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, generated.ListProductsParams) ([]generated.Product, error)); ok {
		return rf(ctx, arg)
	}
	if rf, ok := ret.Get(0).(func(context.Context, generated.ListProductsParams) []generated.Product); ok {
		r0 = rf(ctx, arg)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]generated.Product)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, generated.ListProductsParams) error); ok {
		r1 = rf(ctx, arg)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewQuerier creates a new instance of Querier. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewQuerier(t interface {
	mock.TestingT
	Cleanup(func())
}) *Querier {
	mock := &Querier{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
