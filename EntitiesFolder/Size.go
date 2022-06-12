package EntitiesFolder

type Size int

const (
	Small       Size = iota // 0
	Medium                  // 1
	Large                   // 2
	UnknownSize = -1
)

func CreateSize(s string) Size {
	if s == "Small" {
		return Small
	} else if s == "Medium" {
		return Medium
	} else if s == "Large" {
		return Large
	} else {
		return UnknownSize
	}
}

func SizeToString(s Size) string {
	if s == 0 {
		return "Small"
	} else if s == 1 {
		return "Medium"
	} else if s == 2 {
		return "Large"
	} else {
		return "UnknownSize"
	}
}
