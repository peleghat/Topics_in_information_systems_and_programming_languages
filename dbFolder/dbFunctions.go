package dbFolder

import (
	"miniProjectV2/EntitiesFolder"
)

// Persons Table Functions

// InsertPerson function inserts a new Person to the person table,
// returns a boolean which says if the insertion was a success or a failure
func InsertPerson(p EntitiesFolder.Person) bool {
	err, db := connectToDb()
	if err != nil {
		panic(err)
		return false
	}
	defer db.Close()
	q := "INSERT INTO Persons VALUES ( ?, ?, ?, ?, ?)"
	insertResult, err := db.Query(q, p.GetId(), p.GetName(), p.GetEmail(), p.GetFavProg(), p.GetActiveTaskCount())
	if err != nil {
		panic(err.Error())
		return false
	}
	defer insertResult.Close()
	return true
}

// DeletePerson function deletes a Person from the person table,
// returns a boolean which says if the deletion was a success or a failure
func DeletePerson(p EntitiesFolder.Person) bool {
	err, db := connectToDb()
	if err != nil {
		panic(err)
		return false
	}
	defer db.Close()
	q := "DELETE FROM Persons WHERE id =?"
	_, err = db.Query(q, p.GetId())
	if err != nil {
		panic(err.Error())
		return false
	}
	return true
}

// GetPerson function gets a Person from the person table by his id,
// returns the person's instance if it succeeds else returns an empty person instance
func GetPerson(id string) EntitiesFolder.Person {
	err, db := connectToDb()
	if err != nil {
		panic(err)
		return EntitiesFolder.Person{}
	}
	defer db.Close()
	var personOutput EntitiesFolder.Person
	q := "SELECT * FROM Persons WHERE id =?"
	err = db.QueryRow(q, id).Scan(&personOutput.Id, &personOutput.Name, &personOutput.Email, &personOutput.FavProg, &personOutput.ActiveTaskCount)
	if err != nil {
		panic(err)
		return EntitiesFolder.Person{}
	}
	return personOutput
}

// GetAllPersons function returns the list of all the persons in the person table
// if it succeeds else returns an empty person list
func GetAllPersons() []EntitiesFolder.Person {
	err, db := connectToDb()
	if err != nil {
		panic(err)
		return []EntitiesFolder.Person{}
	}
	defer db.Close()
	personsRes, err := db.Query("SELECT * FROM Persons")
	if err != nil {
		panic(err.Error())
		return []EntitiesFolder.Person{}
	}
	var personList []EntitiesFolder.Person
	for personsRes.Next() {
		var person EntitiesFolder.Person
		err = personsRes.Scan(&person.Id, &person.Name, &person.Email, &person.FavProg, &person.ActiveTaskCount)
		if err != nil {
			panic(err)
			return []EntitiesFolder.Person{}
		}
		personList = append(personList, person)
	}
	return personList
}

// UpdatePerson function update a Person's details (getting person by his id)
// returns a boolean which says if the update was a success or a failure
// TODO - Maybe Change to seperate functions in the API Level
func UpdatePerson(p EntitiesFolder.Person) bool {
	err, db := connectToDb()
	if err != nil {
		panic(err)
		return false
	}
	defer db.Close()
	q := "UPDATE Persons SET name = ?, email = ? , favProg = ? where id = ?"
	res, err := db.Query(q, p.GetName(), p.GetEmail(), p.GetFavProg(), p.GetId())
	if err != nil {
		panic(err)
		return false
	}
	defer res.Close()
	return true
}

// Tasks Table Functions

// AddHomeWork function inserts a new HomeWork to the tasks table,
// returns a boolean which says if the insertion was a success or a failure
func AddHomeWork(h EntitiesFolder.HomeWork) bool {
	err, db := connectToDb()
	if err != nil {
		panic(err)
		return false
	}
	defer db.Close()
	q := "INSERT INTO Tasks VALUES ( ?, ?, ?, ?, ?, ?, ?, ?)"
	insertResult, err := db.Query(q, h.GetTask().GetId(), h.GetTask().GetOwnerId(), h.GetTask().GetStatus(),
		h.GetTask().GetTaskType(), h.GetTask().GetDescription(), h.GetCourse(), h.GetDueDate(), -1) // Size is -1
	if err != nil {
		panic(err)
		return false
	}
	defer insertResult.Close()
	return true
}

// AddChore function inserts a new HomeWork to the tasks table,
// returns a boolean which says if the insertion was a success or a failure
func AddChore(c EntitiesFolder.Chore) bool {
	err, db := connectToDb()
	if err != nil {
		panic(err)
		return false
	}
	defer db.Close()
	q := "INSERT INTO Tasks VALUES ( ?, ?, ?, ?, ?, ?, ?, ?)"
	insertResult, err := db.Query(q, c.GetTask().GetId(), c.GetTask().GetOwnerId(), c.GetTask().GetStatus(),
		c.GetTask().GetTaskType(), c.GetTask().GetDescription(), nil, nil, c.GetSize()) // CourseName and DueDate is nil
	if err != nil {
		panic(err)
		return false
	}
	defer insertResult.Close()
	return true
}

// GetTaskFromDb is a helper function which gets a task id, and returns the corresponding task instance
// if it succeeds, else returns an empty task instance
func GetTaskFromDb(id string) EntitiesFolder.Task {
	err, db := connectToDb()
	if err != nil {
		panic(err)
		return EntitiesFolder.Task{}
	}
	defer db.Close()
	q := "SELECT id,ownerId,status,taskType,description FROM Tasks WHERE id =?"
	var taskOutput EntitiesFolder.Task
	err = db.QueryRow(q, id).Scan(&taskOutput.Id, &taskOutput.OwnerId, &taskOutput.Status, &taskOutput.TaskType, &taskOutput.Description)
	if err != nil {
		panic(err)
		return EntitiesFolder.Task{}
	}
	return taskOutput
}

// GetChoreFromDb is a helper function which gets a task id, and returns the corresponding Chore instance
// if it succeeds, else returns an empty Chore instance
func GetChoreFromDb(id string) EntitiesFolder.Chore {
	err, db := connectToDb()
	if err != nil {
		panic(err)
		return EntitiesFolder.Chore{}
	}
	defer db.Close()
	q := "SELECT size_chore FROM Tasks WHERE id =?"
	var size int
	err = db.QueryRow(q, id).Scan(&size)
	if err != nil {
		panic(err)
		return EntitiesFolder.Chore{}
	}
	return EntitiesFolder.NewChore(EntitiesFolder.Size(size), GetTaskFromDb(id))
}

// GetHomeWorkFromDb is a helper function which gets a task id, and returns the corresponding HomeWork instance
// if it succeeds, else returns an empty HomeWork instance
func GetHomeWorkFromDb(id string) EntitiesFolder.HomeWork {
	err, db := connectToDb()
	if err != nil {
		panic(err)
		return EntitiesFolder.HomeWork{}
	}
	defer db.Close()
	q := "SELECT course_homework, dueDate_homework FROM Tasks WHERE id =?"
	var _Course string
	var _DueDate string
	err = db.QueryRow(q, id).Scan(&_Course, &_DueDate)
	if err != nil {
		panic(err)
		return EntitiesFolder.HomeWork{}
	}
	return EntitiesFolder.NewHomeWork(_Course, EntitiesFolder.ClockUpdate(_DueDate), GetTaskFromDb(id))
}

// GetTask function gets a task id and returns the corresponding Chore/HomeWork instance
// if it succeeds, else returns an empty tuple of Chore and HomeWork instances
func GetTask(id string) (EntitiesFolder.Chore, EntitiesFolder.HomeWork) {
	err, db := connectToDb()
	if err != nil {
		panic(err)
		return EntitiesFolder.Chore{}, EntitiesFolder.HomeWork{}
	}
	defer db.Close()
	task := GetTaskFromDb(id)
	if task.GetTaskType() == "Chore" {
		return GetChoreFromDb(id), EntitiesFolder.HomeWork{}
	} else if task.GetTaskType() == "Homework" {
		return EntitiesFolder.Chore{}, GetHomeWorkFromDb(id)
	} else {
		return EntitiesFolder.Chore{}, EntitiesFolder.HomeWork{}
	}
}
