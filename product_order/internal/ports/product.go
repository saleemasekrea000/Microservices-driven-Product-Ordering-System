package ports

import (
	"context"
	"product_order/internal/application/core/domains"
)

type ProductPort interface {
	GetProduct(ctx context.Context, id int32) (domains.Product, error)
}
