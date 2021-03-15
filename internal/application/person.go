package application

import (
	"github.com/Jeeo/golang-ddd-boilerplate/internal/application/repository"
	"github.com/Jeeo/golang-ddd-boilerplate/internal/entrypoint/dto"
	"github.com/Jeeo/golang-ddd-boilerplate/internal/entrypoint/mapper"
)

type PersonApplicationImpl struct {
	repository repository.PersonRepository
	mapper     mapper.PersonMapper
}

func (pa *PersonApplicationImpl) CreatePerson(payload dto.PersonDTO) dto.PersonDTO {
	person := pa.repository.Create(pa.mapper.FromDTO(payload))

	return pa.mapper.ToDTO(person)
}

func (pa *PersonApplicationImpl) FindOne(id int32) dto.PersonDTO {
	person := pa.repository.GetOne(id)

	return pa.mapper.ToDTO(person)
}

func (pa *PersonApplicationImpl) FindAll() []dto.PersonDTO {
	people := pa.repository.GetAll()

	return pa.mapper.ToManyDTO(people)
}

func (pa *PersonApplicationImpl) UpdatePerson(id int32, payload dto.PersonDTO) dto.PersonDTO {
	person := pa.repository.Update(id, pa.mapper.FromDTO(payload))

	return pa.mapper.ToDTO(person)
}

func (pa *PersonApplicationImpl) DeletePerson(id int32) bool {
	success := pa.repository.Delete(id)

	return success
}

func ProvidePersonApplication(r repository.PersonRepository, m mapper.PersonMapper) *PersonApplicationImpl {
	return &PersonApplicationImpl{
		repository: r,
		mapper:     m,
	}
}
