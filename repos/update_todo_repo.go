package repos

import (
	"errors"
	"todolist-api/handlers/contracts"
	"todolist-api/models"
)

// UpdateTodo update todos
// Params id uint, updatedTodo contracts.TodoIDTO
// Returns updated contracts.TodoIDTO, error
func (t *TodoRepo) UpdateTodo(id uint, updatedTodo contracts.TodoIDTO) (contracts.TodoIDTO, error) {
	// Get old todo from db
	var todo models.Todo
	t.DB.Find(&todo, id)
	//Check if todo not exist
	if &todo == nil {
		return contracts.TodoIDTO{}, errors.New("todo not found")
	}

	// update multiple columns
	t.DB.Model(&todo).Updates(
		models.Todo{Title: updatedTodo.Title, Description: updatedTodo.Description, DueDay: updatedTodo.DueDay})
	return contracts.ToTodoIDTO(todo), nil
}
