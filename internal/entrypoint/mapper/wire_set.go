package mapper

import (
	"github.com/google/wire"
)

var MapperSet = wire.NewSet(
	ProvidePersonMapper,
)
