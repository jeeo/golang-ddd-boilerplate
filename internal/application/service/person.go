package service

import (
	"github.com/Jeeo/golang-ddd-boilerplate/internal/domain/repository"
	"github.com/Jeeo/golang-ddd-boilerplate/internal/entrypoint/dto"
	"github.com/Jeeo/golang-ddd-boilerplate/internal/entrypoint/mapper"
)

type PersonService struct {
	repository repository.PersonRepository
	mapper     mapper.PersonMapper
}

func (ps *PersonService) Create(payload dto.PersonDTO) dto.PersonDTO {
	person := ps.repository.Create(ps.mapper.FromDTO(payload))

	return ps.mapper.ToDTO(person)
}

func (ps *PersonService) GetById(id int32) dto.PersonDTO {
	person := ps.repository.GetOne(id)

	return ps.mapper.ToDTO(person)
}

func (ps *PersonService) GetAll() []dto.PersonDTO {
	people := ps.repository.GetAll()

	return ps.mapper.ToManyDTO(people)
}

func (ps *PersonService) Update(payload dto.PersonDTO) dto.PersonDTO {
	person := ps.repository.Update(ps.mapper.FromDTO(payload))

	return ps.mapper.ToDTO(person)
}

func (ps *PersonService) Delete(id int32) bool {
	success := ps.repository.Delete(id)

	return success
}

func ProvidePersonService(r repository.PersonRepository, m mapper.PersonMapper) *PersonService {
	return &PersonService{
		repository: r,
		mapper:     m,
	}
}
