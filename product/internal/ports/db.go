package ports

import (
	"context"
	"product/internal/adapters/db/generated"
)

type DBPort interface {
	CreateProduct(ctx context.Context, arg generated.CreateProductParams) (generated.Product, error)
	GetProductById(ctx context.Context, id int32) (generated.Product, error)
	ListProducts(ctx context.Context, arg generated.ListProductsParams) ([]generated.Product, error)
}
