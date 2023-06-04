package redis

import (
	"todolist-api/models"
)

type CacheRedis interface {
	Set(key uint64, value models.Todo) error
	Get(key uint64) (models.Todo, error)
}
