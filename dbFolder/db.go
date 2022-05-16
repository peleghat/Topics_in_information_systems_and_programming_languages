package dbFolder

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

const DATABASE_NAME string = "minidb"

const CREATE_PERSONS_TABLE = "CREATE TABLE IF NOT EXISTS Persons(" +
	"id varchar(255) NOT NULL PRIMARY KEY, " +
	"name varchar(255), " +
	"email varchar(255) UNIQUE," +
	"favProg varchar(255));"

const CREATE_TASKS_TABLE = "CREATE TABLE IF NOT EXISTS Tasks(" +
	"id varchar(255) NOT NULL PRIMARY KEY, " +
	"ownerId varchar(255) NOT NULL FOREIGN KEY REFERENCES Persons(id), " +
	"status integer NOT NULL, " +
	"type string NOT NULL," +
	"course_homework string," +
	"dueDate_homework date" +
	"details_homework string," +
	"description_chore string," +
	"size_chore integer);"

func CreateDb() {
	//config := mysql.Config{
	//	User:   "root",
	//	Passwd: "edenandpelegdb",
	//	Net:    "tcp",
	//	Addr:   "127.0.0.1:3306",
	//}
	//db, err := sql.Open("mysql", config.FormatDSN())
	db, err := sql.Open("mysql", "root:edenandpelegdb@tcp(127.0.0.1:3306)/minidb")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	//_, err = db.Exec("CREATE DATABASE IF NOT EXISTS" + DATABASE_NAME)
	//if err != nil {
	//	panic(err)
	//}
	//_, err = db.Exec("USE " + DATABASE_NAME)
	//if err != nil {
	//	panic(err)
	//}
	//_, err = db.Exec(CREATE_PERSONS_TABLE)
	//if err != nil {
	//	panic(err)
	//}
	//
	////_, err = db.Exec(CREATE_TASKS_TABLE)
	////if err != nil {
	////	panic(err)
	////}
	//
	//insert, err := db.Query("INSERT INTO persons NAME('eden')")
	//if err != nil {
	//	panic(err)
	//}

	//defer insert.Close()
}
