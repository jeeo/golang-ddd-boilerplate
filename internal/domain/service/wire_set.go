package service

import (
	"github.com/google/wire"
)

var DomainServiceSet = wire.NewSet(
	ProvidePersonService,
)
