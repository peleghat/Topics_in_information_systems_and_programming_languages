package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"miniProjectV2/EntitiesFolder"
	"miniProjectV2/dbFolder"
)

func main() {
	dbFolder.CreateDb()
	fmt.Println("db was created")
	defer dbFolder.KillDb()
	p1 := EntitiesFolder.NewPerson("peleg", "gmail1", "go")
	p2 := EntitiesFolder.NewPerson("peleg", "gmail2", "go")
	dbFolder.InsertPerson(p1)
	dbFolder.InsertPerson(p2)

	tHomework := EntitiesFolder.NewTask(p1.GetId(), 1, "Homework", "pliz finish")
	h := EntitiesFolder.NewHomeWork("Hedva", EntitiesFolder.ClockUpdate("1996-10-10"), tHomework)
	dbFolder.AddHomeWork(h)

	tChore := EntitiesFolder.NewTask(p2.GetId(), 2, "Chore", "pliz finish 2")
	c := EntitiesFolder.NewChore(2, tChore)
	dbFolder.AddChore(c)
	chore, homework := dbFolder.GetTask(tHomework.GetId())
	fmt.Println("chore is", chore == EntitiesFolder.Chore{})
	fmt.Println("homerwork is", homework == EntitiesFolder.HomeWork{})

	/*p1 := EntitiesFolder.NewPerson("peleg", "gmail1", "go")
	dbFolder.InsertPerson(p1)
	fmt.Println(dbFolder.GetAllPersons())
	dbFolder.DeletePerson(p1)*/
}
