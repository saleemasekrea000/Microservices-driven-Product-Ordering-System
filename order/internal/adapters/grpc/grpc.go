package grpc

import (
	"context"
	"log"
	"order/internal/application/core/domain"
	"order/proto_gen"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (a *GrpcAdapter) CreateOrder(ctx context.Context, request *proto_gen.CreateOrderRequest) (*proto_gen.CreateOrderResponse, error) {
	if request == nil {
		return nil, status.Error(codes.InvalidArgument, "request is nil")
	}
	var orderItems []*domain.OrderItem
	for _, v := range request.GetOrderItems() {
		orderItems = append(orderItems, &domain.OrderItem{
			OrderID:   v.OrderId,
			ProductID: v.ProductId,
			Amount:    v.Amount,
			Price:     v.Price,
		})
	}
	order := domain.NewOrder(request.Order.TotalPrice, request.Order.UserId, orderItems)
	err := a.apiPort.CreateOrder(ctx, *order)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &proto_gen.CreateOrderResponse{}, nil

}

func (a *GrpcAdapter) GetOrder(ctx context.Context, request *proto_gen.GetOrderRequest) (*proto_gen.GetOrderResponse, error) {
	log.Println("request", request)
	if request == nil {
		log.Println("request is nil")
		return nil, status.Error(codes.InvalidArgument, "request is nil")
	}
	order, err := a.apiPort.GetOrder(ctx, request.OrderId)
	if err != nil {
		log.Println("err", err)
		return nil, status.Error(codes.Internal, err.Error())
	}
	if order == nil {
		return nil, status.Error(codes.NotFound, "order not found")
	}
	orderItemGrpc := []*proto_gen.OrderItem{}
	for _, v := range order.OrderItems {
		orderItemGrpc = append(orderItemGrpc, &proto_gen.OrderItem{
			OrderId:   v.OrderID,
			ProductId: v.ProductID,
			Amount:    v.Amount,
			Price:     v.Price,
		})
	}
	log.Println("orderItemGrpc seccsessfef")
	return &proto_gen.GetOrderResponse{Order: &proto_gen.Order{
		TotalPrice: order.TotalPrice,
		UserId:     order.UserID,
		Id:         order.ID,
		CreatedAt:  timestamppb.New(order.CreatedAt),
		UpdatedAt:  timestamppb.New(order.UpdatedAt),
		OrderItems: orderItemGrpc,
	}}, nil
}
func (a *GrpcAdapter) GetOrders(ctx context.Context, request *proto_gen.GetOrdersRequest) (*proto_gen.GetOrdersResponse, error) {
	orders, err := a.apiPort.GetOrders(ctx, request.Limit, request.Offset)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	orderGrpc := []*proto_gen.Order{}
	for _, v := range orders {
		orderItemGrpc := []*proto_gen.OrderItem{}
		for _, v := range v.OrderItems {
			orderItemGrpc = append(orderItemGrpc, &proto_gen.OrderItem{
				OrderId:   v.OrderID,
				ProductId: v.ProductID,
				Amount:    v.Amount,
				Price:     v.Price,
			})
		}
		orderGrpc = append(orderGrpc, &proto_gen.Order{
			TotalPrice: v.TotalPrice,
			UserId:     v.UserID,
			Id:         v.ID,
			CreatedAt:  timestamppb.New(v.CreatedAt),
			UpdatedAt:  timestamppb.New(v.UpdatedAt),
			OrderItems: orderItemGrpc,
		})
	}
	return &proto_gen.GetOrdersResponse{Orders: orderGrpc}, nil
}
