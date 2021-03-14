package mapper

import (
	"github.com/Jeeo/golang-ddd-boilerplate/internal/domain/entity"
	"github.com/Jeeo/golang-ddd-boilerplate/internal/entrypoint/dto"
	"github.com/google/wire"
)

type PersonMapper interface {
	ToDTO(entity.Person) dto.PersonDTO
	FromDTO(dto.PersonDTO) entity.Person
	ToManyDTO([]entity.Person) []dto.PersonDTO
}

var PersonMapperSet = wire.NewSet(
	ProvidePersonMapper,
	wire.Bind(
		new(PersonMapper),
		new(*PersonMapperImpl),
	),
)
