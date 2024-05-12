package ports

import (
	"context"
	"product_order/internal/application/core/domains"
)

type ApiPort interface {
	GetFullOrder(ctx context.Context, orderID int32) (domains.Order, error)
}
