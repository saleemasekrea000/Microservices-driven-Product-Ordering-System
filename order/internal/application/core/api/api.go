package api

import (
	"context"
	"log"
	"order/internal/adapters/db/generated"
	"order/internal/application/core/domain"
	"order/internal/ports"
	"time"

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

func (app *Application) CreateOrder(ctx context.Context, order domain.Order) error {
	arg := generated.CreateOrderParams{
		TotalPrice: order.TotalPrice,
		UserID:     order.UserID,
		UpdatedAt:  pgtype.Timestamp{Time: time.Now(), Valid: true},
	}
	err := app.db.CreateOrder(ctx, arg)
	if err != nil {
		return err
	}
	for _, orderItem := range order.OrderItems {
		arg := generated.CreateOrderItemsParams{
			OrderID:   pgtype.Int4{Int32: orderItem.OrderID, Valid: true},
			ProductID: pgtype.Int4{Int32: orderItem.ProductID, Valid: true},
			Amount:    pgtype.Int4{Int32: orderItem.Amount, Valid: true},
			Price:     pgtype.Float8{Float64: orderItem.Price, Valid: true},
		}
		log.Println("arg", arg)
		err := app.db.CreateOrderItems(ctx, arg)
		if err != nil {
			return err
		}
	}
	return nil
}
func (app *Application) GetOrders(ctx context.Context, limit int32, offset int32) ([]domain.Order, error) {
	arg := generated.GetOrdersParams{
		Limit:  limit,
		Offset: offset,
	}
	orders, err := app.db.GetOrders(ctx, arg)
	log.Println("orders", orders)
	if err != nil {
		return nil, err
	}

	return domain.TransformGetOrdersToOrder(&orders), nil
}
func (app *Application) GetOrder(ctx context.Context, orderID int32) (*domain.Order, error) {
	order, err := app.db.GetOrder(ctx, orderID)
	if err != nil {
		return nil, err
	}
	if len(order) == 0 {
		return nil, nil
	}
	return &domain.TransformGetOrderToOrder(&order)[0], nil
}
