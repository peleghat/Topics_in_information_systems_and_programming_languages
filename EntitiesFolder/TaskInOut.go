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

type ChoreOutput struct {
	Id          string `json:"id"`
	TaskType    string `json:"type"`
	OwnerId     string `json:"ownerId"`
	Status      string `json:"status"`
	Description string `json:"description"`
	Size        string `json:"size"`
}

type HomeWorkOutput struct {
	Id          string `json:"id"`
	TaskType    string `json:"type"`
	OwnerId     string `json:"ownerId"`
	Status      string `json:"status"`
	Description string `json:"details"`
	Course      string `json:"course"`
	DueDate     string `json:"dueDate"`
}

func ChoreToChoreOutPut(c Chore) ChoreOutput {
	var cOutPut ChoreOutput
	cOutPut.Id = c.GetTask().GetId()
	cOutPut.TaskType = c.GetTask().GetTaskType()
	cOutPut.OwnerId = c.GetTask().GetOwnerId()
	cOutPut.Status = StatusToString(c.GetTask().Status)
	cOutPut.Description = c.GetTask().GetDescription()
	cOutPut.Size = SizeToString(c.GetSize())
	return cOutPut
}

func HomeWorkToHomeWorkOutPut(h HomeWork) HomeWorkOutput {
	var houtPut HomeWorkOutput
	houtPut.Id = h.GetTask().GetId()
	houtPut.TaskType = h.GetTask().GetTaskType()
	houtPut.OwnerId = h.GetTask().GetOwnerId()
	houtPut.Status = StatusToString(h.GetTask().Status)
	houtPut.Description = h.GetTask().GetDescription()
	houtPut.Course = h.GetCourse()
	houtPut.DueDate = h.GetDueDate().Format("2006-01-02")
	return houtPut
}

func ChoreListToChoreOutPutList(choreList []Chore) []ChoreOutput {
	var OutPutList []ChoreOutput
	for _, chore := range choreList {
		OutPutList = append(OutPutList, ChoreToChoreOutPut(chore))
	}
	return OutPutList
}

func HomeWorkListToHomeWorkOutPutList(homeWorkList []HomeWork) []HomeWorkOutput {
	var OutPutList []HomeWorkOutput
	for _, homework := range homeWorkList {
		OutPutList = append(OutPutList, HomeWorkToHomeWorkOutPut(homework))
	}
	return OutPutList
}
