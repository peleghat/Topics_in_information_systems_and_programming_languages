package EntitiesFolder

import "github.com/google/uuid"

type Person struct {
	id                          string
	name                        string
	email                       string
	favoriteProgrammingLanguage string
}

func NewPerson(name string, email string, favoriteProgrammingLanguage string) *Person {
	id := uuid.New()
	return &Person{id: id.String(), name: name, email: email, favoriteProgrammingLanguage: favoriteProgrammingLanguage}
}

// TODO - addTask(t), listTasks()
