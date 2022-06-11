package EntitiesFolder

import (
	"miniProject/ErrorsFolder"
	"strings"
	_ "strings"
)

type PersonInput struct {
	Name    string `json:"name"`
	Email   string `json:"emails"`
	FavProg string `json:"favoriteProgrammingLanguage"`
}

func NewPersonInput(name string, email string, favProg string) PersonInput {
	return PersonInput{Name: name, Email: email, FavProg: favProg}
}

func IsValidEmail(p PersonInput) error {
	if strings.Contains(p.Email, "@") && strings.Contains(p.Email, ".") {
		return nil
	} else {
		return ErrorsFolder.ErrIllegalValues
	}
}
