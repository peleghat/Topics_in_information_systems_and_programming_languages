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

func GetTask(id string) (EntitiesFolder.Chore, EntitiesFolder.HomeWork) {
	err, db := connectToDb()
	if err != nil {
		panic(err)
		fmt.Println("Unable to connect to the DB")
		return EntitiesFolder.Chore{}, EntitiesFolder.HomeWork{}
	}
	defer db.Close()
	var taskOutput EntitiesFolder.Task
	q := "SELECT * FROM Tasks WHERE id =?"
	err = db.QueryRow(q, id).Scan(&taskOutput.Id, &taskOutput.OwnerId, &taskOutput.Status, &taskOutput.TaskType, &taskOutput.Description)
	if err != nil {
		panic(err.Error())
		fmt.Println("unable to get the person from the DB")
		return EntitiesFolder.Chore{}, EntitiesFolder.HomeWork{}
	}
	if taskOutput.GetTaskType() == "Chore" {
		var choreOutput EntitiesFolder.Chore
		err = db.QueryRow(q, id).Scan(&choreOutput.Size)
		fmt.Println("success getting chore from the DB")
		return EntitiesFolder.NewChore(choreOutput.GetSize(), taskOutput), EntitiesFolder.HomeWork{}
	} else {
		var homeworkOutput EntitiesFolder.HomeWork
		err = db.QueryRow(q, id).Scan(&homeworkOutput.Course, &homeworkOutput.DueDate)
		fmt.Println("success getting homework from the DB")
		return EntitiesFolder.Chore{}, EntitiesFolder.NewHomeWork(homeworkOutput.GetCourse(), homeworkOutput.GetDueDate(), taskOutput)
	}
}
