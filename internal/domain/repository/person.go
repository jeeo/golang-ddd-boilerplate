package repository

import (
	"github.com/Jeeo/golang-ddd-boilerplate/internal/domain/entity"
	"github.com/Jeeo/golang-ddd-boilerplate/internal/infrastructure/repository"

	"github.com/google/wire"
)

// PersonRepository is a interface for person repository
type PersonRepository interface {
	Create(entity.Person) entity.Person
	GetOne(int32) entity.Person
	GetAll() []entity.Person
	Update(entity.Person) entity.Person
	Delete(int32) bool
}

var PersonRepositorySet = wire.NewSet(
	repository.ProvidePersonRepository,
	wire.Bind(
		new(PersonRepository),
		new(*repository.PersonRepository),
	),
)
