package domains

import "time"

type Product struct {
	ID        int32     `json:"id"`
	Name      string    `json:"name"`
	Price     float64   `json:"price"`
	Image     string    `json:"image"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Order struct {
	ID         int32        `json:"id"`
	TotalPrice float64      `json:"total_price"`
	UserID     int32        `json:"user_id"`
	CreatedAt  time.Time    `json:"created_at"`
	UpdatedAt  time.Time    `json:"updated_at"`
	OrderItems []*OrderItem `json:"order_items"`
}

type OrderItem struct {
	ID      int32    `json:"id"`
	OrderID int32    `json:"order_id"`
	Product *Product `json:"product"`
	Amount  int32    `json:"amount"`
	Price   float64  `json:"price"`
}
