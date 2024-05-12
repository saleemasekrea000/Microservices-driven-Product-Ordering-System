package test

import (
	"context"
	"order/internal/adapters/db/generated"
	"os"
	"testing"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/suite"
)

type DataBaseSuite struct {
	suite.Suite
	db *generated.Queries
}

func TestDatabaseSuite(t *testing.T) {
	suite.Run(t, new(DataBaseSuite))
}

func (suite *DataBaseSuite) SetupSuite() {
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
	suite.db = generated.New(pg)
}

func (suite *DataBaseSuite) TestCreateOrder() {

	err := suite.db.CreateOrder(context.Background(), generated.CreateOrderParams{
		TotalPrice: 100,
		UserID:     1,
		UpdatedAt: pgtype.Timestamp{
			Time:  time.Now(),
			Valid: true,
		},
	})
	suite.NoError(err)
}
