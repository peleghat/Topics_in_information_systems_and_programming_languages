package EntitiesFolder

import (
	"github.com/google/uuid"
	"time"
)

type Task struct {
	Id          string
	OwnerId     string
	Status      Status
	TaskType    string
	Description string
}

type HomeWork struct {
	Course  string
	DueDate time.Time
	Task    Task
}

type Chore struct {
	Size Size
	Task Task
}

//Task functions

func NewTask(ownerId string, status Status, taskType string, Description string) Task {
	id := uuid.New()
	return Task{Id: id.String(), OwnerId: ownerId, Status: status, TaskType: taskType, Description: Description}
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

func (t *Task) GetTaskType() string {
	return t.TaskType
}

func (t *Task) SetTaskType(taskType string) {
	t.TaskType = taskType
}

func (t *Task) GetDescription() string {
	return t.Description
}

func (t *Task) SetDescription(Description string) {
	t.Description = Description
}

//Chore functions

func NewChore(size Size, task Task) Chore {
	return Chore{Size: size, Task: task}
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

func NewHomeWork(course string, dueDate time.Time, task Task) HomeWork {
	return HomeWork{Course: course, DueDate: dueDate, Task: task}
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
