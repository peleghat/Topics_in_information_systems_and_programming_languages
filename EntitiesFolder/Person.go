package EntitiesFolder

import "github.com/google/uuid"

type Person struct {
	id                          string `json:"id"`
	name                        string `json:"name"`
	email                       string `json:"email"`
	favoriteProgrammingLanguage string `json:"favoriteProgrammingLanguage"`
}

func NewPerson(name string, email string, favoriteProgrammingLanguage string) *Person {
	id := uuid.New()
	return &Person{id: id.String(), name: name, email: email, favoriteProgrammingLanguage: favoriteProgrammingLanguage}
}

// TODO - addTask(t), listTasks()
