package main

import (
	"net/http"
	"time"
	"todolist-api/cache/redis"
	"todolist-api/handlers"
	"todolist-api/services"
)

const ExpireAfter time.Duration = 10

// InitServer inits server.
func InitServer() {
	// get DB connection.
	//connection := storage.GetDBConnection()
	//todoRepo := repos.TodoRepo{
	//	DB: connection,
	//}

	// create new redis cache
	cache := redis.NewRedisCache("localhost:6379", 0, ExpireAfter)

	// define todo services.
	creatorSrv := services.Creator{}
	deleterSrv := services.Deleter{}
	oneGetterSrv := services.Getter{
		TodoCache: cache,
	}
	getterSrv := services.AllGetter{}
	updaterSrv := services.Updater{}

	mux := http.NewServeMux()
	// handle incoming requests.
	mux.Handle("/api/v1/todos/", &handlers.TodosHandler{
		CreateTodo:  &creatorSrv,
		DeleteTodo:  &deleterSrv,
		GetTodo:     &oneGetterSrv,
		GetAllTodos: &getterSrv,
		UpdateTodo:  &updaterSrv,
	})
	// handle incoming requests.
	mux.Handle("/api/v1/todos", &handlers.TodosHandler{
		CreateTodo:  &creatorSrv,
		DeleteTodo:  &deleterSrv,
		GetTodo:     &oneGetterSrv,
		GetAllTodos: &getterSrv,
		UpdateTodo:  &updaterSrv,
	})
	http.ListenAndServe("localhost:8000", mux)
}
