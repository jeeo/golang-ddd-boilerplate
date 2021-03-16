package repository

import (
	"github.com/Jeeo/golang-ddd-boilerplate/internal/domain/entity"
)

// PersonRepository is a interface for person repository
type PersonRepository interface {
	CreatePerson(entity.Person) entity.Person
	GetPerson(int32) entity.Person
	GetPeople() []entity.Person
	UpdatePerson(int32, entity.Person) entity.Person
	DeletePerson(int32) bool
}
