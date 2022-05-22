package EntitiesFolder

type Status int

const (
	Active        Status = iota // 0
	Done                        // 1
	UnknownStatus = -1
)

func createStatus(s string) Status {
	if s == "Active" {
		return Active
	} else if s == "Done" {
		return Done
	} else {
		return UnknownStatus
	}
}
