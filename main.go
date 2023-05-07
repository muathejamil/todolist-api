package main

import (
	"net/http"
	"todolist-api/handlers"
)

func init() {

}
func main() {
	mux := http.NewServeMux()
	mux.Handle("/todos/", &handlers.TodosHandler{})
	mux.Handle("/todos", &handlers.TodosHandler{})
	http.ListenAndServe("localhost:8000", mux)
}
