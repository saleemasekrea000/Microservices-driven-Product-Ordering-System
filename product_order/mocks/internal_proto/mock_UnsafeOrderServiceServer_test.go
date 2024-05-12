// Code generated by mockery v2.42.2. DO NOT EDIT.

package proto_gen

import mock "github.com/stretchr/testify/mock"

// UnsafeOrderServiceServer is an autogenerated mock type for the UnsafeOrderServiceServer type
type UnsafeOrderServiceServer struct {
	mock.Mock
}

type UnsafeOrderServiceServer_Expecter struct {
	mock *mock.Mock
}

func (_m *UnsafeOrderServiceServer) EXPECT() *UnsafeOrderServiceServer_Expecter {
	return &UnsafeOrderServiceServer_Expecter{mock: &_m.Mock}
}

// mustEmbedUnimplementedOrderServiceServer provides a mock function with given fields:
func (_m *UnsafeOrderServiceServer) mustEmbedUnimplementedOrderServiceServer() {
	_m.Called()
}

// UnsafeOrderServiceServer_mustEmbedUnimplementedOrderServiceServer_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'mustEmbedUnimplementedOrderServiceServer'
type UnsafeOrderServiceServer_mustEmbedUnimplementedOrderServiceServer_Call struct {
	*mock.Call
}

// mustEmbedUnimplementedOrderServiceServer is a helper method to define mock.On call
func (_e *UnsafeOrderServiceServer_Expecter) mustEmbedUnimplementedOrderServiceServer() *UnsafeOrderServiceServer_mustEmbedUnimplementedOrderServiceServer_Call {
	return &UnsafeOrderServiceServer_mustEmbedUnimplementedOrderServiceServer_Call{Call: _e.mock.On("mustEmbedUnimplementedOrderServiceServer")}
}

func (_c *UnsafeOrderServiceServer_mustEmbedUnimplementedOrderServiceServer_Call) Run(run func()) *UnsafeOrderServiceServer_mustEmbedUnimplementedOrderServiceServer_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *UnsafeOrderServiceServer_mustEmbedUnimplementedOrderServiceServer_Call) Return() *UnsafeOrderServiceServer_mustEmbedUnimplementedOrderServiceServer_Call {
	_c.Call.Return()
	return _c
}

func (_c *UnsafeOrderServiceServer_mustEmbedUnimplementedOrderServiceServer_Call) RunAndReturn(run func()) *UnsafeOrderServiceServer_mustEmbedUnimplementedOrderServiceServer_Call {
	_c.Call.Return(run)
	return _c
}

// NewUnsafeOrderServiceServer creates a new instance of UnsafeOrderServiceServer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewUnsafeOrderServiceServer(t interface {
	mock.TestingT
	Cleanup(func())
}) *UnsafeOrderServiceServer {
	mock := &UnsafeOrderServiceServer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}