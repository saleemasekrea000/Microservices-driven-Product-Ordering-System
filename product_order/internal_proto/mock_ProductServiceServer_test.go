// Code generated by mockery v2.42.2. DO NOT EDIT.

package proto_gen

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// MockProductServiceServer is an autogenerated mock type for the ProductServiceServer type
type MockProductServiceServer struct {
	mock.Mock
}

type MockProductServiceServer_Expecter struct {
	mock *mock.Mock
}

func (_m *MockProductServiceServer) EXPECT() *MockProductServiceServer_Expecter {
	return &MockProductServiceServer_Expecter{mock: &_m.Mock}
}

// CreateProduct provides a mock function with given fields: _a0, _a1
func (_m *MockProductServiceServer) CreateProduct(_a0 context.Context, _a1 *CreateProductRequest) (*CreateProductResponse, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for CreateProduct")
	}

	var r0 *CreateProductResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *CreateProductRequest) (*CreateProductResponse, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *CreateProductRequest) *CreateProductResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*CreateProductResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *CreateProductRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockProductServiceServer_CreateProduct_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateProduct'
type MockProductServiceServer_CreateProduct_Call struct {
	*mock.Call
}

// CreateProduct is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 *CreateProductRequest
func (_e *MockProductServiceServer_Expecter) CreateProduct(_a0 interface{}, _a1 interface{}) *MockProductServiceServer_CreateProduct_Call {
	return &MockProductServiceServer_CreateProduct_Call{Call: _e.mock.On("CreateProduct", _a0, _a1)}
}

func (_c *MockProductServiceServer_CreateProduct_Call) Run(run func(_a0 context.Context, _a1 *CreateProductRequest)) *MockProductServiceServer_CreateProduct_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*CreateProductRequest))
	})
	return _c
}

func (_c *MockProductServiceServer_CreateProduct_Call) Return(_a0 *CreateProductResponse, _a1 error) *MockProductServiceServer_CreateProduct_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockProductServiceServer_CreateProduct_Call) RunAndReturn(run func(context.Context, *CreateProductRequest) (*CreateProductResponse, error)) *MockProductServiceServer_CreateProduct_Call {
	_c.Call.Return(run)
	return _c
}

// GetProduct provides a mock function with given fields: _a0, _a1
func (_m *MockProductServiceServer) GetProduct(_a0 context.Context, _a1 *GetProductRequest) (*GetProductResponse, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for GetProduct")
	}

	var r0 *GetProductResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *GetProductRequest) (*GetProductResponse, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *GetProductRequest) *GetProductResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*GetProductResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *GetProductRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockProductServiceServer_GetProduct_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetProduct'
type MockProductServiceServer_GetProduct_Call struct {
	*mock.Call
}

// GetProduct is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 *GetProductRequest
func (_e *MockProductServiceServer_Expecter) GetProduct(_a0 interface{}, _a1 interface{}) *MockProductServiceServer_GetProduct_Call {
	return &MockProductServiceServer_GetProduct_Call{Call: _e.mock.On("GetProduct", _a0, _a1)}
}

func (_c *MockProductServiceServer_GetProduct_Call) Run(run func(_a0 context.Context, _a1 *GetProductRequest)) *MockProductServiceServer_GetProduct_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*GetProductRequest))
	})
	return _c
}

func (_c *MockProductServiceServer_GetProduct_Call) Return(_a0 *GetProductResponse, _a1 error) *MockProductServiceServer_GetProduct_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockProductServiceServer_GetProduct_Call) RunAndReturn(run func(context.Context, *GetProductRequest) (*GetProductResponse, error)) *MockProductServiceServer_GetProduct_Call {
	_c.Call.Return(run)
	return _c
}

// GetProducts provides a mock function with given fields: _a0, _a1
func (_m *MockProductServiceServer) GetProducts(_a0 context.Context, _a1 *GetProductsRequest) (*GetProductsResponse, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for GetProducts")
	}

	var r0 *GetProductsResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *GetProductsRequest) (*GetProductsResponse, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *GetProductsRequest) *GetProductsResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*GetProductsResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *GetProductsRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockProductServiceServer_GetProducts_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetProducts'
type MockProductServiceServer_GetProducts_Call struct {
	*mock.Call
}

// GetProducts is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 *GetProductsRequest
func (_e *MockProductServiceServer_Expecter) GetProducts(_a0 interface{}, _a1 interface{}) *MockProductServiceServer_GetProducts_Call {
	return &MockProductServiceServer_GetProducts_Call{Call: _e.mock.On("GetProducts", _a0, _a1)}
}

func (_c *MockProductServiceServer_GetProducts_Call) Run(run func(_a0 context.Context, _a1 *GetProductsRequest)) *MockProductServiceServer_GetProducts_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*GetProductsRequest))
	})
	return _c
}

func (_c *MockProductServiceServer_GetProducts_Call) Return(_a0 *GetProductsResponse, _a1 error) *MockProductServiceServer_GetProducts_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockProductServiceServer_GetProducts_Call) RunAndReturn(run func(context.Context, *GetProductsRequest) (*GetProductsResponse, error)) *MockProductServiceServer_GetProducts_Call {
	_c.Call.Return(run)
	return _c
}

// mustEmbedUnimplementedProductServiceServer provides a mock function with given fields:
func (_m *MockProductServiceServer) mustEmbedUnimplementedProductServiceServer() {
	_m.Called()
}

// MockProductServiceServer_mustEmbedUnimplementedProductServiceServer_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'mustEmbedUnimplementedProductServiceServer'
type MockProductServiceServer_mustEmbedUnimplementedProductServiceServer_Call struct {
	*mock.Call
}

// mustEmbedUnimplementedProductServiceServer is a helper method to define mock.On call
func (_e *MockProductServiceServer_Expecter) mustEmbedUnimplementedProductServiceServer() *MockProductServiceServer_mustEmbedUnimplementedProductServiceServer_Call {
	return &MockProductServiceServer_mustEmbedUnimplementedProductServiceServer_Call{Call: _e.mock.On("mustEmbedUnimplementedProductServiceServer")}
}

func (_c *MockProductServiceServer_mustEmbedUnimplementedProductServiceServer_Call) Run(run func()) *MockProductServiceServer_mustEmbedUnimplementedProductServiceServer_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockProductServiceServer_mustEmbedUnimplementedProductServiceServer_Call) Return() *MockProductServiceServer_mustEmbedUnimplementedProductServiceServer_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockProductServiceServer_mustEmbedUnimplementedProductServiceServer_Call) RunAndReturn(run func()) *MockProductServiceServer_mustEmbedUnimplementedProductServiceServer_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockProductServiceServer creates a new instance of MockProductServiceServer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockProductServiceServer(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockProductServiceServer {
	mock := &MockProductServiceServer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}