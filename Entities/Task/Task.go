package Task

type Task struct {
	id      string
	ownerId string
	status  Status
}

func NewTask(ownerId string, status Status) *Task {
	//TODO id
	id := "1"
	return &Task{id: id, ownerId: ownerId, status: status}
}
