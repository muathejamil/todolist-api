package repos

import (
	"todolist-api/models"
)

// CreateTodo creates a new Todo
// Params todoIDTO contracts.TodoIDTO
// returns contracts.TodoIDTO, error
func (t *TodoRepo) CreateTodo(todo models.Todo) (models.Todo, error) {
	// TODO: ID is our responsibility not DB.
	dbTodo := t.DB.Create(&todo)
	if dbTodo.Error != nil {
		return models.Todo{}, dbTodo.Error
	}
	return todo, nil
}
