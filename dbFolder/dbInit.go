package dbFolder

import (
	"database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"log"
)

const DatabaseName = "minidb"

const CreatePersonsTable = "CREATE TABLE IF NOT EXISTS Persons(" +
	"id varchar(255) NOT NULL, " +
	"name varchar(255), " +
	"email varchar(255) UNIQUE, " +
	"favProg varchar(255), " +
	"ActiveTaskCount integer NOT NULL, " +
	"PRIMARY KEY (id));"

const CreateTasksTable = "CREATE TABLE IF NOT EXISTS Tasks(" +
	"id varchar(255) NOT NULL PRIMARY KEY, " +
	"ownerId varchar(255) NOT NULL, " +
	"status integer NOT NULL, " +
	"taskType varchar(255) NOT NULL, " +
	"description varchar(255) NOT NULL, " +
	"course_homework varchar(255), " +
	"dueDate_homework date, " +
	"size_chore integer, " +
	"FOREIGN KEY (ownerId) REFERENCES Persons(id));"

// CreateDb function crates a new database and set up the new tables
func CreateDb() {
	config := mysql.Config{
		User:   "root",
		Passwd: "edenandpelegdb", //  Pelegedendb
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
	fmt.Println("db was created1")
}

// KillDb function Used for debugging and testing purposes
// Drops the database after main
func KillDb() {
	err, db := connectToDb()
	_, err = db.Exec("DROP DATABASE minidb")
	if err != nil {
		panic(err)
	}
}

// the function connects to the database, if succeed returns the database itself, else returns err
func connectToDb() (error, *sql.DB) {
	config := mysql.Config{
		User:   "root",
		Passwd: "edenandpelegdb", // Pelegedendb
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: DatabaseName,
	}
	db, err := sql.Open("mysql", config.FormatDSN())
	if err != nil {
		log.Fatal("Failed To connect")
		return err, nil
	}
	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
		return pingErr, nil
	}
	return nil, db
}
