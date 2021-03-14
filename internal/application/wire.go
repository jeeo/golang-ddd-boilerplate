package application

import (
	"github.com/google/wire"
)

var ApplicationSet = wire.NewSet(
	PersonApplicationSet,
)
