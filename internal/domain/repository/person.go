package repository

import (
	"github.com/Jeeo/golang-ddd-boilerplate/internal/infrastructure/repository"
	"github.com/google/wire"
)

// PersonRepository is a interface for person repository
var PersonRepositorySet = wire.NewSet(
	repository.ProvidePersonRepository,
	wire.Bind(
		new(PersonRepository),
		new(*repository.PersonRepositoryImpl),
	),
)
