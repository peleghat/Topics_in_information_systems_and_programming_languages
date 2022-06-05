package dbFolder

import (
	"fmt"
	"miniProjectV2/EntitiesFolder"
)

func InsertPerson(p EntitiesFolder.Person) bool {
	err, db := connectToDb()
	if err != nil {
		panic(err)
		fmt.Println("Unable to connect to the DB")
		return false
	}
	defer db.Close()
	q := "INSERT INTO Persons VALUES ( ?, ?, ?, ?, ?)"
	insertResult, err := db.Query(q, p.GetId(), p.GetName(), p.GetEmail(), p.GetFavProg(), p.GetActiveTaskCount())
	if err != nil {
		panic(err.Error())
		fmt.Println("unable to insert the person to DB")
		return false
	}
	defer insertResult.Close()
	fmt.Println("success adding person to DB")
	return true
}

func DeletePerson(p EntitiesFolder.Person) bool {
	err, db := connectToDb()
	if err != nil {
		panic(err)
		fmt.Println("Unable to connect to the DB")
		return false
	}
	defer db.Close()
	q := "DELETE FROM Persons WHERE id =?"
	_, err = db.Query(q, p.GetId())
	if err != nil {
		panic(err.Error())
		fmt.Println("unable to delete the person from the DB")
		return false
	}
	fmt.Println("success deleting person from the DB")
	return true
}

func GetPerson(id string) EntitiesFolder.Person {
	err, db := connectToDb()
	if err != nil {
		panic(err)
		fmt.Println("Unable to connect to the DB")
		return EntitiesFolder.Person{}
	}
	defer db.Close()
	var personOutput EntitiesFolder.Person
	q := "SELECT * FROM Persons WHERE id =?"
	err = db.QueryRow(q, id).Scan(&personOutput.Id, &personOutput.Name, &personOutput.Email, &personOutput.FavProg, &personOutput.ActiveTaskCount)
	if err != nil {
		panic(err.Error())
		fmt.Println("unable to get the person from the DB")
		return EntitiesFolder.Person{}
	}
	fmt.Println("success getting person from the DB")
	return personOutput
}

func GetAllPersons() []EntitiesFolder.Person {
	err, db := connectToDb()
	if err != nil {
		panic(err)
		fmt.Println("Unable to connect to the DB")
		return []EntitiesFolder.Person{}
	}
	defer db.Close()
	personsRes, err := db.Query("SELECT * FROM Persons")
	if err != nil {
		panic(err.Error())
		fmt.Println("unable to get all persons from the DB")
		return []EntitiesFolder.Person{}
	}
	personListOutput := []EntitiesFolder.Person{}
	for personsRes.Next() {
		var personOutput EntitiesFolder.Person
		err = personsRes.Scan(&personOutput.Id, &personOutput.Name, &personOutput.Email, &personOutput.FavProg, &personOutput.ActiveTaskCount)
		if err != nil {
			panic(err.Error())
			fmt.Println("unable to get all persons from the DB")
			return []EntitiesFolder.Person{}
		}
		personListOutput = append(personListOutput, personOutput)
	}
	return personListOutput
}

func UpdatePerson(p EntitiesFolder.Person) bool {
	err, db := connectToDb()
	if err != nil {
		panic(err)
		fmt.Println("Unable to connect to the DB")
		return false
	}
	defer db.Close()
	q := "UPDATE Persons SET name = ?, email = ? , favProg = ? where id = ?"
	res, err := db.Query(q, p.Name, p.Email, p.FavProg, p.GetId())
	if err != nil {
		panic(err)
		fmt.Println("Unable to Update person")
		return false
	}
	defer res.Close()
	return true
}

func AddHomeWork(h EntitiesFolder.HomeWork) bool {
	err, db := connectToDb()
	if err != nil {
		panic(err)
		fmt.Println("Unable to connect to the DB")
		return false
	}
	defer db.Close()

	q := "INSERT INTO Tasks VALUES ( ?, ?, ?, ?, ?, ?, ?, ?)"
	insertResult, err := db.Query(q, h.Task.GetId(), h.Task.GetOwnerId(), h.Task.GetStatus(), h.Task.GetTaskType(),
		h.Task.GetDescription(), h.GetCourse(), h.GetDueDate(), -1)
	if err != nil {
		panic(err.Error())
		fmt.Println(" unable to insert the Homework to the DB")
		return false
	}
	defer insertResult.Close()
	fmt.Println("success adding Homework to DB")
	return true
}

func AddChore(c EntitiesFolder.Chore) bool {
	err, db := connectToDb()
	if err != nil {
		panic(err)
		fmt.Println("Unable to connect to the DB")
		return false
	}
	defer db.Close()

	q := "INSERT INTO Tasks VALUES ( ?, ?, ?, ?, ?, ?, ?, ?)"
	insertResult, err := db.Query(q, c.Task.GetId(), c.Task.GetOwnerId(), c.Task.GetStatus(), c.Task.GetTaskType(),
		c.Task.GetDescription(), nil, nil, c.GetSize())
	if err != nil {
		panic(err.Error())
		fmt.Println(" unable to insert the Homework to the DB")
		return false
	}
	defer insertResult.Close()
	fmt.Println("success adding Homework to DB")
	return true
}

func GetTaskFromDb(id string) EntitiesFolder.Task {
	err, db := connectToDb()
	if err != nil {
		panic(err)
		fmt.Println("Unable to connect to the DB")
		return EntitiesFolder.Task{}
	}
	defer db.Close()
	q := "SELECT id,ownerId,status,taskType,description FROM Tasks WHERE id =?"
	var taskOutput EntitiesFolder.Task
	err = db.QueryRow(q, id).Scan(&taskOutput.Id, &taskOutput.OwnerId, &taskOutput.Status, &taskOutput.TaskType, &taskOutput.Description)
	if err != nil {
		panic(err.Error())
		fmt.Println("unable to get the task from the DB")
		return EntitiesFolder.Task{}
	}
	return taskOutput
}

func GetChoreFromDb(id string) EntitiesFolder.Chore {
	err, db := connectToDb()
	if err != nil {
		panic(err)
		fmt.Println("Unable to connect to the DB")
		return EntitiesFolder.Chore{}
	}
	defer db.Close()
	q := "SELECT size_chore FROM Tasks WHERE id =?"
	var size int
	err = db.QueryRow(q, id).Scan(&size)
	if err != nil {
		panic(err.Error())
		fmt.Println("unable to get the task from the DB")
		return EntitiesFolder.Chore{}
	}
	return EntitiesFolder.NewChore(EntitiesFolder.Size(size), GetTaskFromDb(id))
}

func GetHomeWorkFromDb(id string) EntitiesFolder.HomeWork {
	err, db := connectToDb()
	if err != nil {
		panic(err)
		fmt.Println("Unable to connect to the DB")
		return EntitiesFolder.HomeWork{}
	}
	defer db.Close()
	q := "SELECT course_homework, dueDate_homework FROM Tasks WHERE id =?"
	var _Course string
	var _DueDate string
	err = db.QueryRow(q, id).Scan(&_Course, &_DueDate)
	if err != nil {
		panic(err.Error())
		fmt.Println("unable to get the task from the DB")
		return EntitiesFolder.HomeWork{}
	}
	return EntitiesFolder.NewHomeWork(_Course, EntitiesFolder.ClockUpdate(_DueDate), GetTaskFromDb(id))
}

func GetTask(id string) (EntitiesFolder.Chore, EntitiesFolder.HomeWork) {
	err, db := connectToDb()
	if err != nil {
		panic(err)
		fmt.Println("Unable to connect to the DB")
		return EntitiesFolder.Chore{}, EntitiesFolder.HomeWork{}
	}
	defer db.Close()
	//var taskOutput EntitiesFolder.Task
	task := GetTaskFromDb(id)
	if task.GetTaskType() == "Chore" {
		return GetChoreFromDb(id), EntitiesFolder.HomeWork{}
	} else if task.GetTaskType() == "Homework" {
		return EntitiesFolder.Chore{}, GetHomeWorkFromDb(id)
	} else {
		return EntitiesFolder.Chore{}, EntitiesFolder.HomeWork{}
	}
}
