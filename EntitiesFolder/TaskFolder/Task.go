package TaskFolder

type Task struct {
	id      string
	ownerId string
	status  Status
}

//func NewTask(ownerId string, status Status) *Task {
//	id := uuid.New()
//	return &Task{id: id.String(), ownerId: ownerId, status: status}
//}

func (t *Task) GetId() string {
	return t.id
}

func (t *Task) SetId(id string) {
	t.id = id
}

func (t *Task) GetOwnerId() string {
	return t.ownerId
}

func (t *Task) SetOwnerId(newOwnerId string) {
	t.ownerId = newOwnerId
}

func (t *Task) GetStatus() Status {
	return t.status
}

func (t *Task) SetStatus(newStatus Status) {
	t.status = newStatus
}
