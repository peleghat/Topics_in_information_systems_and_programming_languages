//package main
//
//import (
//	"fmt"
//	"miniProject/dbFolder"
//	"net/http"
//)
//
//// TODO - Remove later
//func homeLink(w http.ResponseWriter, r *http.Request) {
//	fmt.Fprintf(w, "Welcome home!")
//}
//
//func main() {
//	//router := mux.NewRouter().StrictSlash(true)
//	//router.HandleFunc("/", homeLink)
//	//log.Fatal(http.ListenAndServe(":8080", router))
//	dbFolder.CreateDb()
//	//config := mysql.Config{
//	//	User:   os.Getenv("root"),
//	//	Passwd: os.Getenv("edenandpelegdb"),
//	//	Net:    "tcp",
//	//	Addr:   "localhost:3306",
//	//}
//	//db, err := sql.Open("mysql", config.FormatDSN())
//	//if err != nil {
//	//	panic(err)
//	//}
//	//defer db.Close()
//	//insert, err := db.Query("INSERT INTO persons NAME('eden')")
//	//if err != nil {
//	//	panic(err)
//	//}
//	//defer insert.Close()
//	//
//	//fmt.Println("connect to db!")
//
//}

//package main
//
//import (
//	"database/sql"
//	"fmt"
//	_ "github.com/go-sql-driver/mysql"
//)

package main

import (
	"database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"miniProject/EntitiesFolder"
)

const DatabaseName = "minidb"

const CreatePersonsTable = "CREATE TABLE IF NOT EXISTS Persons(" +
	"id varchar(255) NOT NULL, " +
	"name varchar(255), " +
	"email varchar(255) UNIQUE," +
	"favProg varchar(255)," +
	"PRIMARY KEY (id));"

const CreateTasksTable = "CREATE TABLE IF NOT EXISTS Tasks(" +
	"id varchar(255) NOT NULL PRIMARY KEY, " +
	"ownerId varchar(255) NOT NULL, " +
	"status integer NOT NULL, " +
	"taskType varchar(255) NOT NULL, " +
	"course_homework varchar(255), " +
	"dueDate_homework date, " +
	"details_homework varchar(255), " +
	"description_chore varchar(255), " +
	"size_chore integer, " +
	"FOREIGN KEY (ownerId) REFERENCES Persons(id));"

func main() {

	config := mysql.Config{
		User:   "root",
		Passwd: "edenandpelegdb",
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
	}
	db, err := sql.Open("mysql", config.FormatDSN())
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	_, err = db.Exec("CREATE DATABASE IF NOT EXISTS " + DatabaseName)
	if err != nil {
		panic(err)
	}
	_, err = db.Exec("USE " + DatabaseName)
	if err != nil {
		panic(err)
	}
	_, err = db.Exec(CreatePersonsTable)
	if err != nil {
		panic(err)
	}
	_, err = db.Exec(CreateTasksTable)
	if err != nil {
		panic(err)
	}

	p := EntitiesFolder.NewPerson("peleg", "gmail", "go")
	fmt.Println("p is", p)

	//db := openConnection()
	q := "INSERT INTO Persons VALUES ( ?, ? ,?, ?) "
	insertResult, err := db.Query(q, p.GetId(), p.GetName(), p.GetEmail(), p.GetFavProg())
	if err != nil {
		panic(err.Error())
		fmt.Println(" err- no success adding person to table")
		defer db.Close()
	}
	defer insertResult.Close()
	fmt.Println("success adding person to table")
	defer db.Close()

}

func openConnection() *sql.DB {
	config := mysql.Config{
		User:   "root",
		Passwd: "edenandpelegdb",
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
	}
	db, err := sql.Open("mysql", config.FormatDSN())
	if err != nil {
		panic(err.Error())
		return nil
	}
	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
		return nil
	}
	return db
}

//func AddPerson(p EntitiesFolder.Person) bool {
//
//	//db := openConnection()
//	if db != nil {
//		q := "INSERT INTO Persons VALUES ( ?, ? ,?, ?) "
//		insertResult, err := db.Query(q, p.GetId(), p.GetName(), p.GetEmail(), p.GetFavProg())
//		if err != nil {
//			panic(err.Error())
//			defer db.Close()
//			return false
//		}
//		defer insertResult.Close()
//		//defer db.Close()
//		return true
//	}
//	defer db.Close()
//	return false
//}
