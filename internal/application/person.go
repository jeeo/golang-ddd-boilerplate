package application

import (
	"github.com/Jeeo/golang-ddd-boilerplate/internal/domain/service"
	"github.com/Jeeo/golang-ddd-boilerplate/internal/entrypoint/dto"
	"github.com/Jeeo/golang-ddd-boilerplate/internal/entrypoint/mapper"
)

type PersonApplicationImpl struct {
	service service.PersonService
	mapper  mapper.PersonMapper
}

func (pa *PersonApplicationImpl) CreatePerson(payload dto.PersonDTO) dto.PersonDTO {
	person := pa.service.Create(pa.mapper.FromDTO(payload))

	return pa.mapper.ToDTO(person)
}

func (pa *PersonApplicationImpl) FindOne(id int32) dto.PersonDTO {
	person := pa.service.GetById(id)

	return pa.mapper.ToDTO(person)
}

func (pa *PersonApplicationImpl) FindAll() []dto.PersonDTO {
	people := pa.service.GetAll()

	return pa.mapper.ToManyDTO(people)
}

func (pa *PersonApplicationImpl) UpdatePerson(id int32, payload dto.PersonDTO) dto.PersonDTO {
	person := pa.service.Update(id, pa.mapper.FromDTO(payload))

	return pa.mapper.ToDTO(person)
}

func (pa *PersonApplicationImpl) DeletePerson(id int32) bool {
	success := pa.service.Delete(id)

	return success
}

func ProvidePersonApplication(s service.PersonService, m mapper.PersonMapper) *PersonApplicationImpl {
	return &PersonApplicationImpl{
		service: s,
		mapper:  m,
	}
}
