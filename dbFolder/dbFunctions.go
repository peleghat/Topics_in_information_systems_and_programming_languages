package dbFolder

import (
	"fmt"
	"miniProject/EntitiesFolder"
	"miniProject/ErrorsFolder"
)

// Persons Table Functions

// InsertPerson function inserts a new Person to the person table,
// returns nil if the insertion was a success else an error
func InsertPerson(p EntitiesFolder.Person) error {
	err, db := connectToDb()
	if err != nil {
		return ErrorsFolder.ErrDbConnection
	}
	defer db.Close()
	q := "INSERT INTO Persons VALUES ( ?, ?, ?, ?, ?)"
	insertResult, err := db.Query(q, p.GetId(), p.GetName(), p.GetEmail(), p.GetFavProg(), p.GetActiveTaskCount())
	if err != nil {
		return ErrorsFolder.ErrAlreadyExist
	}
	defer insertResult.Close()
	return nil
}

// DeletePerson function deletes a Person from the person table,
// returns nil if the deletion was a success else an error
func DeletePerson(s string) error {
	err, db := connectToDb()
	if err != nil {
		return ErrorsFolder.ErrDbConnection
	}
	defer db.Close()
	q := "DELETE FROM Persons WHERE id =?"
	_, err = db.Query(q, s)
	if err != nil {
		return ErrorsFolder.ErrDbQuery
	}
	return nil
}

// GetPerson function gets a Person from the person table by his id,
// returns the person's instance if it succeeds else returns an empty person instance and an error
func GetPerson(id string) (error, EntitiesFolder.Person) {
	err, db := connectToDb()
	if err != nil {
		return ErrorsFolder.ErrDbConnection, EntitiesFolder.Person{}
	}
	defer db.Close()
	var personOutput EntitiesFolder.Person
	q := "SELECT * FROM Persons WHERE id =?"
	err = db.QueryRow(q, id).Scan(&personOutput.ID, &personOutput.Name, &personOutput.Email, &personOutput.FavProg, &personOutput.ActiveTaskCount)
	if err != nil {
		return ErrorsFolder.ErrNotExist, EntitiesFolder.Person{}
	}
	return nil, personOutput
}

// GetAllPersons function returns the list of all the persons in the person table
// if it succeeds else returns an empty person list and error
func GetAllPersons() (error, []EntitiesFolder.Person) {
	err, db := connectToDb()
	if err != nil {
		return ErrorsFolder.ErrDbConnection, []EntitiesFolder.Person{}
	}
	defer db.Close()
	personId, err := db.Query("SELECT id FROM Persons")
	if err != nil {
		return ErrorsFolder.ErrDbQuery, []EntitiesFolder.Person{}
	}
	var personList []EntitiesFolder.Person
	for personId.Next() {
		var _id string
		err = personId.Scan(&_id)
		if err != nil {
			return ErrorsFolder.ErrIllegalValues, []EntitiesFolder.Person{}
		}
		err, personId := GetPerson(_id)
		if err != nil {
			return ErrorsFolder.ErrNotExist, []EntitiesFolder.Person{}

		} else {
			personList = append(personList, personId)
		}
	}
	return nil, personList
}

// UpdatePerson function update a Person's details (getting person by his id)
// returns nil if the update was a success else an error
func UpdatePerson(p EntitiesFolder.Person) error {
	err, db := connectToDb()
	if err != nil {
		return ErrorsFolder.ErrDbConnection
	}
	defer db.Close()
	q := "UPDATE Persons SET name = ?, email = ? , favProg = ? where id = ?"
	res, err := db.Query(q, p.GetName(), p.GetEmail(), p.GetFavProg(), p.GetId())
	if err != nil {
		return ErrorsFolder.ErrDbQuery
	}
	defer res.Close()
	return nil
}

// Tasks Table Functions

// AddHomeWork function inserts a new HomeWork to the tasks table,
// returns an error which says if the insertion was a success or a failure
func AddHomeWork(h EntitiesFolder.HomeWork) error {
	err1, db := connectToDb()
	if err1 != nil {
		return ErrorsFolder.ErrDbConnection
	}
	defer db.Close()
	q := "INSERT INTO Tasks VALUES ( ?, ?, ?, ?, ?, ?, ?, ?)"
	insertResult, err2 := db.Query(q, h.GetTask().GetId(), h.GetTask().GetOwnerId(), h.GetTask().GetStatus(),
		h.GetTask().GetTaskType(), h.GetTask().GetDescription(), h.GetCourse(), h.GetDueDate(), -1) // Size is -1
	if err2 != nil {
		return ErrorsFolder.ErrDbQuery
	}
	err3 := IncTaskToPerson(h.GetTask().GetOwnerId())
	if err3 != nil {
		return ErrorsFolder.ErrDbQuery
	}
	defer insertResult.Close()
	return nil
}

// IncTaskToPerson function increment active task count field on a given person
// returns nil if the insertion was a success else an error
func IncTaskToPerson(id string) error {
	err, db := connectToDb()
	if err != nil {
		return ErrorsFolder.ErrDbConnection
	}
	defer db.Close()
	err, personToInc := GetPerson(id)
	if err != nil {
		return ErrorsFolder.ErrNotExist
	}

	update := personToInc.IncActiveTaskCount()
	q := "UPDATE Persons SET ActiveTaskCount = ? where id = ?"
	res, err := db.Query(q, update, id)
	if err != nil {
		return ErrorsFolder.ErrDbQuery
	}
	defer res.Close()
	return nil
}

// DecTaskToPerson function decrement active task count field on a given person
// returns nil if the insertion was a success else an error
func DecTaskToPerson(id string) error {
	err, db := connectToDb()
	if err != nil {
		return ErrorsFolder.ErrDbConnection
	}
	defer db.Close()
	err, personToDec := GetPerson(id)
	if err != nil {
		return ErrorsFolder.ErrNotExist
	}

	update := personToDec.DecActiveTaskCount()
	q := "UPDATE Persons SET ActiveTaskCount = ? where id = ?"
	res, err := db.Query(q, update, id)
	if err != nil {
		return ErrorsFolder.ErrDbQuery
	}
	defer res.Close()
	return nil
}

// AddChore function inserts a new HomeWork to the tasks table,
// returns an error which says if the insertion was a success or a failure
func AddChore(c EntitiesFolder.Chore) error {
	err1, db := connectToDb()
	if err1 != nil {
		return ErrorsFolder.ErrDbConnection
	}
	defer db.Close()
	q := "INSERT INTO Tasks VALUES ( ?, ?, ?, ?, ?, ?, ?, ?)"
	insertResult, err2 := db.Query(q, c.GetTask().GetId(), c.GetTask().GetOwnerId(), c.GetTask().GetStatus(),
		c.GetTask().GetTaskType(), c.GetTask().GetDescription(), nil, nil, c.GetSize()) // CourseName and DueDate is nil
	if err2 != nil {
		return ErrorsFolder.ErrDbQuery
	}
	err3 := IncTaskToPerson(c.GetTask().GetOwnerId())
	if err3 != nil {
		return ErrorsFolder.ErrDbQuery
	}
	defer insertResult.Close()
	return nil
}

// GetTaskFromDb is a helper function which gets a task id, and returns the corresponding task instance
// if it succeeds, else returns an empty task instance and an error
func GetTaskFromDb(id string) (error, EntitiesFolder.Task) {
	err, db := connectToDb()
	if err != nil {
		return ErrorsFolder.ErrDbConnection, EntitiesFolder.Task{}
	}
	defer db.Close()
	q := "SELECT id,ownerId,status,taskType,description FROM Tasks WHERE id =?"
	var taskOutput EntitiesFolder.Task
	err = db.QueryRow(q, id).Scan(&taskOutput.Id, &taskOutput.OwnerId, &taskOutput.Status, &taskOutput.TaskType, &taskOutput.Description)
	if err != nil {
		return ErrorsFolder.ErrNotExist, EntitiesFolder.Task{}
	}
	return nil, taskOutput
}

// GetChoreFromDb is a helper function which gets a task id, and returns the corresponding Chore instance
// if it succeeds, else returns an empty Chore instance and an error
func GetChoreFromDb(id string) (error, EntitiesFolder.Chore) {
	err, db := connectToDb()
	if err != nil {
		return ErrorsFolder.ErrDbConnection, EntitiesFolder.Chore{}
	}
	defer db.Close()
	q := "SELECT size_chore FROM Tasks WHERE id =?"
	var size int
	err = db.QueryRow(q, id).Scan(&size)
	if err != nil {
		return ErrorsFolder.ErrNotExist, EntitiesFolder.Chore{}
	}
	err2, task := GetTaskFromDb(id)
	if err2 != nil {
		return ErrorsFolder.ErrNotExist, EntitiesFolder.Chore{}
	}
	return nil, EntitiesFolder.NewChore(EntitiesFolder.Size(size), task)
}

// GetHomeWorkFromDb is a helper function which gets a task id, and returns the corresponding HomeWork instance
// if it succeeds, else returns an empty HomeWork instance and an error
func GetHomeWorkFromDb(id string) (error, EntitiesFolder.HomeWork) {
	err, db := connectToDb()
	if err != nil {
		return ErrorsFolder.ErrDbConnection, EntitiesFolder.HomeWork{}
	}
	defer db.Close()
	q := "SELECT course_homework, dueDate_homework FROM Tasks WHERE id =?"
	var _Course string
	var _DueDate string
	err = db.QueryRow(q, id).Scan(&_Course, &_DueDate)
	if err != nil {
		return ErrorsFolder.ErrNotExist, EntitiesFolder.HomeWork{}
	}
	err2, task := GetTaskFromDb(id)
	if err2 != nil {
		return ErrorsFolder.ErrNotExist, EntitiesFolder.HomeWork{}
	}
	return nil, EntitiesFolder.NewHomeWork(_Course, EntitiesFolder.ClockUpdate(_DueDate), task)
}

// GetTask function gets a task id and returns the corresponding Chore/HomeWork instance
// if it succeeds, else returns an empty tuple of Chore and HomeWork instances and an error
func GetTask(id string) (EntitiesFolder.Chore, EntitiesFolder.HomeWork, error) {
	err, db := connectToDb()
	if err != nil {
		return EntitiesFolder.Chore{}, EntitiesFolder.HomeWork{}, ErrorsFolder.ErrDbConnection
	}
	defer db.Close()
	err2, task := GetTaskFromDb(id)
	if err2 != nil {
		return EntitiesFolder.Chore{}, EntitiesFolder.HomeWork{}, ErrorsFolder.ErrNotExist
	}
	if task.GetTaskType() == "Chore" {
		err1, chore := GetChoreFromDb(id)
		if err1 != nil {
			return EntitiesFolder.Chore{}, EntitiesFolder.HomeWork{}, ErrorsFolder.ErrNotExist
		}
		return chore, EntitiesFolder.HomeWork{}, nil
	} else if task.GetTaskType() == "Homework" {
		err3, homework := GetHomeWorkFromDb(id)
		if err3 != nil {
			return EntitiesFolder.Chore{}, EntitiesFolder.HomeWork{}, ErrorsFolder.ErrNotExist
		}
		return EntitiesFolder.Chore{}, homework, nil
	} else {
		return EntitiesFolder.Chore{}, EntitiesFolder.HomeWork{}, ErrorsFolder.ErrIllegalValues
	}
}

// GetAllTTasks function returns the list of all the Tasks in the Task table
// if it succeeds else returns an empty HomeWork list,empty Chore List and an error
func GetAllTTasks() ([]EntitiesFolder.Chore, []EntitiesFolder.HomeWork, error) {
	err, db := connectToDb()
	if err != nil {
		return []EntitiesFolder.Chore{}, []EntitiesFolder.HomeWork{}, ErrorsFolder.ErrDbConnection
	}
	defer db.Close()
	TaskIds, err := db.Query("SELECT id FROM Tasks")
	if err != nil {
		return []EntitiesFolder.Chore{}, []EntitiesFolder.HomeWork{}, ErrorsFolder.ErrDbQuery
	}
	var ChoreList []EntitiesFolder.Chore
	var HomeWorkList []EntitiesFolder.HomeWork
	emptyChore := EntitiesFolder.Chore{}
	emptyHomeWork := EntitiesFolder.HomeWork{}
	for TaskIds.Next() {
		var _id string
		err = TaskIds.Scan(&_id)
		if err != nil {
			return []EntitiesFolder.Chore{}, []EntitiesFolder.HomeWork{}, ErrorsFolder.ErrNotExist
		}
		Chore, HomeWork, err2 := GetTask(_id)
		if err2 != nil {
			return []EntitiesFolder.Chore{}, []EntitiesFolder.HomeWork{}, ErrorsFolder.ErrNotExist
		}
		if Chore != emptyChore {
			ChoreList = append(ChoreList, Chore)
		}
		if HomeWork != emptyHomeWork {
			HomeWorkList = append(HomeWorkList, HomeWork)
		}
	}
	return ChoreList, HomeWorkList, nil
}

// DeleteTask function deletes a Task from the Tasks table,
// returns nil if the deletion was a success else an error
func DeleteTask(taskId string) error {
	err, db := connectToDb()
	if err != nil {
		return ErrorsFolder.ErrDbConnection
	}
	err, task := GetTaskFromDb(taskId)
	if err != nil {
		return ErrorsFolder.ErrNotExist
	}
	err3 := DecTaskToPerson(task.GetOwnerId())
	if err3 != nil {
		return ErrorsFolder.ErrNotExist
	}
	defer db.Close()
	q := "DELETE FROM Tasks WHERE id =?"
	_, err = db.Query(q, taskId)
	if err != nil {
		return ErrorsFolder.ErrNotExist
	}
	return nil
}

// UpdateTask function update a Task's details (getting the task by its id)
// returns nil if the update was a success else an error
func UpdateTask(c EntitiesFolder.ChoreOutput, h EntitiesFolder.HomeWorkOutput) error {
	err, db := connectToDb()
	if err != nil {
		return ErrorsFolder.ErrDbConnection
	}
	defer db.Close()
	emptyChore := EntitiesFolder.ChoreOutput{}
	emptyHomework := EntitiesFolder.HomeWorkOutput{}
	if c == emptyChore {
		//homework update
		status := EntitiesFolder.StatusStrToInt(h.Status)
		dueDate := EntitiesFolder.ClockUpdate(h.DueDate)
		q := "UPDATE Tasks SET status = ?, description = ?, course_homework=?, dueDate_homework=? where id = ?"
		res, err := db.Query(q, status, h.Description, h.Course, dueDate, h.Id)
		if err != nil {
			return ErrorsFolder.ErrDbQuery
		}
		defer res.Close()
		return nil

	} else if h == emptyHomework {
		// chore update
		status := EntitiesFolder.StatusStrToInt(c.Status)
		//desc := c.Description
		size := EntitiesFolder.SizeStrToInt(c.Size)
		cid := c.Id
		q := "UPDATE Tasks SET status = ?, description = ?, size_chore= ? where id = ?"
		res, err := db.Query(q, status, c.Description, size, cid)
		if err != nil {
			fmt.Println("chore failed")
			return ErrorsFolder.ErrDbQuery
		}
		defer res.Close()
		return nil
	} else {
		return ErrorsFolder.ErrNotExist

	}
}

// GetTasksFromPerson function returns the list of tasks of a specific person
// if succeeds else, returns an empty Chore list, empty HomeWork List and an error
func GetTasksFromPerson(personId string) ([]EntitiesFolder.Chore, []EntitiesFolder.HomeWork, error) {
	err, db := connectToDb()
	if err != nil {
		return []EntitiesFolder.Chore{}, []EntitiesFolder.HomeWork{}, ErrorsFolder.ErrDbConnection
	}
	defer db.Close()
	q := "SELECT id FROM Tasks WHERE ownerId = ?"
	TaskIds, err := db.Query(q, personId)
	if err != nil {
		return []EntitiesFolder.Chore{}, []EntitiesFolder.HomeWork{}, ErrorsFolder.ErrDbQuery
	}
	var ChoreList []EntitiesFolder.Chore
	var HomeWorkList []EntitiesFolder.HomeWork
	emptyChore := EntitiesFolder.Chore{}
	emptyHomeWork := EntitiesFolder.HomeWork{}
	for TaskIds.Next() {
		var _id string
		err = TaskIds.Scan(&_id)
		if err != nil {
			return []EntitiesFolder.Chore{}, []EntitiesFolder.HomeWork{}, ErrorsFolder.ErrNotExist
		}
		Chore, HomeWork, err2 := GetTask(_id)
		if err2 != nil {
			return []EntitiesFolder.Chore{}, []EntitiesFolder.HomeWork{}, ErrorsFolder.ErrNotExist
		}
		if Chore != emptyChore {
			ChoreList = append(ChoreList, Chore)
		}
		if HomeWork != emptyHomeWork {
			HomeWorkList = append(HomeWorkList, HomeWork)
		}
	}
	return ChoreList, HomeWorkList, nil
}

// GetPersonFromTask function returns the corresponding person to a specific task
// if succeeds, else returns an empty person instance and an error
func GetPersonFromTask(t EntitiesFolder.Task) (error, EntitiesFolder.Person) {
	err, db := connectToDb()
	if err != nil {
		return ErrorsFolder.ErrDbConnection, EntitiesFolder.Person{}
	}
	defer db.Close()
	err, personId := GetPerson(t.GetOwnerId())
	if err != nil {
		return ErrorsFolder.ErrNotExist, EntitiesFolder.Person{}

	} else {
		return nil, personId
	}
}

// UpdateTaskStatus function updates the status of the task
// returns nil if the update was a success else an error
func UpdateTaskStatus(taskId string, newStatus EntitiesFolder.Status) error {
	err, db := connectToDb()
	if err != nil {
		return ErrorsFolder.ErrDbConnection
	}
	defer db.Close()
	err3, _ := GetTaskFromDb(taskId)
	if err3 != nil {
		return ErrorsFolder.ErrNotExist
	}
	q := "UPDATE Tasks SET status = ? where id = ?"
	res, err2 := db.Query(q, newStatus, taskId)

	if err2 != nil {
		return ErrorsFolder.ErrNotExist
	}
	defer res.Close()
	return nil
}

// SetTaskOwner function updates the owner of the task
// returns nil if the update was a success else an error
func SetTaskOwner(taskId string, newOwnerId string) error {
	err1, db := connectToDb()
	if err1 != nil {
		return ErrorsFolder.ErrDbConnection
	}
	defer db.Close()
	err2, _ := GetPerson(newOwnerId)
	if err2 != nil { // owner not exists
		return ErrorsFolder.ErrNotExist
	}
	err3, task := GetTaskFromDb(taskId)
	if err3 != nil { // task not exists
		return ErrorsFolder.ErrNotExist
	}
	personToDec := task.OwnerId
	q := "UPDATE Tasks SET ownerId = ? where id = ?"
	res, err4 := db.Query(q, newOwnerId, taskId)
	if err4 != nil {
		return ErrorsFolder.ErrDbQuery
	}
	err5 := IncTaskToPerson(newOwnerId)
	if err5 != nil {
		return ErrorsFolder.ErrDbQuery
	}
	err6 := DecTaskToPerson(personToDec)
	if err6 != nil {
		return ErrorsFolder.ErrDbQuery
	}
	defer res.Close()
	return nil
}
