package EntitiesFolder

type Status int

const (
	Active        Status = iota // 0
	Done                        // 1
	UnknownStatus = -1
)

func createStatus(s string) Status {
	if s == "Active" || s == "active" {
		return Active
	} else if s == "Done" || s == "done" {
		return Done
	} else {
		return UnknownStatus
	}
}
