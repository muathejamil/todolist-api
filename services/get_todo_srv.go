package services

import "todolist-api/contracts"

// Get gets todo by id
// Params id uint
// returns contracts.TodoIDTO, error
func (s *Service) Get(id uint) (contracts.TodoDTO, error) {
	todoIDTO, err := s.GetTodo.GetTodo(id)
	if err != nil {
		return contracts.TodoDTO{}, err
	}
	return contracts.ToTodoDTO(todoIDTO), nil
}
