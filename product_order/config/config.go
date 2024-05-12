package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Configration struct {
	Port       int
	OrderUrl   string
	ProductUrl string
}

func New() *Configration {
	godotenv.Load("../")
	port := os.Getenv("PORT")
	portInt, err := strconv.Atoi(port)
	if err != nil {
		panic(err)
	}
	OrderUrl := os.Getenv("ORDER_URL")
	ProductUrl := os.Getenv("PRODUCT_URL")
	return &Configration{
		Port:       portInt,
		OrderUrl:   OrderUrl,
		ProductUrl: ProductUrl,
	}
}
