package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

// Golang program to illustrate the
// concept of inheritance

// declaring a struct
type Comic struct {

	// declaring struct variable
	Universe string
}

// function to return the
// universe of the comic
func (comic Comic) ComicUniverse() string {

	// returns comic universe
	return comic.Universe
}

// declaring a struct
type Marvel struct {

	// anonymous field,
	// this is composition where
	// the struct is embedded
	Comic
}

// declaring a struct
type DC struct {

	// anonymous field
	Comic
}

// main function
func main() {

	// creating an instance
	c1 := Marvel{

		// child struct can directly
		// access base struct variables
		Comic{
			Universe: "MCU",
		},
	}

	// child struct can directly
	// access base struct methods

	// printing base method using child
	fmt.Println("Universe is:", c1.ComicUniverse())

	c2 := DC{
		Comic{
			Universe: "DC",
		},
	}

	// printing base method using child
	fmt.Println("Universe is:", c2.ComicUniverse())
}

//func main() {
//
//	//h := TaskFolder.NewHomeWork("mivne", time.Parse("2006-01-02", "1996-10-14"), "blah")
//	//dbFolder.CreateDb()
//	//fmt.Println("db was created")
//	//p := EntitiesFolder.NewPerson("peleg", "gmail", "go")
//	//fmt.Println("p is", p)
//	//if dbFolder.InsertPerson(p) {
//	//	fmt.Println("succeed to insert: p is", p)
//	//} else {
//	//	fmt.Println("unable to insert: p is", p)
//	//}
//
//}
