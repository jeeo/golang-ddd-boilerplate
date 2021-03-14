package http

import (
	"fmt"

	"github.com/Jeeo/golang-ddd-boilerplate/configs"
	"github.com/Jeeo/golang-ddd-boilerplate/internal/entrypoint/handler"

	"github.com/google/wire"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type EchoServer struct {
	config  configs.Config
	handler handler.PersonHandler
}

func (s *EchoServer) Init() {
	port := s.config.Server.Port
	if port == "" {
		port = "3000"
	}
	e := echo.New()
	// it would be really good if was injected =) but i'm lazy
	e.Use(middleware.Logger())
	e.Use(middleware.CORS())

	e.GET("/person/:id", s.handler.GetById)
	e.GET("/person", s.handler.GetAll)
	e.POST("/person", s.handler.Create)
	e.PUT("/person/:id", s.handler.Update)
	e.DELETE("/person/:id", s.handler.Delete)

	e.Start(fmt.Sprintf(":%s", port))
}

func ProvideEchoServer(h *handler.PersonHandlerImpl, c configs.Config) *EchoServer {
	return &EchoServer{
		handler: h,
		config:  c,
	}
}

var EchoSet = wire.NewSet(
	ProvideEchoServer,
	wire.Bind(
		new(HttpService),
		new(*EchoServer),
	),
)
