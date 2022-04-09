package Chore

type chore struct {
	description string
	size        Size
}

func newChore(description string, size Size) *chore {
	return &chore{description: description, size: size}
}
