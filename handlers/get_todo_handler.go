package handlers

import (
	"errors"
	"todolist-api/cache/redis"
	"todolist-api/handlers/contracts"
	srvContracts "todolist-api/services/contracts"
)

// TodoGetterSrv todo getter service interface.
type TodoGetterSrv interface {
	Get(id uint) (srvContracts.TodoIDTO, error)
}

// OneGetterSrv todo getter service interface.
type OneGetterSrv struct {
	serv      TodoGetterSrv
	TodoCache redis.CacheRedis
}

// Get get todo handler.
// Params writer http.ResponseWriter, request *http.Request.
// Returns todo if exist.
func (g *OneGetterSrv) Get(id uint64) (contracts.TodoDTO, error) {
	// Check if the value exist in the cache.
	todo, err := g.TodoCache.Get(id)
	if todo.Id == 0 {
		// Get the value from the service
		var todoIdto srvContracts.TodoIDTO
		todoIdto, err = g.serv.Get(uint(id))
		if err != nil {
			return contracts.TodoDTO{}, errors.New("unable to find todo")
		}
		todo = contracts.ToTodoDTO(todoIdto)
		//store the value in the cache.
		g.TodoCache.Set(id, todo)
	}
	return todo, nil
}
