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
	p1 := EntitiesFolder.NewPerson("peleg", "gmail1", "go")
	p2 := EntitiesFolder.NewPerson("peleg", "gmail2", "go")
	dbFolder.InsertPerson(p1)
	dbFolder.InsertPerson(p2)

	t_homework := EntitiesFolder.NewTask(p1.GetId(), 1, "Homework", "pliz finish")
	h := EntitiesFolder.NewHomeWork("Hedva", EntitiesFolder.ClockUpdate("1996-10-10"), t_homework)
	dbFolder.AddHomeWork(h)

	t_chore := EntitiesFolder.NewTask(p2.GetId(), 2, "Chore", "pliz finish 2")
	c := EntitiesFolder.NewChore(2, t_chore)
	dbFolder.AddChore(c)
	// getTask is not working, because we passed 5 arguments to scan instead of 8
	// hello world
	//chore, _ := dbFolder.GetTask(t_homework.GetId())
	//fmt.Println("the homework is ", chore)

	/*p1 := EntitiesFolder.NewPerson("peleg", "gmail1", "go")
	dbFolder.InsertPerson(p1)
	fmt.Println(dbFolder.GetAllPersons())
	dbFolder.DeletePerson(p1)*/
}
