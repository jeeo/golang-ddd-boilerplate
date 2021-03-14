package mapper

import (
	"github.com/Jeeo/golang-ddd-boilerplate/internal/domain/entity"
	"github.com/Jeeo/golang-ddd-boilerplate/internal/entrypoint/dto"
)

type PersonMapperImpl struct{}

func (pm *PersonMapperImpl) ToDTO(p entity.Person) dto.PersonDTO {
	return dto.PersonDTO{ID: p.ID, Name: p.Name, Age: p.Age}
}

func (pm *PersonMapperImpl) ToManyDTO(people []entity.Person) []dto.PersonDTO {
	var peopleDTO []dto.PersonDTO
	for _, p := range people {
		peopleDTO = append(peopleDTO, dto.PersonDTO{ID: p.ID, Name: p.Name, Age: p.Age})
	}

	return peopleDTO
}

func (pm *PersonMapperImpl) FromDTO(p dto.PersonDTO) entity.Person {
	return entity.Person{ID: p.ID, Name: p.Name, Age: p.Age}
}

func ProvidePersonMapper() *PersonMapperImpl {
	return &PersonMapperImpl{}
}
