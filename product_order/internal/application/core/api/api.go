package api

import (
	"context"
	"log"
	"product_order/internal/application/core/domains"
	"product_order/internal/ports"
)

type ApiAdapter struct {
	order   ports.OrderPort
	product ports.ProductPort
}

func New(order ports.OrderPort, product ports.ProductPort) *ApiAdapter {
	return &ApiAdapter{
		order:   order,
		product: product,
	}
}

func (api *ApiAdapter) GetFullOrder(ctx context.Context, orderID int32) (domains.Order, error) {

	order, err := api.order.GetOrder(ctx, orderID)
	if err != nil {
		log.Println("err", err)
		return domains.Order{}, err
	}
	log.Println("pass order")
	log.Print("order id", orderID)

	for _, orderItem := range order.OrderItems {
		log.Println("orderItem", orderItem.Product.ID)
		if orderItem.Product.ID == 0 {
			continue
		}
		product, err := api.product.GetProduct(ctx, orderItem.Product.ID)
		if err != nil {
			log.Println("err", err)
			return domains.Order{}, err
		}
		orderItem.Product = &product
	}
	return order, nil
}
