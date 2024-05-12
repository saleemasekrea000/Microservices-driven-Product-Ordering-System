package api

import (
	"context"
	"product/internal/adapters/db/generated"
	"product/internal/application/core/domain"
	"product/mocks"
	"testing"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/assert"
)

func TestGetProducts(t *testing.T) {
	mockesDb := mocks.NewDBPort(t)
	mockesDb.On("ListProducts", context.Background(), generated.ListProductsParams{
		Limit:  10,
		Offset: 0,
	}).Return([]generated.Product{}, nil)
	apiAdapter := New(mockesDb)
	products, err := apiAdapter.ListProducts(context.Background(), 10, 0)
	assert.Nil(t, err)
	assert.Equal(t, 0, len(products))

}

func TestCreateProduct(t *testing.T) {

	mockesDb := mocks.NewDBPort(t)
	mockesDb.On("CreateProduct", context.Background(), generated.CreateProductParams{
		Name:  pgtype.Text{String: "test", Valid: true},
		Price: pgtype.Float8{Float64: 10, Valid: true},
		Image: pgtype.Text{String: "test", Valid: true},
	}).Return(generated.Product{
		ID:        1,
		Name:      pgtype.Text{String: "test", Valid: true},
		Price:     pgtype.Float8{Float64: 10, Valid: true},
		Image:     pgtype.Text{String: "test", Valid: true},
		CreatedAt: pgtype.Timestamp{Time: time.Now(), Valid: true},
		UpdatedAt: pgtype.Timestamp{Time: time.Now(), Valid: true},
	}, nil)
	apiAdapter := New(mockesDb)
	product, err := apiAdapter.CreateProduct(context.Background(), domain.Product{
		Name:  "test",
		Price: 10,
		Image: "test",
	})
	assert.Nil(t, err)
	assert.Equal(t, int32(1), product.ID)
	assert.Equal(t, "test", product.Name)
	assert.Equal(t, float64(10), product.Price)
	assert.Equal(t, "test", product.Image)

}
func TestGetProduct(t *testing.T) {

	mockesDb := mocks.NewDBPort(t)
	mockesDb.On("GetProductById", context.Background(), int32(1)).Return(generated.Product{
		ID:        1,
		Name:      pgtype.Text{String: "test", Valid: true},
		Price:     pgtype.Float8{Float64: 10, Valid: true},
		Image:     pgtype.Text{String: "test", Valid: true},
		CreatedAt: pgtype.Timestamp{Time: time.Now(), Valid: true},
		UpdatedAt: pgtype.Timestamp{Time: time.Now(), Valid: true},
	}, nil)
	apiAdapter := New(mockesDb)
	product, err := apiAdapter.GetProductById(context.Background(), 1)
	assert.Nil(t, err)
	assert.Equal(t, int32(1), product.ID)
	assert.Equal(t, "test", product.Name)
	assert.Equal(t, float64(10), product.Price)
	assert.Equal(t, "test", product.Image)
}
