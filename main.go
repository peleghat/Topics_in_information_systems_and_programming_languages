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
}
