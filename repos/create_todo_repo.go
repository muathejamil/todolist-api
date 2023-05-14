package repos

import (
	"todolist-api/handlers/contracts"
	"todolist-api/models"
)

// CreateTodo creates a new Todo
// Params todoIDTO contracts.TodoIDTO
// returns contracts.TodoIDTO, error
func (t *TodoRepo) CreateTodo(todoIDTO contracts.TodoIDTO) (contracts.TodoIDTO, error) {
	// TODO: ID is our responsibility not DB.
	todo := models.NewTodo(todoIDTO.Id, todoIDTO.Title, todoIDTO.Description, todoIDTO.DueDay)
	dbTodo := t.DB.Create(&todo)
	if dbTodo.Error != nil {
		return contracts.TodoIDTO{}, dbTodo.Error
	}
	return contracts.ToTodoIDTO(todo), nil
}
