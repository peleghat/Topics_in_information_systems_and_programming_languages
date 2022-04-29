package TaskFolder

import "github.com/google/uuid"

type Task struct {
	id      string
	ownerId string
	status  Status
}

func NewTask(ownerId string, status Status) *Task {
	id := uuid.New()
	return &Task{id: id.String(), ownerId: ownerId, status: status}
}

func setOwner(tsk *Task, newOwnerId string) {
	tsk.ownerId = newOwnerId
}

func setStatus(tsk *Task, newStatus Status) {
	tsk.status = newStatus
}
