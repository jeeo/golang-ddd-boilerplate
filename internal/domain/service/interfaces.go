package service

import (
	"github.com/Jeeo/golang-ddd-boilerplate/internal/domain/entity"
	"github.com/google/wire"
)

type PersonService interface {
	Create(entity.Person) entity.Person
	GetById(int32) entity.Person
	GetAll() []entity.Person
	Update(int32, entity.Person) entity.Person
	Delete(int32) bool
}

var PersonServiceSet = wire.NewSet(
	ProvidePersonService,
	wire.Bind(
		new(PersonService),
		new(*PersonServiceImpl),
	),
)
