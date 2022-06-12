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
	"strings"
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
		w.Write([]byte(fmt.Errorf("A person with the id %s does not exist", params["id"]).Error()))
		w.WriteHeader(http.StatusBadRequest)

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
		w.Write([]byte("Person updated successfully. Response body contains updated data.\n"))
		json.NewEncoder(w).Encode(PersonToUpdate)
		w.WriteHeader(http.StatusOK)

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
		// TODO - add all 4 cases to interface
		ans := []interface{}{EntitiesFolder.ChoreListToChoreOutPutList(choreList),
			EntitiesFolder.HomeWorkListToHomeWorkOutPutList(homeworkList)}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(ans)
	}
}

// AddTaskToPerson fix add task, active count ++, if not word "active" do aturomatically delete fmt.println
func AddTaskToPerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var TaskToAdd EntitiesFolder.TaskInput
	json.NewDecoder(r.Body).Decode(&TaskToAdd)
	if TaskToAdd.TaskType == "chore" || TaskToAdd.TaskType == "Chore" {
		choreToAdd := EntitiesFolder.TaskToChore(TaskToAdd, params["id"])
		err := dbFolder.AddChore(choreToAdd)
		if err != nil {
			w.Header().Set("Content-Type", "text/plain")
			switch err {
			case ErrorsFolder.ErrDbConnection:
				{
					w.Write([]byte(fmt.Errorf("failed to connect to db").Error()))
					w.WriteHeader(http.StatusBadRequest) //400
					return
				}
			case ErrorsFolder.ErrDbQuery:
				{
					w.Write([]byte(fmt.Errorf("A person with the id %s does not exist", params["id"]).Error()))
					w.Write([]byte("Requested person is not present.\n"))
					w.WriteHeader(http.StatusNotFound) //404
					return
				}
			default:
				{
					w.Write([]byte(fmt.Errorf("unknown error has occured").Error()))
				}
			}
		}
		w.Header().Set("Location", fmt.Sprintf("/api/people/%s/tasks", choreToAdd.GetTask().GetId()))
		w.Header().Set("x-Created-Id", choreToAdd.GetTask().GetId())
		w.Write([]byte("Task created and assigned successfully\n"))
		w.WriteHeader(http.StatusCreated)
		return
	} else if TaskToAdd.TaskType == "homework" || TaskToAdd.TaskType == "Homework" {
		homeworkToAdd := EntitiesFolder.TaskToHomework(TaskToAdd, params["id"])
		err := dbFolder.AddHomeWork(homeworkToAdd)
		if err != nil {
			w.Header().Set("Content-Type", "text/plain")
			switch err {
			case ErrorsFolder.ErrDbConnection:
				{
					w.Write([]byte(fmt.Errorf("failed to connect to db").Error()))
					w.WriteHeader(http.StatusBadRequest) //400
					return
				}
			case ErrorsFolder.ErrDbQuery:
				{
					w.Write([]byte(fmt.Errorf("A person with the id %s does not exist", params["id"]).Error()))
					w.Write([]byte("Requested person is not present.\n"))
					w.WriteHeader(http.StatusNotFound) //404
					return
				}
			default:
				{
					w.Write([]byte(fmt.Errorf("unknown error has occured").Error()))
				}
			}
		}
		w.Header().Set("Location", fmt.Sprintf("/api/people/%s/tasks", homeworkToAdd.GetTask().GetId()))
		w.Header().Set("x-Created-Id", homeworkToAdd.GetTask().GetId())
		w.Write([]byte("Task created and assigned successfully\n"))
		w.WriteHeader(http.StatusCreated)
		return
	} else {
		w.Write([]byte(fmt.Errorf("Task type: %s does not exist", TaskToAdd.TaskType).Error()))
		w.WriteHeader(http.StatusNotFound) //404
		return
	}
}

func GetPersonsTasksByStatus(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	fmt.Println(params["status"])
	choreList, homeworkList, err := dbFolder.GetTasksFromPerson(params["id"])
	choreOut := EntitiesFolder.ChoreListToChoreOutPutList(choreList)
	HomeOut := EntitiesFolder.HomeWorkListToHomeWorkOutPutList(homeworkList)
	var filteredChoreList []EntitiesFolder.ChoreOutput
	var filteredHomeWorkList []EntitiesFolder.HomeWorkOutput
	if err != nil {
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Errorf("A person with the id %s does not exist", params["id"]).Error()))
	} else {
		for _, chore := range choreOut {
			if strings.ToLower(chore.Status) == params["status"] {
				filteredChoreList = append(filteredChoreList, chore)
			}
		}
		for _, homeWork := range HomeOut {
			if strings.ToLower(homeWork.Status) == params["status"] {
				filteredHomeWorkList = append(filteredHomeWorkList, homeWork)
			}
		}
		if filteredChoreList == nil && filteredHomeWorkList == nil {
			var ans []interface{}
			json.NewEncoder(w).Encode(ans)
			return
		} else if filteredChoreList == nil {
			ans := []interface{}{filteredHomeWorkList}
			json.NewEncoder(w).Encode(ans)
			return
		} else if filteredHomeWorkList == nil {
			ans := []interface{}{filteredChoreList}
			json.NewEncoder(w).Encode(ans)
			return
		} else {
			ans := []interface{}{filteredChoreList, filteredHomeWorkList}
			json.NewEncoder(w).Encode(ans)
			return
		}
		w.WriteHeader(http.StatusOK)
	}
}

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
			json.NewEncoder(w).Encode(EntitiesFolder.ChoreToChoreOutPut(Chore))
		} else {
			json.NewEncoder(w).Encode(EntitiesFolder.HomeWorkToHomeWorkOutPut(HomeWork))
		}
	}
}

//func UpdateChore(choreToUpdate EntitiesFolder.Chore, TaskInput EntitiesFolder.TaskInput) error {
//	if TaskInput.Size != "" {
//		choreToUpdate.Size = EntitiesFolder.CreateSize(TaskInput.Size)
//	}
//}
//
//func UpdateHomework(homeworkToUpdate EntitiesFolder.HomeWork, TaskInput EntitiesFolder.TaskInput) error {
//
//}
func UpdateTask(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var TaskInput EntitiesFolder.TaskInput
	json.NewDecoder(r.Body).Decode(&TaskInput)
	emptyChore := EntitiesFolder.Chore{}
	emptyHomework := EntitiesFolder.HomeWork{}
	choreToUpdate, homeworkToUpdate, err1 := dbFolder.GetTask(params["id"])
	if err1 != nil {
		w.Header().Set("Content-Type", "text/plain")
		w.Write([]byte(fmt.Errorf("A task with the id %s does not exist", params["id"]).Error()))
		w.Write([]byte("Requested task is not present\n"))
		w.WriteHeader(http.StatusNotFound)
		return

	} //chore update
	if homeworkToUpdate == emptyHomework {
		err := dbFolder.UpdateTask(TaskInput, params["id"])
		if err != nil {
			w.Header().Set("Content-Type", "text/plain")
			w.Write([]byte(fmt.Errorf("Chore with task id %s failed to update", params["id"]).Error()))
			w.WriteHeader(http.StatusNotFound)
			return
		}
	} else if choreToUpdate == emptyChore {
		err := dbFolder.UpdateTask(TaskInput, params["id"])
		if err != nil {
			w.Header().Set("Content-Type", "text/plain")
			w.Write([]byte(fmt.Errorf("Homework with task id %s failed to update", params["id"]).Error()))
			w.WriteHeader(http.StatusNotFound)
			return
		}
	} else {
		w.Header().Set("Content-Type", "text/plain")
		w.Write([]byte(fmt.Errorf("A task with the id %s does not exist", params["id"]).Error()))
		w.WriteHeader(http.StatusNotFound)
		return
	}

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
