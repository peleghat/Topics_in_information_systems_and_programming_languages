package APIFolder

// TODO - Add comments
// TODO - Errors, and comments for the 4 functions we wrote

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"miniProject/EntitiesFolder"
	"miniProject/ErrorsFolder"
	"miniProject/dbFolder"
	"net/http"
)

// AddPerson function adds a person to the database.
// Checks if the person email is unique, if not return error
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

// GetAllPersons functions return all persons in the db.
func GetAllPersons(w http.ResponseWriter, r *http.Request) {
	err, persons := dbFolder.GetAllPersons()
	if err != nil {
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Errorf("query to the db has failed").Error()))
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(EntitiesFolder.PersonsToOutput(persons))
	}
}

func GetPerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	err, person := dbFolder.GetPerson(params["id"])
	if err != nil {
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Errorf("A person with the id %s does not exist", params["id"]).Error()))
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(EntitiesFolder.PersonToOutput(person))
	}
}

func UpdatePerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var PersonInput EntitiesFolder.PersonInput
	json.NewDecoder(r.Body).Decode(&PersonInput)
	err1, PersonToUpdate := dbFolder.GetPerson(params["id"])
	if err1 != nil {
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Errorf("A person with the id %s does not exist", params["id"]).Error()))
	}
	if PersonInput.Name != "" {
		PersonToUpdate.Name = PersonInput.Name
	}
	if PersonInput.Email != "" {
		PersonToUpdate.Email = PersonInput.Email
	}
	if PersonInput.FavProg != "" {
		PersonToUpdate.FavProg = PersonInput.FavProg
	}
	err2 := dbFolder.UpdatePerson(PersonToUpdate)
	if err2 != nil {
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusBadRequest)
		switch err2 {
		case ErrorsFolder.ErrDbConnection:
			{
				w.Write([]byte(fmt.Errorf("failed to connect to db").Error()))
			}
		case ErrorsFolder.ErrDbQuery:
			{
				w.Write([]byte(fmt.Errorf("failed to update the person with id %s to db", params["id"]).Error()))
			}
		default:
			{
				w.Write([]byte(fmt.Errorf("unknown error has occured").Error()))
			}
		}
	} else {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Person updated successfully. Response body contains updated data.\n"))
		json.NewEncoder(w).Encode(EntitiesFolder.PersonToOutput(PersonToUpdate))
	}
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
