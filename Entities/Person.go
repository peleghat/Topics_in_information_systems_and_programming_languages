package Entities

type Person struct {
	id                          string
	name                        string
	email                       string
	favoriteProgrammingLanguage string
}

func NewPerson(name string, email string, favoriteProgrammingLanguage string) *Person {
	//TODO id
	id := "1"
	return &Person{id: id, name: name, email: email, favoriteProgrammingLanguage: favoriteProgrammingLanguage}
}
