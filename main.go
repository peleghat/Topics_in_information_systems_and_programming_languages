package main

import (
	_ "github.com/go-sql-driver/mysql"
	"miniProject/APIFolder"
	"miniProject/dbFolder"
)

func main() {
	dbFolder.CreateDb()
	//defer dbFolder.KillDb()
	APIFolder.InitServer()

	//p1 := EntitiesFolder.NewPerson("peleg", "peleg@gmail.com", "go")
	//p2 := EntitiesFolder.NewPerson("peleg", "gmail1", "go")
	//
	//dbFolder.InsertPerson(p1)
	//dbFolder.InsertPerson(p2)
	//
	//tHomework := EntitiesFolder.NewTask(p1.GetId(), 1, "Homework", "pliz finish")
	//h := EntitiesFolder.NewHomeWork("Hedva", EntitiesFolder.ClockUpdate("1996-10-10"), tHomework)
	//dbFolder.AddHomeWork(h)
	//
	//tChore := EntitiesFolder.NewTask(p1.GetId(), 0, "Chore", "pliz finish 2")
	//c := EntitiesFolder.NewChore(2, tChore)
	//dbFolder.AddChore(c)
}
