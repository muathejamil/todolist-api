package services

import (
	"todolist-api/cache/redis"
	"todolist-api/models"
	"todolist-api/services/contracts"
)

// TodoGetter gets todo interface.
type TodoGetter interface {
	GetTodo(id uint) (models.Todo, error)
}

// Getter gets todo struct.
type Getter struct {
	repo      TodoGetter
	TodoCache redis.CacheRedis
}

// Get gets todo by id
// Params id uint
// returns contracts.TodoIDTO, error
func (g *Getter) Get(id uint) (contracts.TodoIDTO, error) {
	todo, err := g.repo.GetTodo(id)
	if err != nil {
		return contracts.TodoIDTO{}, err
	}
	return contracts.ToIDTO(todo), nil
}
