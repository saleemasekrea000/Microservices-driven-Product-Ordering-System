package ports

import (
	"context"
	"order/internal/adapters/db/generated"
)

type DBPort interface {
	CreateOrder(ctx context.Context, arg generated.CreateOrderParams) error
	CreateOrderItems(ctx context.Context, arg generated.CreateOrderItemsParams) error
	DeleteOrder(ctx context.Context, id int32) error
	DeleteOrderItems(ctx context.Context, id int32) error
	GetOrder(ctx context.Context, id int32) ([]generated.GetOrderRow, error)
	GetOrderItems(ctx context.Context, id int32) ([]generated.OrderItem, error)
	GetOrders(ctx context.Context, arg generated.GetOrdersParams) ([]generated.GetOrdersRow, error)
	UpdateOrder(ctx context.Context, arg generated.UpdateOrderParams) error
	UpdateOrderItems(ctx context.Context, arg generated.UpdateOrderItemsParams) error
}
