package EntitiesFolder

import "github.com/google/uuid"

type Person struct {
	Id              string `json:"id"`
	Name            string `json:"name"`
	Email           string `json:"email"`
	FavProg         string `json:"favProg"`
	ActiveTaskCount int    `json:"activeTaskCount"`
}

func NewPerson(name string, email string, favProg string) Person {
	id := uuid.New()
	return Person{Id: id.String(), Name: name, Email: email, FavProg: favProg, ActiveTaskCount: 0}
}

// TODO - addTask(t), listTasks()

func (p Person) GetId() string {
	return p.Id
}
func (p Person) GetName() string {
	return p.Name
}
func (p Person) GetEmail() string {
	return p.Email
}
func (p Person) GetFavProg() string {
	return p.FavProg
}

func (p Person) GetActiveTaskCount() int {
	return p.ActiveTaskCount
}

func (p Person) IncActiveTaskCount() {
	p.ActiveTaskCount++
}

func (p Person) DecActiveTaskCount() {
	p.ActiveTaskCount--
}

func (p Person) SetActiveTaskCount(NewActiveTaskCount int) {
	p.ActiveTaskCount = NewActiveTaskCount
}
