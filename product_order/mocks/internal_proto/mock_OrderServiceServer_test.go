// Code generated by mockery v2.42.2. DO NOT EDIT.

package proto_gen

import (
	context "context"
	internal_proto "product_order/internal_proto"

	mock "github.com/stretchr/testify/mock"
)

// OrderServiceServer is an autogenerated mock type for the OrderServiceServer type
type OrderServiceServer struct {
	mock.Mock
}

type OrderServiceServer_Expecter struct {
	mock *mock.Mock
}

func (_m *OrderServiceServer) EXPECT() *OrderServiceServer_Expecter {
	return &OrderServiceServer_Expecter{mock: &_m.Mock}
}

// CreateOrder provides a mock function with given fields: _a0, _a1
func (_m *OrderServiceServer) CreateOrder(_a0 context.Context, _a1 *internal_proto.CreateOrderRequest) (*internal_proto.CreateOrderResponse, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for CreateOrder")
	}

	var r0 *internal_proto.CreateOrderResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *internal_proto.CreateOrderRequest) (*internal_proto.CreateOrderResponse, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *internal_proto.CreateOrderRequest) *internal_proto.CreateOrderResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*internal_proto.CreateOrderResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *internal_proto.CreateOrderRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// OrderServiceServer_CreateOrder_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateOrder'
type OrderServiceServer_CreateOrder_Call struct {
	*mock.Call
}

// CreateOrder is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 *internal_proto.CreateOrderRequest
func (_e *OrderServiceServer_Expecter) CreateOrder(_a0 interface{}, _a1 interface{}) *OrderServiceServer_CreateOrder_Call {
	return &OrderServiceServer_CreateOrder_Call{Call: _e.mock.On("CreateOrder", _a0, _a1)}
}

func (_c *OrderServiceServer_CreateOrder_Call) Run(run func(_a0 context.Context, _a1 *internal_proto.CreateOrderRequest)) *OrderServiceServer_CreateOrder_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*internal_proto.CreateOrderRequest))
	})
	return _c
}

func (_c *OrderServiceServer_CreateOrder_Call) Return(_a0 *internal_proto.CreateOrderResponse, _a1 error) *OrderServiceServer_CreateOrder_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *OrderServiceServer_CreateOrder_Call) RunAndReturn(run func(context.Context, *internal_proto.CreateOrderRequest) (*internal_proto.CreateOrderResponse, error)) *OrderServiceServer_CreateOrder_Call {
	_c.Call.Return(run)
	return _c
}

// GetOrder provides a mock function with given fields: _a0, _a1
func (_m *OrderServiceServer) GetOrder(_a0 context.Context, _a1 *internal_proto.GetOrderRequest) (*internal_proto.GetOrderResponse, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for GetOrder")
	}

	var r0 *internal_proto.GetOrderResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *internal_proto.GetOrderRequest) (*internal_proto.GetOrderResponse, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *internal_proto.GetOrderRequest) *internal_proto.GetOrderResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*internal_proto.GetOrderResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *internal_proto.GetOrderRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// OrderServiceServer_GetOrder_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetOrder'
type OrderServiceServer_GetOrder_Call struct {
	*mock.Call
}

// GetOrder is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 *internal_proto.GetOrderRequest
func (_e *OrderServiceServer_Expecter) GetOrder(_a0 interface{}, _a1 interface{}) *OrderServiceServer_GetOrder_Call {
	return &OrderServiceServer_GetOrder_Call{Call: _e.mock.On("GetOrder", _a0, _a1)}
}

func (_c *OrderServiceServer_GetOrder_Call) Run(run func(_a0 context.Context, _a1 *internal_proto.GetOrderRequest)) *OrderServiceServer_GetOrder_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*internal_proto.GetOrderRequest))
	})
	return _c
}

func (_c *OrderServiceServer_GetOrder_Call) Return(_a0 *internal_proto.GetOrderResponse, _a1 error) *OrderServiceServer_GetOrder_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *OrderServiceServer_GetOrder_Call) RunAndReturn(run func(context.Context, *internal_proto.GetOrderRequest) (*internal_proto.GetOrderResponse, error)) *OrderServiceServer_GetOrder_Call {
	_c.Call.Return(run)
	return _c
}

// GetOrders provides a mock function with given fields: _a0, _a1
func (_m *OrderServiceServer) GetOrders(_a0 context.Context, _a1 *internal_proto.GetOrdersRequest) (*internal_proto.GetOrdersResponse, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for GetOrders")
	}

	var r0 *internal_proto.GetOrdersResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *internal_proto.GetOrdersRequest) (*internal_proto.GetOrdersResponse, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *internal_proto.GetOrdersRequest) *internal_proto.GetOrdersResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*internal_proto.GetOrdersResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *internal_proto.GetOrdersRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// OrderServiceServer_GetOrders_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetOrders'
type OrderServiceServer_GetOrders_Call struct {
	*mock.Call
}

// GetOrders is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 *internal_proto.GetOrdersRequest
func (_e *OrderServiceServer_Expecter) GetOrders(_a0 interface{}, _a1 interface{}) *OrderServiceServer_GetOrders_Call {
	return &OrderServiceServer_GetOrders_Call{Call: _e.mock.On("GetOrders", _a0, _a1)}
}

func (_c *OrderServiceServer_GetOrders_Call) Run(run func(_a0 context.Context, _a1 *internal_proto.GetOrdersRequest)) *OrderServiceServer_GetOrders_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*internal_proto.GetOrdersRequest))
	})
	return _c
}

func (_c *OrderServiceServer_GetOrders_Call) Return(_a0 *internal_proto.GetOrdersResponse, _a1 error) *OrderServiceServer_GetOrders_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *OrderServiceServer_GetOrders_Call) RunAndReturn(run func(context.Context, *internal_proto.GetOrdersRequest) (*internal_proto.GetOrdersResponse, error)) *OrderServiceServer_GetOrders_Call {
	_c.Call.Return(run)
	return _c
}

// mustEmbedUnimplementedOrderServiceServer provides a mock function with given fields:
func (_m *OrderServiceServer) mustEmbedUnimplementedOrderServiceServer() {
	_m.Called()
}

// OrderServiceServer_mustEmbedUnimplementedOrderServiceServer_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'mustEmbedUnimplementedOrderServiceServer'
type OrderServiceServer_mustEmbedUnimplementedOrderServiceServer_Call struct {
	*mock.Call
}

// mustEmbedUnimplementedOrderServiceServer is a helper method to define mock.On call
func (_e *OrderServiceServer_Expecter) mustEmbedUnimplementedOrderServiceServer() *OrderServiceServer_mustEmbedUnimplementedOrderServiceServer_Call {
	return &OrderServiceServer_mustEmbedUnimplementedOrderServiceServer_Call{Call: _e.mock.On("mustEmbedUnimplementedOrderServiceServer")}
}

func (_c *OrderServiceServer_mustEmbedUnimplementedOrderServiceServer_Call) Run(run func()) *OrderServiceServer_mustEmbedUnimplementedOrderServiceServer_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *OrderServiceServer_mustEmbedUnimplementedOrderServiceServer_Call) Return() *OrderServiceServer_mustEmbedUnimplementedOrderServiceServer_Call {
	_c.Call.Return()
	return _c
}

func (_c *OrderServiceServer_mustEmbedUnimplementedOrderServiceServer_Call) RunAndReturn(run func()) *OrderServiceServer_mustEmbedUnimplementedOrderServiceServer_Call {
	_c.Call.Return(run)
	return _c
}

// NewOrderServiceServer creates a new instance of OrderServiceServer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewOrderServiceServer(t interface {
	mock.TestingT
	Cleanup(func())
}) *OrderServiceServer {
	mock := &OrderServiceServer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}