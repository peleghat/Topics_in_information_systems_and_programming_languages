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
		json.NewEncoder(w).Encode(persons)
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
		json.NewEncoder(w).Encode(person)
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
		json.NewEncoder(w).Encode(PersonToUpdate)
	}
}

func DeletePerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	choreList, HomeWorkList, _ := dbFolder.GetTasksFromPerson(params["id"])
	for _, chore := range choreList {
		dbFolder.DeleteTask(chore.GetTask())
	}
	for _, homeWork := range HomeWorkList {
		dbFolder.DeleteTask(homeWork.GetTask())
	}
	err2 := dbFolder.DeletePerson(params["id"])
	if err2 != nil {
		w.Header().Set("Content-Type", "text/plain")
		switch err2 {
		case ErrorsFolder.ErrDbConnection:
			{
				w.Write([]byte(fmt.Errorf("failed to connect to db").Error()))
			}
		case ErrorsFolder.ErrDbQuery:
			{
				w.Write([]byte(fmt.Errorf("failed to update the person with id %s to db", params["id"]).Error()))
			}
		}
	} else {
		w.WriteHeader(http.StatusOK)
	}
}

func GetAllPersonTasks(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	choreList, homeworkList, err := dbFolder.GetTasksFromPerson(params["id"])
	if err != nil {
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Errorf("A person with the id %s does not exist", params["id"]).Error()))
	} else {
		ans := []interface{}{choreList, homeworkList}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(ans)
	}
}

func AddTaskToPerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var TaskToAdd EntitiesFolder.TaskInput
	json.NewDecoder(r.Body).Decode(&TaskToAdd)
	if TaskToAdd.TaskType == "chore" || TaskToAdd.TaskType == "Chore" {
		choreToAdd := EntitiesFolder.TaskToChore(TaskToAdd, params["id"])
		err := dbFolder.AddChore(choreToAdd)
		switch err {
		case ErrorsFolder.ErrDbConnection:
			w.Header().Set("Content-Type", "text/plain")
			w.WriteHeader(http.StatusBadRequest) //400
			w.Write([]byte(fmt.Errorf("failed to connect to db").Error()))
		case ErrorsFolder.ErrDbQuery:
			w.Header().Set("Content-Type", "text/plain")
			w.WriteHeader(http.StatusNotFound) //404
			w.Write([]byte(fmt.Errorf("A person with the id %s does not exist", params["id"]).Error()))
			w.Write([]byte("Requested person is not present.\n"))
		}
		w.Header().Set("Location", fmt.Sprintf("/api/people/%s/tasks", choreToAdd.GetTask().GetId()))
		w.Header().Set("x-Created-Id", choreToAdd.GetTask().GetId())
		w.Write([]byte("Task created and assigned successfully\n"))
		w.WriteHeader(http.StatusCreated)
	} else {
		HomeWorkToAdd := EntitiesFolder.TaskToHomework(TaskToAdd, params["id"])
		err := dbFolder.AddHomeWork(HomeWorkToAdd)
		switch err {
		case ErrorsFolder.ErrDbConnection:
			w.Header().Set("Content-Type", "text/plain")
			w.WriteHeader(http.StatusBadRequest) //400
			w.Write([]byte(fmt.Errorf("failed to connect to db").Error()))
		case ErrorsFolder.ErrDbQuery:
			w.Header().Set("Content-Type", "text/plain")
			w.WriteHeader(http.StatusNotFound) //404
			w.Write([]byte(fmt.Errorf("A person with the id %s does not exist", params["id"]).Error()))
			w.Write([]byte("Requested person is not present.\n"))
		default:
			{
				w.Write([]byte(fmt.Errorf("unknown error has occured").Error()))
			}
		}
		w.Header().Set("Location", fmt.Sprintf("/api/tasks/%s", HomeWorkToAdd.GetTask().GetId()))
		w.Header().Set("x-Created-Id", HomeWorkToAdd.GetTask().GetId())
		w.WriteHeader(http.StatusCreated)

	}

}

func GetPersonsTasksByStatus(w http.ResponseWriter, r *http.Request) {
	// lol
}

// TODO - CHANGE ACTIVE and SIZE
func GetTask(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	Chore, HomeWork, err := dbFolder.GetTask(params["id"])
	if err != nil {
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusBadRequest)
		switch err {
		case ErrorsFolder.ErrDbConnection:
			{
				w.Write([]byte(fmt.Errorf("failed to connect to db").Error()))
			}
		case ErrorsFolder.ErrIllegalValues:
			{
				w.Write([]byte(fmt.Errorf("A task with the id %s does not exist", params["id"]).Error()))
			}
		}
	} else {
		w.WriteHeader(http.StatusOK)
		emptyChore := EntitiesFolder.Chore{}
		if Chore != emptyChore {
			json.NewEncoder(w).Encode(Chore)
		} else {
			json.NewEncoder(w).Encode(HomeWork)
		}
	}
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
