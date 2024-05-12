package grpc

import (
	"context"
	"product/internal/application/core/domain"
	"product/proto_gen"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (a *GrpcAdapter) CreateProduct(ctx context.Context, request *proto_gen.CreateProductRequest) (*proto_gen.CreateProductResponse, error) {

	if request == nil {
		return nil, status.Error(codes.InvalidArgument, "request is nil")
	}
	product := domain.NewProduct(request.Name, request.Price, request.Image)
	product, err := a.ApiPort.CreateProduct(ctx, product)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &proto_gen.CreateProductResponse{Product: &proto_gen.Product{
		Id:        product.ID,
		Name:      product.Name,
		Price:     product.Price,
		Image:     product.Image,
		CreatedAt: timestamppb.New(product.CreatedAt),
		UpdatedAt: timestamppb.New(product.UpdatedAt),
	}}, nil
}
func (a *GrpcAdapter) GetProduct(ctx context.Context, request *proto_gen.GetProductRequest) (*proto_gen.GetProductResponse, error) {

	if request == nil {
		return nil, status.Error(codes.InvalidArgument, "request is nil")
	}
	product, err := a.ApiPort.GetProductById(ctx, request.Id)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &proto_gen.GetProductResponse{Product: &proto_gen.Product{
		Id:        product.ID,
		Name:      product.Name,
		Price:     product.Price,
		Image:     product.Image,
		CreatedAt: timestamppb.New(product.CreatedAt),
		UpdatedAt: timestamppb.New(product.UpdatedAt),
	}}, nil
}

func (a *GrpcAdapter) GetProducts(ctx context.Context, requesr *proto_gen.GetProductsRequest) (*proto_gen.GetProductsResponse, error) {

	products, err := a.ApiPort.ListProducts(ctx, requesr.Limit, requesr.Offset)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	var productsGrpc []*proto_gen.Product
	for _, v := range products {
		productsGrpc = append(productsGrpc, &proto_gen.Product{
			Id:        v.ID,
			Name:      v.Name,
			Price:     v.Price,
			Image:     v.Image,
			CreatedAt: timestamppb.New(v.CreatedAt),
			UpdatedAt: timestamppb.New(v.UpdatedAt),
		})
	}
	return &proto_gen.GetProductsResponse{Products: productsGrpc}, nil
}
