package repository

import (
	"github.com/Jeeo/golang-ddd-boilerplate/internal/domain/entity"
)

// PersonRepository is a interface for person repository
type PersonRepository interface {
	Create(entity.Person) entity.Person
	GetOne(int32) entity.Person
	GetAll() []entity.Person
	Update(int32, entity.Person) entity.Person
	Delete(int32) bool
}
