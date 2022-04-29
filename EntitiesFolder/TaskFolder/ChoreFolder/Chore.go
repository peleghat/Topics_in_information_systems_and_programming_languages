package ChoreFolder

import "miniProject/EntitiesFolder/TaskFolder"

type chore struct {
	description string
	size        Size
	TaskFolder.Task
}

func newChore(description string, size Size) *chore {
	return &chore{description: description, size: size}
}
