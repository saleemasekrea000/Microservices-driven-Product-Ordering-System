package order

import (
	"context"
	"log"
	"product_order/internal/application/core/domains"
	proto_gen "product_order/internal_proto"
	"time"

	grpc_retry "github.com/grpc-ecosystem/go-grpc-middleware/retry"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
)

type OrderAdtapter struct {
	client proto_gen.OrderServiceClient
}

func New(order_url string) (*OrderAdtapter, error) {
	var opts []grpc.DialOption
	opts = append(opts,
		grpc.WithUnaryInterceptor(grpc_retry.UnaryClientInterceptor(
			grpc_retry.WithCodes(codes.Unavailable, codes.ResourceExhausted),
			grpc_retry.WithMax(5),
			grpc_retry.WithBackoff(grpc_retry.BackoffLinear(time.Second)),
		)))
	opts = append(opts,
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	conn, err := grpc.Dial(order_url, opts...)
	if err != nil {
		return nil, err
	}
	cleint := proto_gen.NewOrderServiceClient(conn)
	return &OrderAdtapter{client: cleint}, nil
}

func (a *OrderAdtapter) GetOrder(ctx context.Context, id int32) (domains.Order, error) {
	res, err := a.client.GetOrder(ctx, &proto_gen.GetOrderRequest{OrderId: id})
	if err != nil {
		log.Println("err hone", err)
		return domains.Order{}, err
	}
	log.Println("pass")
	var orderItems []*domains.OrderItem
	for _, item := range res.Order.GetOrderItems() {
		log.Println("orderItem", item.Id)
		orderItems = append(orderItems, &domains.OrderItem{
			OrderID: item.OrderId,
			Amount:  item.Amount,
			Price:   item.Price,
			Product: &domains.Product{
				ID: item.Id,
			},
		})
	}
	order := domains.Order{
		ID:         res.Order.Id,
		TotalPrice: res.Order.TotalPrice,
		UserID:     res.Order.UserId,
		OrderItems: orderItems,
		CreatedAt:  res.Order.CreatedAt.AsTime(),
		UpdatedAt:  res.Order.UpdatedAt.AsTime(),
	}
	return order, nil
}
