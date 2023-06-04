package main

import (
	"fmt"
	"net/http"
	"time"
	"todolist-api/config"
	"todolist-api/handlers"
	"todolist-api/repos"
	"todolist-api/services"
	"todolist-api/services/cache/redis"
	"todolist-api/storage"
)

const ExpireAfter time.Duration = 10

// InitServer inits server.
func InitServer(configuration config.Configuration) {

	// get DB connection.
	connection := storage.GetDBConnection(configuration.DataBase)
	repo := repos.TodoRepo{
		DB: connection,
	}

	// create new redis cache
	redisHost := fmt.Sprintf("%s:%d", configuration.Redis.Host, configuration.Redis.Port)
	cache := redis.NewRedisCache(redisHost, configuration.Redis.DBNumber, ExpireAfter)

	// define todo services.
	creatorSrv := services.Creator{
		Repo: &repo,
	}
	deleterSrv := services.Deleter{
		Repo: &repo,
	}
	oneGetterSrv := services.Getter{
		Repo:      &repo,
		TodoCache: cache,
	}
	getterSrv := services.AllGetter{
		Repo: &repo,
	}
	updaterSrv := services.Updater{
		Repo: &repo,
	}

	// new serve mux.
	mux := http.NewServeMux()

	// handle incoming requests.
	mux.Handle("/api/v1/todos", &handlers.TodosHandler{
		CreateTodoSrv:  &creatorSrv,
		DeleteTodoSrv:  &deleterSrv,
		GetTodoSrv:     &oneGetterSrv,
		GetAllTodosSrv: &getterSrv,
		UpdateTodoSrv:  &updaterSrv,
	})

	err := http.ListenAndServe(fmt.Sprintf("%s:%d", configuration.Server.Host, configuration.Server.Port), mux)
	if err != nil {
		return
	}
}
