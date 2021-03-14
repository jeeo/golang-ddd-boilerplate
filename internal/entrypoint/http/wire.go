package http

import (
	"github.com/google/wire"
)

var HttpSet = wire.NewSet(
	EchoSet,
)
