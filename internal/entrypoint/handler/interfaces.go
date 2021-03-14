package handler

import (
	"github.com/google/wire"
	"github.com/labstack/echo/v4"
)

type PersonHandler interface {
	GetById(echo.Context) error
}

var PersonHandlerSet = wire.NewSet(
	ProvidePersonHandler,
	wire.Bind(
		new(PersonHandler),
		new(*PersonHandlerImpl),
	),
)
