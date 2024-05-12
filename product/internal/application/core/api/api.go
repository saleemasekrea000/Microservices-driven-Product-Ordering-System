package api

import (
	"context"
	"product/internal/adapters/db/generated"
	"product/internal/application/core/domain"
	"product/internal/ports"

	"github.com/jackc/pgx/v5/pgtype"
)

type Application struct {
	db ports.DBPort
}

func New(db ports.DBPort) *Application {
	return &Application{
		db: db,
	}
}

func (app *Application) CreateProduct(ctx context.Context, product domain.Product) (domain.Product, error) {

	arg := generated.CreateProductParams{
		Name:  pgtype.Text{String: product.Name, Valid: true},
		Price: pgtype.Float8{Float64: product.Price, Valid: true},
		Image: pgtype.Text{String: product.Image, Valid: true},
	}
	results, err := app.db.CreateProduct(ctx, arg)
	if err != nil {
		return domain.Product{}, err
	}
	return domain.Product{
		ID:        results.ID,
		Name:      results.Name.String,
		Price:     results.Price.Float64,
		Image:     results.Image.String,
		CreatedAt: results.CreatedAt.Time,
		UpdatedAt: results.UpdatedAt.Time,
	}, nil
}

func (app *Application) GetProductById(ctx context.Context, id int32) (domain.Product, error) {

	results, err := app.db.GetProductById(ctx, id)
	if err != nil {
		return domain.Product{}, err
	}
	return domain.Product{
		ID:        results.ID,
		Name:      results.Name.String,
		Price:     results.Price.Float64,
		Image:     results.Image.String,
		CreatedAt: results.CreatedAt.Time,
		UpdatedAt: results.UpdatedAt.Time,
	}, nil
}
func (app *Application) ListProducts(ctx context.Context, limit int32, offset int32) ([]domain.Product, error) {

	results, err := app.db.ListProducts(ctx, generated.ListProductsParams{
		Limit:  limit,
		Offset: offset,
	})
	if err != nil {
		return nil, err
	}
	var products []domain.Product
	for _, result := range results {
		products = append(products, domain.Product{
			ID:        result.ID,
			Name:      result.Name.String,
			Price:     result.Price.Float64,
			Image:     result.Image.String,
			CreatedAt: result.CreatedAt.Time,
			UpdatedAt: result.UpdatedAt.Time,
		})
	}
	return products, nil
}
