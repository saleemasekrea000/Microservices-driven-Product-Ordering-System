package domain

import "time"

type Product struct {
	ID        int32     `json:"id"`
	Name      string    `json:"name"`
	Price     float64   `json:"price"`
	Image     string    `json:"image"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewProduct(name string, price float64, image string) Product {
	return Product{
		Name:      name,
		Price:     price,
		Image:     image,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
