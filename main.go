package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"miniProject/EntitiesFolder/TaskFolder"
)

func main() {
	//t := TaskFolder.Task{OwnerId: "eden"}
	h := TaskFolder.HomeWork{TaskFolder.Task{"eden"}}
	fmt.Println(h.Task.GetOwnerId())

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
