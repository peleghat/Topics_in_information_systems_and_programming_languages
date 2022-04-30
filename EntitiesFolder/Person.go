package EntitiesFolder

import "github.com/google/uuid"

type Person struct {
	id      string `json:"id"`
	name    string `json:"name"`
	email   string `json:"email"`
	favProg string `json:"favProg"`
}

func NewPerson(name string, email string, favProg string) *Person {
	id := uuid.New()
	return &Person{id: id.String(), name: name, email: email, favProg: favProg}
}

// TODO - addTask(t), listTasks()
