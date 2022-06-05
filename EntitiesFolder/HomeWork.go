package EntitiesFolder

import (
	"time"
)

type HomeWork struct {
	Course  string
	DueDate time.Time
	Task    Task
}

// Homework functions
// Constructor

func NewHomeWork(course string, dueDate time.Time, task Task) HomeWork {
	return HomeWork{Course: course, DueDate: dueDate, Task: task}
}

// Getters

func (h *HomeWork) GetCourse() string {
	return h.Course
}
func (h *HomeWork) GetDueDate() time.Time {
	return h.DueDate
}
func (h *HomeWork) GetTask() Task {
	return h.Task
}

// Setters

func (h *HomeWork) SetCourse(course string) {
	h.Course = course
}
func (h *HomeWork) SetDueDate(dueDate time.Time) {
	h.DueDate = dueDate
}
func (h *HomeWork) SetTask(task Task) {
	h.Task = task
}

// ClockUpdate function gets a string and creates a time.Time instance
func ClockUpdate(toUpdate string) time.Time {
	myClock, err := time.Parse("2006-01-02", toUpdate)
	if err != nil {
		panic(err)
	}
	return myClock
}
