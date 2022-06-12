package APIFolder

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"miniProject/EntitiesFolder"
	"miniProject/ErrorsFolder"
	"miniProject/dbFolder"
	"net/http"
	"strings"
)

// AddPerson function adds a person to the database.
// Checks if the person email is unique. if the insertion was a success return 201, else return error 400
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
		w.Write([]byte("person created successfully"))
		w.WriteHeader(http.StatusCreated)
		return
	}
}

// GetAllPersons functions return all persons in the db. if success return 200 else returns error 400
func GetAllPersons(w http.ResponseWriter, r *http.Request) {
	err, persons := dbFolder.GetAllPersons()
	if err != nil {
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Errorf("query to the db has failed").Error()))
	} else {
		json.NewEncoder(w).Encode(persons)
		w.WriteHeader(http.StatusOK)
	}
}

// GetPerson function returns the corresponding person by the given id from the api request,
// if success return 200 and the person details as json, else returns error 404
func GetPerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	err, person := dbFolder.GetPerson(params["id"])
	if err != nil {
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Requested person is not present.\n"))
		w.Write([]byte(fmt.Errorf("A person with the id %s does not exist", params["id"]).Error()))
	} else {
		w.Write([]byte("Person data provided."))
		json.NewEncoder(w).Encode(person)
		w.WriteHeader(http.StatusOK)
	}
}

// UpdatePerson function updates a specific person, with id given, with optional parameters (name, email and favoriteProgrammingLanguage).
// the function decode the params from the json body request.
// if success return 200 else returns error 404
func UpdatePerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var PersonInput EntitiesFolder.PersonInput
	json.NewDecoder(r.Body).Decode(&PersonInput)
	err1, PersonToUpdate := dbFolder.GetPerson(params["id"])
	if err1 != nil {
		w.Header().Set("Content-Type", "text/plain")
		w.Write([]byte("Requested person is not present.\n"))
		w.Write([]byte(fmt.Errorf("A person with the id %s does not exist", params["id"]).Error()))
		w.WriteHeader(http.StatusNotFound)
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
		w.Write([]byte("Requested person is not present.\n"))
		w.WriteHeader(http.StatusNotFound)
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

// DeletePerson function deletes a person from the database according to its id given.
// if success return 200 else returns error 404
func DeletePerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	choreList, HomeWorkList, _ := dbFolder.GetTasksFromPerson(params["id"])
	for _, chore := range choreList {
		dbFolder.DeleteTask(chore.GetTask().GetId())
	}
	for _, homeWork := range HomeWorkList {
		dbFolder.DeleteTask(homeWork.GetTask().GetId())
	}
	err2 := dbFolder.DeletePerson(params["id"])
	if err2 != nil {
		w.Header().Set("Content-Type", "text/plain")
		w.Write([]byte("Requested person is not present.\n"))
		w.WriteHeader(http.StatusNotFound)
		switch err2 {
		case ErrorsFolder.ErrDbConnection:
			{
				w.Write([]byte(fmt.Errorf("failed to connect to db").Error()))
			}
		case ErrorsFolder.ErrDbQuery:
			{
				w.Write([]byte(fmt.Errorf("A person with id %s does not exists", params["id"]).Error()))
			}
		}
	} else {
		w.Write([]byte("Person removed successfully.\n"))
		w.WriteHeader(http.StatusOK)
	}
}

// GetAllPersonTasks return all the tasks of a given person by its id.
// if success return 200 and the tasks, else returns error 404
func GetAllPersonTasks(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	choreList, homeworkList, err := dbFolder.GetTasksFromPerson(params["id"])
	if err != nil {
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Requested person is not present.\n"))
		w.Write([]byte(fmt.Errorf("A person with the id %s does not exist", params["id"]).Error()))
	} else {
		if choreList == nil && homeworkList == nil {
			var ans []interface{}
			w.Write([]byte("Task found and provided.\n"))
			json.NewEncoder(w).Encode(ans)
			w.WriteHeader(http.StatusOK)
			return
		} else if choreList == nil {
			ans2 := []interface{}{EntitiesFolder.HomeWorkListToHomeWorkOutPutList(homeworkList)}
			w.Write([]byte("Task found and provided.\n"))
			json.NewEncoder(w).Encode(ans2)
			w.WriteHeader(http.StatusOK)
			return
		} else if homeworkList == nil {
			ans := []interface{}{EntitiesFolder.ChoreListToChoreOutPutList(choreList)}
			w.Write([]byte("Task found and provided.\n"))
			json.NewEncoder(w).Encode(ans)
			w.WriteHeader(http.StatusOK)
			return
		} else {
			ans := []interface{}{EntitiesFolder.ChoreListToChoreOutPutList(choreList),
				EntitiesFolder.HomeWorkListToHomeWorkOutPutList(homeworkList)}
			w.Write([]byte("Task found and provided.\n"))
			json.NewEncoder(w).Encode(ans)
			w.WriteHeader(http.StatusOK)
			return
		}
	}
}

// AddTaskToPerson functions adds a task to specific person by its id given.
// the task itself is in the api request body represent by json.
// if the status field is not specified in the request body, the server will default to marking the newly created task as active.
// the function increases the active task count of the owner of this task by 1
// if success return 201 and the tasks, else returns error 400/404
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
					w.Write([]byte("Requested person is not present.\n"))
					w.Write([]byte(fmt.Errorf("A person with the id %s does not exist", params["id"]).Error()))
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
					w.Write([]byte("Requested person is not present.\n"))
					w.Write([]byte(fmt.Errorf("A person with the id %s does not exist", params["id"]).Error()))
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
		w.Write([]byte(fmt.Errorf("task type: %s does not exist", TaskToAdd.TaskType).Error()))
		w.WriteHeader(http.StatusNotFound) //404
		return
	}
}

// GetPersonsTasksByStatus function return all the tasks of a given person by its id.
// Returns an array of tasks that the person with id owns.
// The optional parameter status allows the caller to filter by task status.
// When status is not present, return an array of all tasks, regardless os their status.
// For example, when status=done, the server will return only the tasks' person id whose status is done.
// if success return 200 and the filtered tasks, else returns error 404
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
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Requested person is not present.\n"))
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
		w.Write([]byte("Task found and provided.\n"))
		w.WriteHeader(http.StatusOK)
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
	}
}

// GetTask function return a task from its id provided.
// if success return 200 and the task as json, else returns error 404
func GetTask(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	Chore, HomeWork, err := dbFolder.GetTask(params["id"])
	if err != nil {
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("requested task is not presented.\n"))
		switch err {
		case ErrorsFolder.ErrDbConnection:
			{
				w.Write([]byte(fmt.Errorf("failed to connect to db").Error()))
			}
		case ErrorsFolder.ErrIllegalValues:
			{
				w.Write([]byte(fmt.Errorf("A task with the id %s does not exist", params["id"]).Error()))
			}
		case ErrorsFolder.ErrNotExist:
			{
				w.Write([]byte(fmt.Errorf("A task with the id %s does not exist", params["id"]).Error()))
			}
		}
	} else {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Task found and provided.\n"))
		emptyChore := EntitiesFolder.Chore{}
		if Chore != emptyChore {
			json.NewEncoder(w).Encode(EntitiesFolder.ChoreToChoreOutPut(Chore))
		} else {
			json.NewEncoder(w).Encode(EntitiesFolder.HomeWorkToHomeWorkOutPut(HomeWork))
		}
	}
}

// UpdateChore is a helper function to UpdateTask.
// the functions check which fields need to be updated
func UpdateChore(choreToUpdate EntitiesFolder.ChoreOutput, TaskInput EntitiesFolder.TaskInput) EntitiesFolder.ChoreOutput {
	if TaskInput.Size != "" {
		choreToUpdate.Size = TaskInput.Size
	}
	if TaskInput.Description != "" {
		choreToUpdate.Description = TaskInput.Description
	}
	if TaskInput.Status != "" {
		choreToUpdate.Status = TaskInput.Status
	}
	return choreToUpdate
}

// UpdateHomework is a helper function to UpdateTask.
// the functions check which fields need to be updated
func UpdateHomework(homeworkToUpdate EntitiesFolder.HomeWorkOutput, TaskInput EntitiesFolder.TaskInput) EntitiesFolder.HomeWorkOutput {
	if TaskInput.Status != "" {
		homeworkToUpdate.Status = TaskInput.Status
	}
	if TaskInput.Description != "" {
		homeworkToUpdate.Description = TaskInput.Description
	}
	if TaskInput.Course != "" {
		homeworkToUpdate.Course = TaskInput.Course
	}
	if TaskInput.DueDate != "" {
		homeworkToUpdate.DueDate = TaskInput.DueDate
	}
	return homeworkToUpdate
}

// UpdateTask function updates a specific task, given the task id as a parameter.
// the function check if the task is a chore or homework, and updates it accordingly (using the helper functions).
// Data fields that should be updated. ALL FIELDS ARE OPTIONAL - NO FIELD ARE REQUIRED IN THIS CONTEXT.
// the fields that need to be updated are given as a json in the api body request.
// if success return 200, else returns error 404
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
		choreToDb1 := EntitiesFolder.ChoreToChoreOutPut(choreToUpdate)
		choreToDb2 := UpdateChore(choreToDb1, TaskInput)
		err := dbFolder.UpdateTask(choreToDb2, EntitiesFolder.HomeWorkOutput{})
		if err != nil {
			w.Header().Set("Content-Type", "text/plain")
			w.Write([]byte(fmt.Errorf("chore with task id %s failed to update", params["id"]).Error()))
			w.WriteHeader(http.StatusNotFound)
			return
		} else {
			w.Write([]byte("Task updated successfully.Data contains updated task.\n"))
			json.NewEncoder(w).Encode(choreToDb2)
			w.WriteHeader(http.StatusOK)
			return
		}
	} else if choreToUpdate == emptyChore {
		homeworkToDb1 := EntitiesFolder.HomeWorkToHomeWorkOutPut(homeworkToUpdate)
		homeworkToDb2 := UpdateHomework(homeworkToDb1, TaskInput)
		err := dbFolder.UpdateTask(EntitiesFolder.ChoreOutput{}, homeworkToDb2)
		if err != nil {
			w.Header().Set("Content-Type", "text/plain")
			w.Write([]byte(fmt.Errorf("homework with task id %s failed to update", params["id"]).Error()))
			w.WriteHeader(http.StatusNotFound)
			return
		} else {
			w.Write([]byte("Task updated successfully.Data contains updated task.\n"))
			json.NewEncoder(w).Encode(homeworkToDb2)
			w.WriteHeader(http.StatusOK)
			return
		}
	} else {
		w.Header().Set("Content-Type", "text/plain")
		w.Write([]byte(fmt.Errorf("A task with the id %s does not exist", params["id"]).Error()))
		w.WriteHeader(http.StatusNotFound)
		return
	}
}

// DeleteTask function deletes a task given by its id as a parameter.
// the function decreases the active task count of the owner of this task by 1
// if success return 200, else returns error 404
func DeleteTask(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	err := dbFolder.DeleteTask(params["id"])
	if err != nil {
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Requested task is not present\n"))
		switch err {
		case ErrorsFolder.ErrDbConnection:
			{
				w.Write([]byte(fmt.Errorf("failed to connect to db").Error()))
			}
		case ErrorsFolder.ErrNotExist:
			{
				w.Write([]byte(fmt.Errorf("A task with the id %s does not exist", params["id"]).Error()))
			}
		case ErrorsFolder.ErrDbQuery:
			{
				w.Write([]byte(fmt.Errorf("failed to delete the task with id %s to db", params["id"]).Error()))
			}
		default:
			{
				w.Write([]byte(fmt.Errorf("unknown error has occured").Error()))
			}
		}
	} else {
		w.Write([]byte("Task removed successfully.\n"))
		w.WriteHeader(http.StatusOK)
	}
}

// GetTaskStatus function return the status of a given task by its id
// if success return 200, else returns error 404
func GetTaskStatus(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	err, t := dbFolder.GetTaskFromDb(params["id"])
	if err != nil {
		w.Header().Set("Content-Type", "text/plain")
		w.Write([]byte("Requested task is not present."))
		w.WriteHeader(http.StatusNotFound)
		switch err {
		case ErrorsFolder.ErrDbConnection:
			{
				w.Write([]byte(fmt.Errorf("failed to connect to db").Error()))
			}
		case ErrorsFolder.ErrNotExist:
			{
				w.Write([]byte(fmt.Errorf("A task with the id %s does not exist", params["id"]).Error()))
			}
		}
	} else {
		w.Write([]byte("Task's current status is provided."))
		w.Write([]byte(EntitiesFolder.StatusToString(t.GetStatus())))
		w.WriteHeader(http.StatusOK)
	}
}

// SetTaskStatus sets the status of a given task by its id
// the function checks if the status is valid.
// if success return 204, else returns error 400/404
func SetTaskStatus(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	bodyBytes, _ := ioutil.ReadAll(r.Body)
	output := string(bodyBytes)
	output = output[1 : len(output)-1]
	output = strings.ToLower(output)
	if output != "active" && output != "done" {
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Request contains data that make no sense."))
		w.Write([]byte(fmt.Errorf("value %s is not a legal task status", output).Error()))
		return
	}
	err := dbFolder.UpdateTaskStatus(params["id"], EntitiesFolder.CreateStatus(output))
	if err != nil {
		w.Header().Set("Content-Type", "text/plain")
		switch err {
		case ErrorsFolder.ErrDbConnection:
			{
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte(fmt.Errorf("failed to connect to db").Error()))
			}
		case ErrorsFolder.ErrNotExist:
			{
				w.WriteHeader(http.StatusNotFound)
				w.Write([]byte("Requested task is not present."))
				w.Write([]byte(fmt.Errorf("A task with the id %s does not exist", params["id"]).Error()))
			}
		default:
			{
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte(fmt.Errorf("unknown error has occured").Error()))
			}
		}
	} else {
		w.Write([]byte("task's status updated successfully."))
		w.WriteHeader(http.StatusNoContent)
	}
}

// GetTaskOwner function gets the task owner by a given task by its id.
// if success return 200, else returns error 404
func GetTaskOwner(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	err, t := dbFolder.GetTaskFromDb(params["id"])
	if err != nil {
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Requested task is not present."))
		switch err {
		case ErrorsFolder.ErrDbConnection:
			{
				w.Write([]byte(fmt.Errorf("failed to connect to db").Error()))
			}
		case ErrorsFolder.ErrNotExist:
			{
				w.Write([]byte(fmt.Errorf("A task with the id %s does not exist", params["id"]).Error()))
			}
		}
	} else {
		w.Write([]byte("Id of the owner of the task."))
		w.Write([]byte(t.GetOwnerId()))
		w.WriteHeader(http.StatusOK)
	}
}

// SetTaskOwner function, sets the owner of a given task to a new owner.
// the function decreases the active task count of the first person who holds the task by 1, and increases
// by 1 to the new person who holds the task
// if success return 204, else returns error 404
func SetTaskOwner(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var newOwnerId string
	json.NewDecoder(r.Body).Decode(&newOwnerId)
	err := dbFolder.SetTaskOwner(params["id"], newOwnerId)
	if err != nil {
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusNotFound)
		switch err {
		case ErrorsFolder.ErrDbConnection:
			{
				w.Write([]byte(fmt.Errorf("failed to connect to db").Error()))
			}
		case ErrorsFolder.ErrDbQuery:
			{
				w.Write([]byte(fmt.Errorf("requested task is not present\n").Error()))
				w.Write([]byte(fmt.Errorf("A task with the id %s does not exist", params["id"]).Error()))
			}
		case ErrorsFolder.ErrNotExist:
			{
				w.Write([]byte(fmt.Errorf("requested task is not present\n").Error()))
				w.Write([]byte(fmt.Errorf("A task with the id %s does not exist", params["id"]).Error()))
			}
		default:
			{
				w.Write([]byte(fmt.Errorf("unknown error has occured").Error()))
			}
		}
	} else {
		w.Write([]byte("task owner updated successfully."))
		w.WriteHeader(http.StatusNoContent)
	}
}
