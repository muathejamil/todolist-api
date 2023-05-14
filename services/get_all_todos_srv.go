package services

import "todolist-api/handlers/contracts"

// GetAll get all todos
// Params
// returns slice of contracts.TodoDTO
func (s *Service) GetAll() ([]contracts.TodoDTO, error) {
	todoIDTOS, err := s.GetAllTodos.GetAllTodos()
	if err != nil {
		return nil, err
	}
	todoDTOs := make([]contracts.TodoDTO, len(todoIDTOS))
	for i, todo := range todoIDTOS {
		todoDTOs[i] = contracts.ToTodoDTO(todo)
	}
	return todoDTOs, nil
}
