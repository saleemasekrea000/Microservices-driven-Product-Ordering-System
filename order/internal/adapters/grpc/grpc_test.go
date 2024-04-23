package grpc

import (
	"context"
	"log"
	"net"

	// "order/internal/application/core/domain"
	"order/internal/application/core/domain"
	"order/mocks"
	"order/proto_gen"
	"testing"

	// "order/mocks"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/test/bufconn"
)

type GrpcTestSuite struct {
	adatapter *GrpcAdapter
	suite.Suite
	client proto_gen.OrderServiceClient
}

func TestGrpcTestSuite(t *testing.T) {
	suite.Run(t, new(GrpcTestSuite))
}

func (s *GrpcTestSuite) SetupSuite() {
	lis := bufconn.Listen(1024 * 1024)

	s.T().Cleanup(func() { _ = lis.Close() })

	srv := grpc.NewServer()
	s.T().Cleanup(func() { srv.Stop() })

	s.adatapter = NewAdatpter(0, nil)
	proto_gen.RegisterOrderServiceServer(srv, s.adatapter)
	go func() {
		err := srv.Serve(lis)
		if err != nil {
			log.Fatalf("failed to serve: %v", err)

		}
	}()

	bufDialer := func(context.Context, string) (net.Conn, error) {
		return lis.Dial()
	}

	conn, err := grpc.DialContext(context.Background(), "", grpc.WithContextDialer(bufDialer), grpc.WithTransportCredentials(insecure.NewCredentials()))
	s.T().Cleanup(func() { _ = conn.Close() })
	if err != nil {
		s.Fail(err.Error())
	}

	s.client = proto_gen.NewOrderServiceClient(conn)

}

// 	GetOrder(ctx context.Context, orderID int32) (*domain.Order, error)

func (s *GrpcTestSuite) TestGetOrder() {

	apiMock := mocks.NewAPIPort(s.T())
	// returnedOrder := &domain.Order{ID: 1, TotalPrice: 1, UserID: 1, OrderItems: nil}
	apiMock.On("GetOrder", mock.Anything, int32(1)).Return(&domain.Order{ID: 1, TotalPrice: 1, UserID: 1, OrderItems: nil}, nil)

	s.adatapter.apiPort = apiMock

	order, err := s.client.GetOrder(context.Background(), &proto_gen.GetOrderRequest{OrderId: 1})
	s.NoError(err)
	s.Equal(order.Order.Id, int32(1))
}

// 	CreateOrder(ctx context.Context,order domain.Order) error

func (s *GrpcTestSuite) TestCreateOrder() {

	apiMock := mocks.NewAPIPort(s.T())
	apiMock.On("CreateOrder", mock.Anything, mock.Anything).Return(nil)
	s.adatapter.apiPort = apiMock
	_, err := s.client.CreateOrder(context.Background(), &proto_gen.CreateOrderRequest{Order: &proto_gen.Order{Id: 1, TotalPrice: 1, UserId: 1}, OrderItems: []*proto_gen.OrderItem{}})
	s.NoError(err)
}

// GetOrders(ctx context.Context,limit int32, offset int32) ([]domain.Order, error)

func (s *GrpcTestSuite) TestGetOrders() {

	apiMock := mocks.NewAPIPort(s.T())
	apiMock.On("GetOrders", mock.Anything, int32(1), int32(1)).Return([]domain.Order{{ID: 1, TotalPrice: 1, UserID: 1, OrderItems: nil}}, nil)
	s.adatapter.apiPort = apiMock
	orders, err := s.client.GetOrders(context.Background(), &proto_gen.GetOrdersRequest{Limit: 1, Offset: 1})
	s.NoError(err)
	s.Equal(orders.Orders[0].Id, int32(1))
}
