package EntitiesFolder

import (
	"github.com/google/uuid"
)

type Task struct {
	Id          string
	OwnerId     string
	Status      Status
	TaskType    string
	Description string
}

// Task functions
//  Constructor

func NewTask(ownerId string, status Status, taskType string, Description string) Task {
	id := uuid.New()
	return Task{Id: id.String(), OwnerId: ownerId, Status: status, TaskType: taskType, Description: Description}
}

// Getters

func (t Task) GetId() string {
	return t.Id
}
func (t Task) GetOwnerId() string {
	return t.OwnerId
}
func (t Task) GetStatus() string {
	return string(t.Status)
}
func (t Task) GetStatusStr() string {
	return string(t.Status)
}
func (t Task) GetTaskType() string {
	return t.TaskType
}
func (t Task) GetDescription() string {
	return t.Description
}

// Setters

func (t Task) SetId(id string) {
	t.Id = id
}
func (t Task) SetOwnerId(newOwnerId string) {
	t.OwnerId = newOwnerId
}
func (t Task) SetStatus(newStatus Status) {
	t.Status = newStatus
}
func (t Task) SetTaskType(taskType string) {
	t.TaskType = taskType
}
func (t Task) SetDescription(Description string) {
	t.Description = Description
}
