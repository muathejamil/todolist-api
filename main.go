package main

import (
	"net/http"
	"todolist-api/handlers"
	"todolist-api/repo"
	"todolist-api/resources"
	"todolist-api/services"
)

func init() {
	//resources.LoadEnvVariables()
	resources.ConnectToDb()
}
func main() {
	//TODO: We can use DI here
	todoRepo := repo.TodoRepo{DB: resources.DB}
	todoListService := services.TodoListService{Repo: todoRepo}
	mux := http.NewServeMux()
	mux.Handle("/api/v1/todos/", &handlers.TodosHandler{TodolistService: todoListService})
	mux.Handle("/api/v1/todos", &handlers.TodosHandler{TodolistService: todoListService})
	http.ListenAndServe("localhost:8000", mux)
}

//How to change line 20, 21 https://github.com/golang/go/issues/29261
