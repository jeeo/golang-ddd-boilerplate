package service

import (
	"github.com/Jeeo/golang-ddd-boilerplate/internal/domain/entity"
	"github.com/Jeeo/golang-ddd-boilerplate/internal/domain/repository"
)

type PersonService struct {
	repository repository.PersonRepository
}

func (ps *PersonService) Create(payload entity.Person) entity.Person {
	person := ps.repository.Create(payload)

	return person
}

func (ps *PersonService) GetById(id int32) entity.Person {
	person := ps.repository.GetOne(id)

	return person
}

func (ps *PersonService) GetAll() []entity.Person {
	people := ps.repository.GetAll()

	return people
}

func (ps *PersonService) Update(payload entity.Person) entity.Person {
	person := ps.repository.Update(payload)

	return person
}

func (ps *PersonService) Delete(id int32) bool {
	success := ps.repository.Delete(id)

	return success
}

func ProvidePersonService(r repository.PersonRepository) *PersonService {
	return &PersonService{
		repository: r,
	}
}
