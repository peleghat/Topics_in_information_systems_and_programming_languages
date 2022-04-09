package Task

import "time"

type HomeWork struct {
	course  string
	dueDate time.Time
	details string
	Task
}

func NewHomeWork(course string, dueDate time.Time, details string) *HomeWork {
	return &HomeWork{course: course, dueDate: dueDate, details: details}
}
