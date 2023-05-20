package handlers

import (
	"todolist-api/cache/redis"
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
//func (g *OneGetterSrv) Get(writer http.ResponseWriter, request *http.Request) {
//	id := utils.ExtractIdFromURL(writer, request)
//	// Check if the value exist in the cache.
//	todo, err := g.TodoCache.Get(id)
//	if todo.Id == 0 {
//		// Get the value from the service
//		todo, err = g.serv.Get(uint(id))
//		if err != nil {
//			http.Error(writer, err.Error(), http.StatusBadRequest)
//			return
//		}
//		//store the value in the cache.
//		g.TodoCache.Set(id, todo)
//	}
//	// create new todoDTO slice.
//	todos := make([]contracts.TodoDTO, 0)
//	// append value
//	dtos := append(todos, todo)
//	// write response as to writer.
//	utils.WriteJsonTodosResponse(writer, dtos)
//}
