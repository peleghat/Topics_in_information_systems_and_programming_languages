package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"miniProject/EntitiesFolder"
	"miniProject/dbFolder"
)

func main() {
	dbFolder.CreateDb()
	defer dbFolder.KillDb()
	fmt.Println("Peleg123")

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
}
