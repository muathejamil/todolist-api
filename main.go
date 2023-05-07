package main

import (
	"net/http"
	"todolist-api/handlers"
	"todolist-api/resources"
)

func init() {
	resources.LoadEnvVariables()
	resources.ConnectToDb()
}
func main() {
	mux := http.NewServeMux()
	mux.Handle("/todos/", &handlers.TodosHandler{})
	mux.Handle("/todos", &handlers.TodosHandler{})
	http.ListenAndServe("localhost:8000", mux)
}
