package services

import (
	"todolist-api/models"
	"todolist-api/services/contracts"
)

// TodoUpdater update todo interface
type TodoUpdater interface {
	UpdateTodo(id uint, updatedTodo models.Todo) (models.Todo, error)
}

// Updater update todo struct
type Updater struct {
	repo TodoUpdater
}

// Update update todos
// Params id uint, todoDto contracts.TodoDTO
// Returns updated contracts.TodoDTO, error
func (u *Updater) Update(id uint, todoIDTO contracts.TodoIDTO) (contracts.TodoIDTO, error) {
	dbTodo := models.NewTodo(todoIDTO.Id, todoIDTO.Title, todoIDTO.Description, todoIDTO.DueDay)
	todo, err := u.repo.UpdateTodo(id, dbTodo)
	if err != nil {
		return contracts.TodoIDTO{}, err
	}
	return contracts.ToIDTO(todo), nil
}
