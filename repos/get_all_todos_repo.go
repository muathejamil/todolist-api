package repos

import (
	"errors"
	"todolist-api/contracts"
	"todolist-api/models"
)

// GetAllTodos get all todos
// Params
// returns slice of contracts.TodoIDTO
func (t *TodoRepo) GetAllTodos() ([]contracts.TodoIDTO, error) {
	//TODO: support pagination and sorting
	var todos []models.Todo
	t.DB.Find(&todos)
	if len(todos) == 0 {
		return nil, errors.New("no todos found")
	}
	todoIDTOs := make([]contracts.TodoIDTO, len(todos))
	for i, todo := range todos {
		todoIDTOs[i] = contracts.ToTodoIDTO(todo)
	}
	return todoIDTOs, nil
}
