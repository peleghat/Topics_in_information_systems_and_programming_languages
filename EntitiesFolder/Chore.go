package EntitiesFolder

type Chore struct {
	Size Size
	Task Task
}

// Chore functions
// Constructor

func NewChore(size Size, task Task) Chore {
	return Chore{Size: size, Task: task}
}

// Getters

func (c *Chore) GetSize() Size {
	return c.Size
}
func (c *Chore) GetTask() Task {
	return c.Task
}

// Setters

func (c *Chore) SetSize(size Size) {
	c.Size = size
}
func (c *Chore) SetTask(task Task) {
	c.Task = task
}
