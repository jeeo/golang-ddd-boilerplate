package http

import (
	"github.com/Jeeo/golang-ddd-boilerplate/internal/entrypoint/handler"

	"github.com/google/wire"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type EchoServer struct {
	handler *handler.PersonHandler
}

func (s *EchoServer) Init() {
	e := echo.New()

	// it would be really good if was injected =) but i'm lazy
	e.Use(middleware.Logger())
	e.Use(middleware.CORS())

	e.GET("/person/:id", s.handler.GetById)
	e.Start(":3000")
}

func ProvideEchoServer(h *handler.PersonHandler) *EchoServer {
	return &EchoServer{
		handler: h,
	}
}

var EchoSet = wire.NewSet(
	ProvideEchoServer,
	wire.Bind(
		new(HttpService),
		new(*EchoServer),
	),
)
