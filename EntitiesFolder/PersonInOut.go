package EntitiesFolder

import (
	"miniProject/ErrorsFolder"
	"strings"
	_ "strings"
)

type PersonInput struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	FavProg string `json:"favoriteProgrammingLanguage"`
}

// NewPersonInput Constructor
func NewPersonInput(name string, email string, favProg string) PersonInput {
	return PersonInput{Name: name, Email: email, FavProg: favProg}
}

// IsValidEmail function check if the persons email is valid
func IsValidEmail(p PersonInput) error {
	if strings.Contains(p.Email, "@") && strings.Contains(p.Email, ".") {
		return nil
	} else {
		return ErrorsFolder.ErrIllegalValues
	}
}
