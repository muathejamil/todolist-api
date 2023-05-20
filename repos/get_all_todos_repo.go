package repos

import (
	"errors"
	"todolist-api/models"
)

// GetAllTodos get all todos
// Params
// returns slice of contracts.TodoIDTO
func (t *TodoRepo) GetAllTodos() ([]models.Todo, error) {
	//TODO: support pagination and sorting
	var todos []models.Todo
	t.DB.Find(&todos)
	if len(todos) == 0 {
		return nil, errors.New("no todos found")
	}
	return todos, nil
}
