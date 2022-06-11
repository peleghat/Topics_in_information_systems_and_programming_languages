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
	r.Methods("OPTIONS").HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
		})

	// router handler - Endpoints
	// people endpoint
	r.HandleFunc("/api/people/", APIFunctionHandler).Methods("POST", "GET")
	r.HandleFunc("/api/people/{id}", APIFunctionHandler).Methods("GET", "PATCH", "DELETE")
	r.Path("/api/people/{id}/tasks/").Queries("status", "{status}").HandlerFunc(APIFunctionHandler).Methods("GET")
	r.Path("/api/people/{id}/tasks/").HandlerFunc(APIFunctionHandler).Methods("GET", "POST")

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

	//Different format for the optional query

	log.Fatal(http.ListenAndServe(":9000", c.Handler(r)))

	fmt.Printf("Server start working 9000")

}
