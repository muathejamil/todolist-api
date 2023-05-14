package repos

import (
	"errors"
	"todolist-api/models"
	"todolist-api/services/contracts"
)

// GetTodo gets todo by id
// Params id uint
// returns contracts.TodoIDTO, error
func (t *TodoRepo) GetTodo(id uint) (contracts.TodoIDTO, error) {
	var todo models.Todo
	t.DB.Find(&todo, id)
	if todo.ID == 0 {
		return contracts.TodoIDTO{}, errors.New("todo not found")
	}
	return contracts.ToTodoIDTO(todo), nil
}
