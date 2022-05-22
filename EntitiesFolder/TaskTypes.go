package EntitiesFolder

import (
	"github.com/google/uuid"
	"time"
)

type Task struct {
	Id      string
	OwnerId string
	Status  Status
}

type Chore struct {
	Description string
	Size        Size
	Task        Task
}

type HomeWork struct {
	Course  string
	DueDate time.Time
	Details string
	Task    Task
}

//Task functions

func NewTask(ownerId string, status Status) Task {
	id := uuid.New()
	return Task{Id: id.String(), OwnerId: ownerId, Status: status}
}

func (t *Task) GetId() string {
	return t.Id
}

func (t *Task) SetId(id string) {
	t.Id = id
}

func (t Task) GetOwnerId() string {
	return t.OwnerId
}

func (t *Task) SetOwnerId(newOwnerId string) {
	t.OwnerId = newOwnerId
}

func (t *Task) GetStatus() Status {
	return t.Status
}

func (t *Task) SetStatus(newStatus Status) {
	t.Status = newStatus
}

//Chore functions

func NewChore(description string, size Size, task Task) Chore {
	return Chore{Description: description, Size: size, Task: task}
}

func (c *Chore) GetDescription() string {
	return c.Description
}

func (c *Chore) SetDescription(description string) {
	c.Description = description
}

func (c *Chore) GetSize() Size {
	return c.Size
}

func (c *Chore) SetSize(size Size) {
	c.Size = size
}

//Homework functions

func ClockUpdate(toUpdate string) time.Time {
	myClock, err := time.Parse("2006-01-02", toUpdate)
	if err != nil {
		panic(err)
	}
	return myClock
}

func NewHomeWork(course string, dueDate string, details string, task Task) HomeWork {
	return HomeWork{Course: course, DueDate: ClockUpdate(dueDate), Details: details, Task: task}
}

func (h *HomeWork) GetCourse() string {
	return h.Course
}

func (h *HomeWork) SetCourse(course string) {
	h.Course = course
}

func (h *HomeWork) GetDueDate() time.Time {
	return h.DueDate
}

func (h *HomeWork) SetDueDate(dueDate time.Time) {
	h.DueDate = dueDate
}

func (h *HomeWork) GetDetails() string {
	return h.Details
}

func (h *HomeWork) SetDetails(details string) {
	h.Details = details
}
