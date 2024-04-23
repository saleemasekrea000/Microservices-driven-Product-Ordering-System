package domain

import "time"

type Order struct {
	ID         int32        `json:"id"`
	TotalPrice float64      `json:"total_price"`
	UserID     int32        `json:"user_id"`
	CreatedAt  time.Time    `json:"created_at"`
	UpdatedAt  time.Time    `json:"updated_at"`
	OrderItems []*OrderItem `json:"order_items"`
}

type OrderItem struct {
	ID        int32   `json:"id"`
	OrderID   int32   `json:"order_id"`
	ProductID int32   `json:"product_id"`
	Amount    int32   `json:"amount"`
	Price     float64 `json:"price"`
}

func NewOrder(totalPrice float64, userID int32, orderItems []*OrderItem) *Order {
	return &Order{
		TotalPrice: totalPrice,
		UserID:     userID,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
		OrderItems: orderItems,
	}
}

func NewOrderItem(orderID int32, productID int32, amount int32, price float64) *OrderItem {
	return &OrderItem{
		OrderID:   orderID,
		ProductID: productID,
		Amount:    amount,
		Price:     price,
	}
}
