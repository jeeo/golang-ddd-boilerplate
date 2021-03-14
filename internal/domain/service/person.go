package service

import (
	"github.com/Jeeo/golang-ddd-boilerplate/internal/domain/entity"
	"github.com/Jeeo/golang-ddd-boilerplate/internal/domain/repository"
)

type PersonServiceImpl struct {
	repository repository.PersonRepository
}

func (ps *PersonServiceImpl) Create(payload entity.Person) entity.Person {
	person := ps.repository.Create(payload)

	return person
}

func (ps *PersonServiceImpl) GetById(id int32) entity.Person {
	person := ps.repository.GetOne(id)

	return person
}

func (ps *PersonServiceImpl) GetAll() []entity.Person {
	people := ps.repository.GetAll()

	return people
}

func (ps *PersonServiceImpl) Update(id int32, payload entity.Person) entity.Person {
	person := ps.repository.Update(id, payload)

	return person
}

func (ps *PersonServiceImpl) Delete(id int32) bool {
	success := ps.repository.Delete(id)

	return success
}

func ProvidePersonService(r repository.PersonRepository) *PersonServiceImpl {
	return &PersonServiceImpl{
		repository: r,
	}
}
