package grpc

import (
	"context"
	"log"
	"net"
	"product/internal/application/core/domain"
	"product/mocks"
	"product/proto_gen"
	"testing"
	"time"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/test/bufconn"
)

type GrpcSuite struct {
	suite.Suite
	client  proto_gen.ProductServiceClient
	adapter *GrpcAdapter
}

func TestGrpcSuite(t *testing.T) {
	suite.Run(t, new(GrpcSuite))
}

func (gs *GrpcSuite) SetupSuite() {
	lis := bufconn.Listen(1024 * 1024)
	gs.Suite.T().Cleanup(func() {

		_ = lis.Close()
	})
	srv := grpc.NewServer()
	gs.Suite.T().Cleanup(func() {

		srv.Stop()
	})
	gs.adapter = New(0, nil)
	proto_gen.RegisterProductServiceServer(srv, gs.adapter)

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
	gs.T().Cleanup(func() {

		_ = conn.Close()
	})
	if err != nil {
		gs.T().Fatalf("Failed to dial bufnet: %v", err)
	}
	gs.client = proto_gen.NewProductServiceClient(conn)

}

func (gs *GrpcSuite) TestGetProducts() {

	mockApi := mocks.NewApiPort(gs.T())
	mockApi.On("ListProducts", mock.Anything, int32(0), int32(0)).Return([]domain.Product{}, nil)
	gs.adapter.ApiPort = mockApi
	product, err := gs.client.GetProducts(context.Background(), &proto_gen.GetProductsRequest{
		Limit:  0,
		Offset: 0,
	})
	gs.Suite.Nil(err)
	gs.Suite.Equal(0, len(product.Products))

}
func (gs *GrpcSuite) TestCreateProduct() {

	mockApi := mocks.NewApiPort(gs.T())
	mockApi.On("CreateProduct", mock.Anything, mock.Anything).Return(domain.Product{ID: 1, Name: "test", Price: 100, Image: "test", CreatedAt: time.Now(), UpdatedAt: time.Now()}, nil)
	gs.adapter.ApiPort = mockApi
	product, err := gs.client.CreateProduct(context.Background(), &proto_gen.CreateProductRequest{
		Name:  "test",
		Image: "test",
		Price: 100,
	})
	gs.Nil(err)
	gs.Equal(int32(1), product.Product.Id)
	gs.Equal("test", product.Product.Name)
	gs.Equal(float64(100), product.Product.Price)
	gs.Equal("test", product.Product.Image)
	gs.NotNil(product.Product.CreatedAt)
	gs.NotNil(product.Product.UpdatedAt)
}

func (gs *GrpcSuite) TestGetProductById() {

	mockApi := mocks.NewApiPort(gs.T())
	mockApi.On("GetProductById", mock.Anything, int32(1)).Return(domain.Product{ID: 1, Name: "test", Price: 100, Image: "test", CreatedAt: time.Now(), UpdatedAt: time.Now()}, nil)
	gs.adapter.ApiPort = mockApi
	product, err := gs.client.GetProduct(context.Background(), &proto_gen.GetProductRequest{
		Id: 1,
	})
	gs.Nil(err)
	gs.Equal(int32(1), product.Product.Id)
	gs.Equal("test", product.Product.Name)
	gs.Equal(float64(100), product.Product.Price)
	gs.Equal("test", product.Product.Image)
	gs.NotNil(product.Product.CreatedAt)
	gs.NotNil(product.Product.UpdatedAt)
}
