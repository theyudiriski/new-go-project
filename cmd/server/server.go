package server

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Server interface {
	Start()
}

type server struct {
	app *echo.Echo
}

func NewServer() Server {
	echoApp := echo.New()

	return &server{
		app: echoApp,
	}
}

func (s *server) Start() {
	s.app.Use(middleware.Logger())
	s.app.Use(middleware.Recover())

	// register API routes
	s.app.GET("/hello-world", s.handleHelloWorld)
	s.app.GET("/hello/:name", s.handleHello)

	s.app.Start(":8080")
}
