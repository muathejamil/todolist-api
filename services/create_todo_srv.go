package services

import (
	"todolist-api/models"
	"todolist-api/services/contracts"
)

// TodoCreator create todo interface.
type TodoCreator interface {
	CreateTodo(todo models.Todo) (models.Todo, error)
}

// Creator create todo struct.
type Creator struct {
	repo TodoCreator
}

// Create creates a new Todo
// Params todoDTO contracts.TodoDTO
// returns contracts.TodoDTO, error
func (s *Creator) Create(todoIDTO contracts.TodoIDTO) (contracts.TodoIDTO, error) {
	dbTodo := models.NewTodo(todoIDTO.Id, todoIDTO.Title, todoIDTO.Description, todoIDTO.DueDay)
	todo, err := s.repo.CreateTodo(dbTodo)
	if err != nil {
		return contracts.TodoIDTO{}, err
	}
	return contracts.ToIDTO(todo), nil
}
