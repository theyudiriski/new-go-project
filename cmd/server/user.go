package server

import (
	"encoding/json"
	"log"
	"new-go-project/cmd/service"

	"github.com/gofiber/fiber/v3"
)

func (s *server) handleCreateUser(c fiber.Ctx) error {
	ctx := c.Context()

	// parse request
	var user service.User
	err := json.Unmarshal(c.Body(), &user)
	if err != nil {
		log.Println("failed to unmarshal request", "error", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid request",
		})
	}

	// create user
	err = s.userService.CreateUser(ctx, &user)
	if err != nil {
		log.Println("failed to create user", "error", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to create user",
		})
	}

	// return response
	return c.Status(fiber.StatusCreated).JSON(user)
}

func (s *server) handleGetUsers(c fiber.Ctx) error {
	ctx := c.Context()

	users, err := s.userService.GetUsers(ctx)
	if err != nil {
		log.Println("failed to get users", "error", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to get users",
		})
	}

	return c.JSON(users)
}
