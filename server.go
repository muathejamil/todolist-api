package main

import (
	"net/http"
	"time"
	"todolist-api/cache/redis"
	"todolist-api/handlers"
	"todolist-api/repos"
	"todolist-api/services"
	"todolist-api/storage"
)

const ExpireAfter time.Duration = 10

// InitServer inits server.
func InitServer() {
	todoRepo := repos.TodoRepo{
		DB: storage.DB,
	}
	service := services.Service{
		CreateTodo:  &todoRepo,
		DeleteTodo:  &todoRepo,
		GetTodo:     &todoRepo,
		UpdateTodo:  &todoRepo,
		GetAllTodos: &todoRepo,
	}

	// create new redis cache

	cache := redis.NewRedisCache("localhost:6379", 0, ExpireAfter)
	mux := http.NewServeMux()
	// handle incoming requests.
	mux.Handle("/api/v1/todos/", &handlers.TodosHandler{
		CreateTodo:  &service,
		DeleteTodo:  &service,
		GetTodo:     &service,
		GetAllTodos: &service,
		UpdateTodo:  &service,
		TodoCache:   cache,
	})
	// handle incoming requests.
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
