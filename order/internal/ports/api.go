package ports

import (
	"context"
	"order/internal/application/core/domain"
)

type APIPort interface {
	CreateOrder(ctx context.Context, order domain.Order) error
	GetOrders(ctx context.Context, limit int32, offset int32) ([]domain.Order, error)
	GetOrder(ctx context.Context, orderID int32) (*domain.Order, error)
}
