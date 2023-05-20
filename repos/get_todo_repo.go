package repos

import (
	"errors"
	"todolist-api/models"
)

// GetTodo gets todo by id
// Params id uint
// returns contracts.TodoIDTO, error
func (t *TodoRepo) GetTodo(id uint) (models.Todo, error) {
	var todo models.Todo
	t.DB.Find(&todo, id)
	if todo.ID == 0 {
		return models.Todo{}, errors.New("todo not found")
	}
	return todo, nil
}
