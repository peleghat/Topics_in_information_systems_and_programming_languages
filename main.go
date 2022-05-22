package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"miniProject/EntitiesFolder"
)

func main() {

	////t := TaskFolder.Task{OwnerId: "eden"}
	//h := EntitiesFolder.NewChore("blah", 1, EntitiesFolder.NewTask("a", 1))
	////fmt.Println(h.Task.GetOwnerId())
	//c1 := EntitiesFolder.Marvel{"the", EntitiesFolder.Comic{Universe: "MCU"}}
	//fmt.Println("Universe is:", c1.Str)
	//c2 := EntitiesFolder.DC{true, EntitiesFolder.Comic{Universe: "DC"}}
	//fmt.Println("Universe is:", c2.Comic.Universe)
	c3 := EntitiesFolder.NewChore("5", 1, EntitiesFolder.NewTask("peleg", 1))
	fmt.Println("Universe is:", c3.Description, c3.Size, c3.Task.Id, c3.Task.OwnerId, c3.Task.GetStatus())
	//c4 := EntitiesFolder.NewHomeWork("5", "2022-03-24", "details", EntitiesFolder.NewTask("peleg", 1))
	//fmt.Println("Universe is:", c4.Course, c4.DueDate, c4.Details, c4.Task.Id, c4.Task.OwnerId, c4.Task.Status)

	//c2 := EntitiesFolder.DC{true, EntitiesFolder.Comic{Universe: "DC"}}
	//fmt.Println("Universe is:", c2.Comic.Universe)
	//h := TaskFolder.NewHomeWork("mivne", time.Parse("2006-01-02", "1996-10-14"), "blah")
	//dbFolder.CreateDb()
	//fmt.Println("db was created")
	//p := EntitiesFolder.NewPerson("peleg", "gmail", "go")
	//fmt.Println("p is", p)
	//if dbFolder.InsertPerson(p) {
	//	fmt.Println("succeed to insert: p is", p)
	//} else {
	//	fmt.Println("unable to insert: p is", p)
	//}

}
