package http

import (
	"fmt"
	"log"
	"product_order/internal/ports"

	"github.com/gofiber/fiber/v2"
)

type httpAdapter struct {
	api  ports.ApiPort
	port int
}

func New(api ports.ApiPort, port int) *httpAdapter {
	return &httpAdapter{
		api:  api,
		port: port,
	}
}

func (h *httpAdapter) Run() {

	app := fiber.New()

	log.Println("Starting server on port", h.port)
	app.Get("/order/:id", h.GetOrder)

	log.Fatalf("fiber failed to start: %v\n", app.Listen(fmt.Sprintf(":%d", h.port)))
}
