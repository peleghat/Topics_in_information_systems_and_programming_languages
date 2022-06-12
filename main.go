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
	//p := EntitiesFolder.Person{}
	//if p == EntitiesFolder.Person{} {
	//	fmt.Println()
	//}
	//fmt.Println(p == EntitiesFolder.Person{})
	//p1 := EntitiesFolder.NewPerson("inctest", "inctest@gmail.com", "go")
	////p2 := EntitiesFolder.NewPerson("peleg", "gmail1", "go")
	////
	//dbFolder.InsertPerson(p1)
	//fmt.Println(p1.GetId())
	//fmt.Println(p1.GetId() == "cb96da67-02cf-4e1e-9932-fee6672cf709")
	//err, p := dbFolder.GetPerson("125b1e78-29b1-4229-8332-8de2b818e08c")
	//err := dbFolder.IncTaskToPerson("125b1e78-29b1-4229-8332-8de2b818e08c")
	//if err != nil {
	//	fmt.Println("failed to inc ")
	//}
	//fmt.Println("get id is : ", p.GetId())
	//fmt.Println(" id is : ", p.ID)
	//fmt.Println(" email is : ", p.Email)
	//fmt.Println(" name is : ", p.Name)

	//fmt.Println("get id  p1 is : ", p1.GetId())

	//dbFolder.InsertPerson(p2)
	//
	//tHomework := EntitiesFolder.NewTask(p1.GetId(), 1, "Homework", "pliz finish")
	//h := EntitiesFolder.NewHomeWork("Hedva", EntitiesFolder.ClockUpdate("1996-10-10"), tHomework)
	//err := dbFolder.AddHomeWork(h)
	//if err != nil {
	//	fmt.Println("failed to inc ")
	//}

	//tChore := EntitiesFolder.NewTask(p1.GetId(), 0, "Chore", "pliz finish 2")
	//c := EntitiesFolder.NewChore(2, tChore)
	//err := dbFolder.AddChore(c)
	//if err != nil {
	//	fmt.Println("failed to inc ")
	//}
}
