package mapper

import (
	"reflect"
	"testing"

	"github.com/Jeeo/golang-ddd-boilerplate/internal/domain/entity"
	"github.com/Jeeo/golang-ddd-boilerplate/internal/entrypoint/dto"
)

func TestToDTO(t *testing.T) {
	target := &PersonMapperImpl{}
	somePerson := entity.Person{ID: 123, Name: "Bjorn", Age: 39}
	expected := dto.PersonDTO{ID: 123, Name: "Bjorn", Age: 39}

	result := target.ToDTO(somePerson)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("\nRECEIVED %#v\nEXPECTED: %#v", result, expected)
	}
}

func TestFromDTO(t *testing.T) {
	target := &PersonMapperImpl{}
	someDTO := dto.PersonDTO{ID: 123, Name: "Bjorn", Age: 39}
	expected := entity.Person{ID: 123, Name: "Bjorn", Age: 39}

	result := target.FromDTO(someDTO)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("\nRECEIVED %#v\nEXPECTED: %#v", result, expected)
	}
}

func TestToManyDTO(t *testing.T) {
	target := &PersonMapperImpl{}
	somePeople := []entity.Person{
		{ID: 1, Name: "Aragon", Age: 89},
		{ID: 2, Name: "Legolas", Age: 301},
	}
	expected := []dto.PersonDTO{
		{ID: 1, Name: "Aragon", Age: 89},
		{ID: 2, Name: "Legolas", Age: 301},
	}

	result := target.ToManyDTO(somePeople)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("\nRECEIVED %#v\nEXPECTED: %#v", result, expected)
	}
}

func TestProvidePersonMapper(t *testing.T) {
	expected := &PersonMapperImpl{}

	result := ProvidePersonMapper()

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("\nRECEIVED %#v\nEXPECTED: %#v", result, expected)
	}
}
