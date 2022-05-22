package TaskFolder

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

func (h *HomeWork) Course() string {
	return h.course
}

func (h *HomeWork) SetCourse(course string) {
	h.course = course
}

func (h *HomeWork) DueDate() time.Time {
	return h.dueDate
}

func (h *HomeWork) SetDueDate(dueDate time.Time) {
	h.dueDate = dueDate
}

func (h *HomeWork) Details() string {
	return h.details
}

func (h *HomeWork) SetDetails(details string) {
	h.details = details
}
