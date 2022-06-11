package main

import (
	_ "github.com/go-sql-driver/mysql"
	"miniProject/APIFolder"
	"miniProject/dbFolder"
)

func main() {
	dbFolder.CreateDb()
	APIFolder.InitServer()

	//a := []int{1, 2, 3}
	//b := []string{"a", "b"}
	//ans := []interface{}{a, b}
	//fmt.Println(ans)
	//defer dbFolder.KillDb()

	//p1 := EntitiesFolder.NewPersonInput("peleg", "peleg@gmail.com", "go")

	//p2 := EntitiesFolder.NewPerson("peleg", "gmail1", "go")
	//fmt.Println(dbFolder.InsertPerson(p1))
	//fmt.Println(dbFolder.InsertPerson(p2))
	//
	//tHomework := EntitiesFolder.NewTask(p1.GetId(), 1, "Homework", "pliz finish")
	//h := EntitiesFolder.NewHomeWork("Hedva", EntitiesFolder.ClockUpdate("1996-10-10"), tHomework)
	//dbFolder.AddHomeWork(h)
	//
	//tChore := EntitiesFolder.NewTask(p2.GetId(), 2, "Chore", "pliz finish 2")
	//c := EntitiesFolder.NewChore(2, tChore)
	//dbFolder.AddChore(c)
}
