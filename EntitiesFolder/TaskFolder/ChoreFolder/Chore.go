package ChoreFolder

import "miniProject/EntitiesFolder/TaskFolder"

type Chore struct {
	description string
	size        Size
	TaskFolder.Task
}

//func NewChore(description string, size Size) *chore {
//	return &chore{description: description, size: size}
//}

func (c *Chore) Description() string {
	return c.description
}

func (c *Chore) SetDescription(description string) {
	c.description = description
}

func (c *Chore) Size() Size {
	return c.size
}

func (c *Chore) SetSize(size Size) {
	c.size = size
}
