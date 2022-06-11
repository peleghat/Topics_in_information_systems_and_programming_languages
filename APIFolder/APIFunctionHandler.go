package APIFolder

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func APIFunctionHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	path := r.RequestURI
	method := r.Method
	params := mux.Vars(r)
	switch path {
	case "/api/people/":
		{
			if method == "POST" {
				AddPerson(w, r)
			} else if method == "GET" {
				GetAllPersons(w, r)
			} else {
				w.WriteHeader(http.StatusNotFound)
			}
		}
	case fmt.Sprintf("/api/people/%s", params["id"]):
		{
			if method == "GET" {
				GetPerson(w, r)
			} else if method == "PATCH" {
				UpdatePerson(w, r)
			} else if method == "DELETE" {
				DeletePerson(w, r)
			} else {
				w.WriteHeader(http.StatusNotFound)
			}
		}
	case fmt.Sprintf("/api/people/%s/tasks/", params["id"]):
		{
			if method == "GET" {
				GetAllPersonTasks(w, r)
			} else if method == "POST" {
				AddTaskToPerson(w, r)
			} else {
				w.WriteHeader(http.StatusNotFound)
			}
		}
	case fmt.Sprintf("/api/people/%s/tasks/?status=%s", params["id"], params["status"]):
		{
			if method == "GET" {
				GetPersonsTasksByStatus(w, r)
			} else {
				w.WriteHeader(http.StatusNotFound)
			}
		}
	case fmt.Sprintf("/api/tasks/%s", params["id"]):
		{
			if method == "GET" {
				GetTask(w, r)
			} else if method == "PATCH" {
				UpdateTask(w, r)
			} else if method == "DELETE" {
				DeleteTask(w, r)
			} else {
				w.WriteHeader(http.StatusNotFound)
			}
		}
	case fmt.Sprintf("/api/tasks/%s/status", params["id"]):
		{
			if method == "GET" {
				GetTaskStatus(w, r)
			} else if method == "PUT" {
				SetTaskStatus(w, r)
			} else {
				w.WriteHeader(http.StatusNotFound)
			}
		}
	case fmt.Sprintf("/api/tasks/%s/owner", params["id"]):
		{
			if method == "GET" {
				GetTaskOwner(w, r)
			} else if method == "PUT" {
				SetTaskOwner(w, r)
			} else {
				w.WriteHeader(http.StatusNotFound)
			}
		}
	default:
		{
			w.WriteHeader(http.StatusNotFound)
		}
	}
}
