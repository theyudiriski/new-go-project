package server

import (
	"github.com/gofiber/fiber/v3"
)

type Server interface {
	Start()
}

type server struct {
	app *fiber.App
}

func NewServer() Server {
	app := fiber.New()

	return &server{
		app: app,
	}
}

func (s *server) Start() {
	s.app.Get("/hello-world", s.handleHelloWorld)
	s.app.Get("/hello/:name", s.handleHello)

	s.app.Listen(":8080")
}
