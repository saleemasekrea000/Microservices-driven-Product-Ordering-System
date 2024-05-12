package test

import (
	"context"
	"os"
	"product/internal/adapters/db/generated"
	"testing"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/suite"
)

type DatabaseSuite struct {
	suite.Suite
	db *generated.Queries
}

func TestDatabaseSuite(t *testing.T) {
	suite.Run(t, new(DatabaseSuite))
}

func (suite *DatabaseSuite) SetupSuite() {
	godotenv.Load("../.env")
	dns := os.Getenv("test_dns")
	pg, err := pgxpool.New(context.Background(), dns)
	if err != nil {
		suite.Fail(err.Error())
	}
	err = pg.Ping(context.Background())
	if err != nil {
		suite.Fail(err.Error())
	}
	pg.Exec(context.Background(), "C")

	suite.db = generated.New(pg)

}

func (suite *DatabaseSuite) TestCreateProduct() {
	product, err := suite.db.CreateProduct(context.Background(), generated.CreateProductParams{
		Name:  pgtype.Text{String: "test", Valid: true},
		Price: pgtype.Float8{Float64: 100, Valid: true},
	})
	suite.NoError(err)
	suite.NotNil(product)
	suite.Equal("test", product.Name.String)
	suite.Equal(float64(100), product.Price.Float64)
}
func (suite *DatabaseSuite) TestGetProduct() {
	suite.T().Log("TestGetProduct")
	product, err := suite.db.CreateProduct(context.Background(), generated.CreateProductParams{
		Name:  pgtype.Text{String: "test", Valid: true},
		Price: pgtype.Float8{Float64: 100, Valid: true},
	})
	suite.NoError(err)
	suite.NotNil(product)
	suite.Equal("test", product.Name.String)
	suite.Equal(float64(100), product.Price.Float64)
	product, err = suite.db.GetProductById(context.Background(), 1)
	suite.NoError(err)
	suite.NotNil(product)
	suite.Equal("test", product.Name.String)
	suite.Equal(float64(100), product.Price.Float64)
}

func (suite *DatabaseSuite) TestListProducts() {
	suite.db.CreateProduct(context.Background(), generated.CreateProductParams{
		Name:  pgtype.Text{String: "test", Valid: true},
		Price: pgtype.Float8{Float64: 100, Valid: true},
	})
	suite.db.CreateProduct(context.Background(), generated.CreateProductParams{
		Name:  pgtype.Text{String: "test", Valid: true},
		Price: pgtype.Float8{Float64: 100, Valid: true},
	})
	products, err := suite.db.ListProducts(context.Background(), generated.ListProductsParams{Limit: 2, Offset: 0})
	suite.NoError(err)
	suite.True(len(products) > 1)
	suite.NotZero(products[0].ID)
}
