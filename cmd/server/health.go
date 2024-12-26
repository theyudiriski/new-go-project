package server

import (
	"log"

	"github.com/gofiber/fiber/v3"
)

func (s *server) handleHealth(c fiber.Ctx) error {
	err := s.db.Leader.Ping()
	if err != nil {
		log.Println("failed to ping db", "error", err)
		return c.Status(fiber.StatusInternalServerError).SendString("DB is not healthy")
	}

	return c.SendString("OK")
}
