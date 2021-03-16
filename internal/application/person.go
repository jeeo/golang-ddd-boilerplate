package application

import (
	"github.com/Jeeo/golang-ddd-boilerplate/internal/domain/repository"
	"github.com/Jeeo/golang-ddd-boilerplate/internal/entrypoint/dto"
	"github.com/Jeeo/golang-ddd-boilerplate/internal/entrypoint/mapper"
)

type PersonApplicationImpl struct {
	repository repository.PersonRepository
	mapper     mapper.PersonMapper
}

func (pa *PersonApplicationImpl) CreatePerson(payload dto.PersonDTO) dto.PersonDTO {
	person := pa.repository.CreatePerson(pa.mapper.FromDTO(payload))

	return pa.mapper.ToDTO(person)
}

func (pa *PersonApplicationImpl) FindOne(id int32) dto.PersonDTO {
	person := pa.repository.GetPerson(id)

	return pa.mapper.ToDTO(person)
}

func (pa *PersonApplicationImpl) FindAll() []dto.PersonDTO {
	people := pa.repository.GetPeople()

	return pa.mapper.ToManyDTO(people)
}

func (pa *PersonApplicationImpl) UpdatePerson(id int32, payload dto.PersonDTO) dto.PersonDTO {
	person := pa.repository.UpdatePerson(id, pa.mapper.FromDTO(payload))

	return pa.mapper.ToDTO(person)
}

func (pa *PersonApplicationImpl) DeletePerson(id int32) bool {
	success := pa.repository.DeletePerson(id)

	return success
}

func ProvidePersonApplication(r repository.PersonRepository, m mapper.PersonMapper) *PersonApplicationImpl {
	return &PersonApplicationImpl{
		repository: r,
		mapper:     m,
	}
}
