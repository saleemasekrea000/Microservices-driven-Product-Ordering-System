package http

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func (a *httpAdapter) GetOrder(c *fiber.Ctx) error {
	id := c.Params("id")
	intId, err := strconv.ParseInt(id, 10, 32)
	if id == "" || err != nil || intId < 1 {
		return c.Status(fiber.ErrBadRequest.Code).SendString("invalid id parameter")
	}

	order, err := a.api.GetFullOrder(c.Context(), int32(intId))
	if err != nil {
		return c.Status(fiber.ErrInternalServerError.Code).SendString(err.Error())
	}
	return c.JSON(order)
}
