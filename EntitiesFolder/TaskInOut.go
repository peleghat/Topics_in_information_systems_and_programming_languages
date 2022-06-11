package EntitiesFolder

import "time"

type TaskInput struct {
	Status      string    `json:"status"`
	TaskType    string    `json:"type"`
	Description string    `json:"description"`
	Size        string    `json:"size,omitempty"`
	Course      string    `json:"course,omitempty"`
	DueDate     time.Time `json:"duedate,omitempty"`
}

func TaskToChore(t TaskInput, ownerid string) Chore {
	taskOutput := NewTask(ownerid, createStatus(t.Status), t.TaskType, t.Description)
	output := NewChore(createSize(t.Size), taskOutput)
	return output
}

func TaskToHomework(t TaskInput, ownerid string) HomeWork {
	taskOutput := NewTask(ownerid, createStatus(t.Status), t.TaskType, t.Description)
	output := NewHomeWork(t.Course, t.DueDate, taskOutput)
	return output
}
