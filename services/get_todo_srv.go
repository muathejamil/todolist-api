package services

import (
	"todolist-api/models"
	"todolist-api/services/cache/redis"
	"todolist-api/services/contracts"
)

// TodoGetter gets todo interface.
type TodoGetter interface {
	GetTodo(id uint) (models.Todo, error)
}

// Getter gets todo struct.
type Getter struct {
	Repo      TodoGetter
	TodoCache redis.CacheRedis
}

// Get gets todo by id
// Params id uint
// returns contracts.TodoIDTO, error
func (g *Getter) Get(id uint) (contracts.TodoIDTO, error) {
	// Get from cache
	todo, err := g.TodoCache.Get(uint64(id))
	if err != nil {
		// get from database
		todo, err = g.Repo.GetTodo(id)
		if err != nil {
			return contracts.TodoIDTO{}, err
		}
		// set value to database
		g.TodoCache.Set(uint64(todo.ID), todo)
	}
	return contracts.ToIDTO(todo), nil
}
