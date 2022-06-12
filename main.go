package main

import (
	_ "github.com/go-sql-driver/mysql"
	"miniProject/APIFolder"
	"miniProject/dbFolder"
)

// the main function - create the db, and init the server
func main() {
	dbFolder.CreateDb()
	APIFolder.InitServer()
}
