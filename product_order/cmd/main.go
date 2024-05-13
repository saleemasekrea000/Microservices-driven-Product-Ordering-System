package main

import (
	"product_order/config"
	"product_order/internal/adapters/http"
	"product_order/internal/adapters/order"
	"product_order/internal/adapters/product"
	"product_order/internal/application/core/api"
)

func main() {
	config := config.New()
	orderProt, err := order.New(config.OrderUrl)
	if err != nil {
		panic(err)
	}
	productPort, err := product.New(config.ProductUrl)
	if err != nil {
		panic(err)
	}
	apiAdapter := api.New(orderProt, productPort)
	httpServer := http.New(apiAdapter, config.Port)
	httpServer.Run()

}
