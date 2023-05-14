package redis

import "todolist-api/contracts"

type CacheRedis interface {
	Set(key uint64, value contracts.TodoDTO) error
	Get(key uint64) (contracts.TodoDTO, error)
}
