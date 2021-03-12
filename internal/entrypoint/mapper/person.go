package mapper

import (
	"github.com/Jeeo/golang-ddd-boilerplate/internal/domain/entity"
	"github.com/Jeeo/golang-ddd-boilerplate/internal/entrypoint/dto"
)

type PersonMapper struct{}

func (pm *PersonMapper) ToDTO(p entity.Person) dto.PersonDTO {
	return dto.PersonDTO{ID: p.ID, Name: p.Name, Age: p.Age}
}

func (pm *PersonMapper) ToManyDTO(people []entity.Person) []dto.PersonDTO {
	var peopleDTO []dto.PersonDTO
	for _, p := range people {
		peopleDTO = append(peopleDTO, dto.PersonDTO{ID: p.ID, Name: p.Name, Age: p.Age})
	}

	return peopleDTO
}

func (pm *PersonMapper) FromDTO(p dto.PersonDTO) entity.Person {
	return entity.Person{ID: p.ID, Name: p.Name, Age: p.Age}
}

func ProvidePersonMapper() PersonMapper {
	return PersonMapper{}
}
