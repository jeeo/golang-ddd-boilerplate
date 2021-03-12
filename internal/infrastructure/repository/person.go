package repository

import (
	"jeeo/api-boilerplate/internal/domain/entity"
	"jeeo/api-boilerplate/internal/infrastructure/persistence"
	"log"
)

// PersonRepository is self explanatory
type PersonRepository struct {
	Database *persistence.Database
}

// GetOne returns a persisted Person
func (pr *PersonRepository) GetOne(id int32) entity.Person {
	person := entity.Person{}
	statement := "SELECT id, name, age from person where id = $1"
	preparedStatement, err := pr.Database.Conn.Prepare(statement)
	if err != nil {
		log.Println("Error on prepare statement: ", err.Error())
		return entity.Person{}
	}
	err = preparedStatement.QueryRow(id).Scan(&person.ID, &person.Name, &person.Age)
	if err != nil {
		log.Println("Error on Query: ", err.Error())
		return entity.Person{}
	}

	return person
}

// GetAll returns all persisted Person
func (pr *PersonRepository) GetAll() []entity.Person {
	people := []entity.Person{}
	statement := "SELECT id, name, age from person"
	rows, err := pr.Database.Conn.Query(statement)
	if err != nil {
		log.Println("Error on Query: ", err.Error())
		return nil
	}
	for rows.Next() {
		person := entity.Person{}
		if err = rows.Scan(&person.ID, &person.Name, &person.Age); err != nil {
			log.Println("Error on scan: ", err.Error())
			return nil
		}
		people = append(people, person)
	}
	return people
}

// Create inserts a person into Database
func (pr *PersonRepository) Create(person entity.Person) entity.Person {
	p := entity.Person{}
	statement := "INSERT INTO person (name, age) VALUES ($1, $2) returning id, name, age"
	preparedStatement, err := pr.Database.Conn.Prepare(statement)

	if err = preparedStatement.QueryRow(person.Name, person.Age).Scan(&p.ID, &p.Name, &p.Age); err != nil {
		return entity.Person{}
	}

	return p
}

// Update updates an person
func (pr *PersonRepository) Update(person entity.Person) entity.Person {
	p := entity.Person{}
	statement := "UPDATE person SET name = $1, age = $2 where id = $3 returning id, name, age"
	preparedStatement, err := pr.Database.Conn.Prepare(statement)

	if err = preparedStatement.QueryRow(person.Name, person.Age, person.ID).Scan(&p.ID, &p.Name, &p.Age); err != nil {
		return entity.Person{}
	}

	return p
}

// Delete deeltes a persisted Person
func (pr *PersonRepository) Delete(id int32) bool {
	statement := "DELETE FROM person WHERE id = $1"
	preparedStatement, err := pr.Database.Conn.Prepare(statement)

	if _, err = preparedStatement.Exec(id); err != nil {
		return false
	}

	return true
}

// ProvidePersonRepository providers person repository
func ProvidePersonRepository(database *persistence.Database) *PersonRepository {
	return &PersonRepository{
		Database: database,
	}
}
