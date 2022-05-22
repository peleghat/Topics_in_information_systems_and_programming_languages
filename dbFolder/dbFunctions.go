package dbFolder

import (
	"fmt"
	"miniProject/EntitiesFolder"
)

func InsertPerson(p EntitiesFolder.Person) bool {
	err, db := connectToDb()
	if err != nil {
		panic(err)
		return false
	}
	defer db.Close()
	q := "INSERT INTO Persons VALUES ( ?, ?, ?, ?) "
	insertResult, err := db.Query(q, p.GetId(), p.GetName(), p.GetEmail(), p.GetFavProg())
	if err != nil {
		panic(err.Error())
		fmt.Println(" unable to insert the person to DB")
		return false
	}
	defer insertResult.Close()
	fmt.Println("success adding person to DB")
	return true
}

//
//func InsertHomeWork(t TaskFolder.Task) bool {
//	err, db := connectToDb()
//	if err != nil {
//		panic(err)
//		return false
//	}
//	defer db.Close()
//	q := "INSERT INTO Tasks VALUES ( ?, ?, ?, ?, ?, ?, ?, ?) "
//	insertResult, err := db.Query(q, t.GetOwnerId(), t.GetStatus(), t.)
//	if err != nil {
//		panic(err.Error())
//		fmt.Println(" unable to insert the person to DB")
//		return false
//	}
//	defer insertResult.Close()
//	fmt.Println("success adding person to DB")
//	return true
//}
