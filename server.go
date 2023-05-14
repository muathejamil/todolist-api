package main

import (
	"net/http"
	"todolist-api/cache"
	"todolist-api/handlers"
	"todolist-api/repo"
	"todolist-api/services"
	"todolist-api/storage"
)

func InitServer() {
	todoRepo := repo.TodoRepo{
		DB: storage.DB,
	}
	service := services.Service{
		CreateTodo:  &todoRepo,
		DeleteTodo:  &todoRepo,
		GetTodo:     &todoRepo,
		UpdateTodo:  &todoRepo,
		GetAllTodos: &todoRepo,
	}

	cache := cache.NewRedisCache("localhost:6379", 0, 10)
	mux := http.NewServeMux()
	mux.Handle("/api/v1/todos/", &handlers.TodosHandler{
		CreateTodo:  &service,
		DeleteTodo:  &service,
		GetTodo:     &service,
		GetAllTodos: &service,
		UpdateTodo:  &service,
		TodoCache:   cache,
	})
	mux.Handle("/api/v1/todos", &handlers.TodosHandler{
		CreateTodo:  &service,
		DeleteTodo:  &service,
		GetTodo:     &service,
		GetAllTodos: &service,
		UpdateTodo:  &service,
		TodoCache:   cache,
	})
	http.ListenAndServe("localhost:8000", mux)
}
