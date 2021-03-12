package entity

// Person is a dummy model
type Person struct {
	ID   int32
	Name string
	Age  int32
}

// ProvidePerson provides a person model
func ProvidePerson() Person {
	return Person{}
}
