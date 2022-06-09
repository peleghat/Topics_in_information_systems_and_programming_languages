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

type PersonOutput struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	FavProg     string `json:"favoriteProgrammingLanguage"`
	ActiveTasks int    `json:"activeTaskCount"`
}

func PersonToOutput(person Person) PersonOutput {
	var output PersonOutput
	output.ID = person.GetId()
	output.Name = person.Name
	output.Email = person.Email
	output.ActiveTasks = person.GetActiveTaskCount()
	output.FavProg = person.GetFavProg()
	return output
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

func PersonsToOutput(persons []Person) []PersonOutput {
	var result []PersonOutput
	for _, person := range persons {
		result = append(result, PersonToOutput(person))
	}
	return result
}
