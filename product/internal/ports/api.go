package ports

import (
	"context"
	"product/internal/application/core/domain"
)

type ApiPort interface {
	CreateProduct(ctx context.Context, product domain.Product) (domain.Product, error)
	GetProductById(ctx context.Context, id int32) (domain.Product, error)
	ListProducts(ctx context.Context, limit int32, offset int32) ([]domain.Product, error)
}
