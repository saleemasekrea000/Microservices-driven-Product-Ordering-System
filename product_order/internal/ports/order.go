package ports

import (
	"context"
	"product_order/internal/application/core/domains"
)

//go:generate mockery --name OrderPort
type OrderPort interface {
	GetOrder(ctx context.Context, id int32) (domains.Order, error)
}
