package handlers

import (
	"todolist-api/services/cache/redis"
	srvContracts "todolist-api/services/contracts"
)

// TodoGetterSrv todo getter service interface.
type TodoGetterSrv interface {
	Get(id uint) (srvContracts.TodoIDTO, error)
}

// OneGetterSrv todo getter service struct.
type OneGetterSrv struct {
	serv      TodoGetterSrv
	TodoCache redis.CacheRedis
}
