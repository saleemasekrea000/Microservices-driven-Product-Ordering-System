package domain

import (
	"order/internal/adapters/db/generated"

	"github.com/jackc/pgx/v5/pgtype"
)

func FromGeneratedOrderToDomainOrder(generatedOrder *generated.Order) *Order {

	return &Order{
		ID:         generatedOrder.ID,
		TotalPrice: float64(generatedOrder.TotalPrice),
		UserID:     generatedOrder.UserID,
		CreatedAt:  generatedOrder.CreatedAt.Time,
		UpdatedAt:  generatedOrder.UpdatedAt.Time,
	}
}
func FromDomainOrderToGenetatedOrder(domainOrder *Order) *generated.Order {
	return &generated.Order{
		ID:         domainOrder.ID,
		TotalPrice: domainOrder.TotalPrice,
		UserID:     domainOrder.UserID,
		CreatedAt: pgtype.Timestamp{
			Time:  domainOrder.CreatedAt,
			Valid: true,
		},
		UpdatedAt: pgtype.Timestamp{
			Time:  domainOrder.UpdatedAt,
			Valid: true,
		},
	}
}

func FromDomainOrderItemToGenetatedOrderItem(domainOrderItem *OrderItem) *generated.OrderItem {
	return &generated.OrderItem{
		ID: domainOrderItem.ID,
		OrderID: pgtype.Int4{
			Int32: domainOrderItem.OrderID,
			Valid: true,
		},
		ProductID: pgtype.Int4{
			Int32: domainOrderItem.ProductID,
			Valid: true,
		},
		Amount: pgtype.Int4{Int32: domainOrderItem.Amount, Valid: true},
		Price:  pgtype.Float8{Float64: domainOrderItem.Price, Valid: true},
	}
}

func FromGenetatedOrderItemToDomainOrderItem(generatedOrderItem *generated.OrderItem) *OrderItem {
	return &OrderItem{
		ID:        generatedOrderItem.ID,
		OrderID:   generatedOrderItem.OrderID.Int32,
		ProductID: generatedOrderItem.ProductID.Int32,
		Amount:    generatedOrderItem.Amount.Int32,
		Price:     generatedOrderItem.Price.Float64,
	}
}

func TransformGetOrdersToOrder(getOrders *[]generated.GetOrdersRow) []Order {
	orderMap := make(map[int32]*Order)

	for _, getOrder := range *getOrders {
		orderID := getOrder.ID
		if order, ok := orderMap[orderID]; ok {
			orderItem := &OrderItem{
				ID:        getOrder.ID_2,
				OrderID:   getOrder.OrderID.Int32,
				ProductID: getOrder.ProductID.Int32,
				Amount:    getOrder.Amount.Int32,
				Price:     getOrder.Price.Float64,
			}
			order.OrderItems = append(order.OrderItems, orderItem)
		} else {
			order := &Order{
				ID:         getOrder.ID,
				TotalPrice: getOrder.TotalPrice,
				UserID:     getOrder.UserID,
				CreatedAt:  getOrder.CreatedAt.Time,
				UpdatedAt:  getOrder.UpdatedAt.Time,
				OrderItems: []*OrderItem{
					{
						ID:        getOrder.ID_2,
						OrderID:   getOrder.OrderID.Int32,
						ProductID: getOrder.ProductID.Int32,
						Amount:    getOrder.Amount.Int32,
						Price:     getOrder.Price.Float64,
					},
				},
			}
			orderMap[orderID] = order
		}
	}

	orders := make([]Order, 0, len(orderMap))
	for _, order := range orderMap {
		orders = append(orders, *order)
	}

	return orders
}

func TransformGetOrderToOrder(getOrders *[]generated.GetOrderRow) []Order {
	orderMap := make(map[int32]*Order)

	for _, getOrder := range *getOrders {
		orderID := getOrder.ID
		if order, ok := orderMap[orderID]; ok {
			orderItem := &OrderItem{
				ID:        getOrder.ID_2.Int32,
				OrderID:   getOrder.OrderID.Int32,
				ProductID: getOrder.ProductID.Int32,
				Amount:    getOrder.Amount.Int32,
				Price:     getOrder.Price.Float64,
			}
			order.OrderItems = append(order.OrderItems, orderItem)
		} else {
			order := &Order{
				ID:         getOrder.ID,
				TotalPrice: getOrder.TotalPrice,
				UserID:     getOrder.UserID,
				CreatedAt:  getOrder.CreatedAt.Time,
				UpdatedAt:  getOrder.UpdatedAt.Time,
				OrderItems: []*OrderItem{
					{
						ID:        getOrder.ID_2.Int32,
						OrderID:   getOrder.OrderID.Int32,
						ProductID: getOrder.ProductID.Int32,
						Amount:    getOrder.Amount.Int32,
						Price:     getOrder.Price.Float64,
					},
				},
			}
			orderMap[orderID] = order
		}
	}

	orders := make([]Order, 0, len(orderMap))
	for _, order := range orderMap {
		orders = append(orders, *order)
	}

	return orders
}
