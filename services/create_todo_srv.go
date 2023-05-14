package services

import (
	"todolist-api/handlers/contracts"
	contracts2 "todolist-api/services/contracts"
)

// Create creates a new Todo
// Params todoDTO contracts.TodoDTO
// returns contracts.TodoDTO, error
func (s *Service) Create(todoDTO contracts.TodoDTO) (contracts.TodoDTO, error) {
	todo := contracts2.NewTodoIDTO(todoDTO.Id, todoDTO.Title, todoDTO.Description, todoDTO.DueDay)
	todoIDTO, err := s.CreateTodo.CreateTodo(todo)
	if err != nil {
		return contracts.TodoDTO{}, err
	}
	return contracts.ToTodoDTO(todoIDTO), nil
}
