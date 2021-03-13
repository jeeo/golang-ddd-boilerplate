package application

import (
	"github.com/Jeeo/golang-ddd-boilerplate/internal/application/service"
	"github.com/Jeeo/golang-ddd-boilerplate/internal/entrypoint/dto"
	"github.com/Jeeo/golang-ddd-boilerplate/internal/entrypoint/mapper"
)

type PersonApplication struct {
	service service.PersonService
	mapper  mapper.PersonMapper
}

func (pa *PersonApplication) CreatePerson(payload dto.PersonDTO) dto.PersonDTO {
	person := pa.service.Create(pa.mapper.FromDTO(payload))

	return pa.mapper.ToDTO(person)
}

func (pa *PersonApplication) FindOne(id int32) dto.PersonDTO {
	person := pa.service.GetById(id)

	return pa.mapper.ToDTO(person)
}

func (pa *PersonApplication) FindAll() []dto.PersonDTO {
	people := pa.service.GetAll()

	return pa.mapper.ToManyDTO(people)
}

func (pa *PersonApplication) UpdatePerson(payload dto.PersonDTO) dto.PersonDTO {
	person := pa.service.Update(pa.mapper.FromDTO(payload))

	return pa.mapper.ToDTO(person)
}

func (pa *PersonApplication) DeletePerson(id int32) bool {
	success := pa.service.Delete(id)

	return success
}

func ProvidePersonApplication(s service.PersonService, m mapper.PersonMapper) *PersonApplication {
	return &PersonApplication{
		service: s,
		mapper:  m,
	}
}
