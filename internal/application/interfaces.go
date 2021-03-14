package application

import (
	"github.com/Jeeo/golang-ddd-boilerplate/internal/entrypoint/dto"
	"github.com/google/wire"
)

type PersonApplication interface {
	CreatePerson(dto.PersonDTO) dto.PersonDTO
	FindOne(int32) dto.PersonDTO
	FindAll() []dto.PersonDTO
	UpdatePerson(int32, dto.PersonDTO) dto.PersonDTO
	DeletePerson(int32) bool
}

var PersonApplicationSet = wire.NewSet(
	ProvidePersonApplication,
	wire.Bind(
		new(PersonApplication),
		new(*PersonApplicationImpl),
	),
)
