package APIFolder

import (
	"encoding/json"
	"fmt"
	"miniProject/EntitiesFolder"
	"miniProject/ErrorsFolder"
	"miniProject/dbFolder"
	"net/http"
)

func AddPerson(w http.ResponseWriter, r *http.Request) {
	var newPersonInput EntitiesFolder.PersonInput
	json.NewDecoder(r.Body).Decode(&newPersonInput)
	err := EntitiesFolder.IsValidEmail(newPersonInput)
	if err != nil {
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Errorf("a person with email %s contains illegal values", newPersonInput.Email).Error()))
		return
	}
	p := EntitiesFolder.NewPerson(newPersonInput.Name, newPersonInput.Email, newPersonInput.FavProg)
	dbErr := dbFolder.InsertPerson(p)
	if dbErr != nil {
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusBadRequest)
		switch dbErr {
		case ErrorsFolder.ErrDbConnection:
			{
				w.Write([]byte(fmt.Errorf("failed to connect to db").Error()))
			}
		case ErrorsFolder.ErrAlreadyExist:
			{
				w.Write([]byte(fmt.Errorf("a person with email %s already exists", newPersonInput.Email).Error()))
			}
		default:
			{
				w.Write([]byte(fmt.Errorf("unknown error has occured").Error()))
			}
		}
		return
	} else { // success
		w.Header().Set("Location", fmt.Sprintf("/api/people/%s", p.GetId()))
		w.Header().Set("x-Created-Id", p.GetId())
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte("person created successfuly"))
		return
	}
}

func GetAllPersons(w http.ResponseWriter, r *http.Request) {

}
func GetPerson(w http.ResponseWriter, r *http.Request) {

}
func UpdatePerson(w http.ResponseWriter, r *http.Request) {

}
func DeletePerson(w http.ResponseWriter, r *http.Request) {

}
func GetAllPersonsTasks(w http.ResponseWriter, r *http.Request) {

}
func AddTaskToPerson(w http.ResponseWriter, r *http.Request) {

}
func GetPersonsTasksByStatus(w http.ResponseWriter, r *http.Request) {

}
func GetTask(w http.ResponseWriter, r *http.Request) {

}
func UpdateTask(w http.ResponseWriter, r *http.Request) {

}
func DeleteTask(w http.ResponseWriter, r *http.Request) {

}
func GetTaskStatus(w http.ResponseWriter, r *http.Request) {

}
func SetTaskStatus(w http.ResponseWriter, r *http.Request) {

}
func GetTaskOwner(w http.ResponseWriter, r *http.Request) {

}
func SetTaskOwner(w http.ResponseWriter, r *http.Request) {

}
