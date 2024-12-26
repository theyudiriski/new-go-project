package server

import (
	"fmt"

	"github.com/gofiber/fiber/v3"
)

func (s *server) handleHello(c fiber.Ctx) error {
	name := c.Params("name")

	return c.JSON(map[string]string{
		"message": fmt.Sprintf("Hello, %s 👋!", name),
	})
}
