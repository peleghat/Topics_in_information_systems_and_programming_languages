package EntitiesFolder

// TODO - addTask(t) - Using IncActiveTaskCount and DecActiveTaskCount, listTasks()
import (
	"github.com/google/uuid"
)

type Person struct {
	Id              string `json:"id"`
	Name            string `json:"name"`
	Email           string `json:"email"`
	FavProg         string `json:"favProg"`
	ActiveTaskCount int    `json:"activeTaskCount"`
}

// Constructor

func NewPerson(name string, email string, favProg string) Person {
	id := uuid.New()
	return Person{Id: id.String(), Name: name, Email: email, FavProg: favProg, ActiveTaskCount: 0}
}

// Getters

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

// Setters

func (p Person) SetId(_Id string) {
	p.Id = _Id
}
func (p Person) SetName(_Name string) {
	p.Name = _Name
}
func (p Person) SetEmail(_Email string) {
	p.Email = _Email
}
func (p Person) SetFavProg(_FavProg string) {
	p.FavProg = _FavProg
}
func (p Person) SetActiveTaskCount(_ActiveTaskCount int) {
	p.ActiveTaskCount = _ActiveTaskCount
}

// IncActiveTaskCount Increment Active Task counter
func (p Person) IncActiveTaskCount() {
	p.ActiveTaskCount++
}

// DecActiveTaskCount Decrement Active Task counter
func (p Person) DecActiveTaskCount() {
	p.ActiveTaskCount--
}
