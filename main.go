package main

import (
	"fmt"
	"miniProject/dbFolder"
	"net/http"
)

// TODO - Remove later
func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome home!")
}

func main() {
	//router := mux.NewRouter().StrictSlash(true)
	//router.HandleFunc("/", homeLink)
	//log.Fatal(http.ListenAndServe(":8080", router))
	dbFolder.Create_db()
	//config := mysql.Config{
	//	User:   os.Getenv("root"),
	//	Passwd: os.Getenv("edenandpelegdb"),
	//	Net:    "tcp",
	//	Addr:   "localhost:3306",
	//}
	//db, err := sql.Open("mysql", config.FormatDSN())
	//if err != nil {
	//	panic(err)
	//}
	//defer db.Close()
	//insert, err := db.Query("INSERT INTO persons NAME('eden')")
	//if err != nil {
	//	panic(err)
	//}
	//defer insert.Close()
	//
	//fmt.Println("connect to db!")

}
