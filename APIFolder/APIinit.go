package APIFolder

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"log"
	"net/http"
)

func InitServer() {
	// Init router
	r := mux.NewRouter()

	// router handler - Endpoints
	// people endpoint
	r.HandleFunc("/api/people/", APIFunctionHandler).Methods("POST", "GET")
	r.HandleFunc("/api/people/{id}", APIFunctionHandler).Methods("GET", "PATCH", "DELETE")
	r.HandleFunc("/api/people/{id}/tasks/", APIFunctionHandler).Methods("GET", "POST")

	// tasks endpoint
	r.HandleFunc("/api/tasks/{id}", APIFunctionHandler).Methods("GET", "PATCH", "DELETE")
	r.HandleFunc("/api/tasks/{id}/status", APIFunctionHandler).Methods("GET", "PUT")
	r.HandleFunc("/api/tasks/{id}/owner", APIFunctionHandler).Methods("GET", "PUT")
	http.Handle("/", r)

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"POST", "OPTIONS", "GET", "PATCH", "DELETE", "PUT", "FETCH"},
		AllowedHeaders: []string{"*"},
	})
	log.Fatal(http.ListenAndServe(":9000", c.Handler(r)))

	fmt.Printf("Server start working 9000")

}
