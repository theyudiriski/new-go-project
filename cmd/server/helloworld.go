package server

import "github.com/gofiber/fiber/v3"

func (s *server) handleHelloWorld(c fiber.Ctx) error {
	return c.SendString("Hello, World 👋!")
}
