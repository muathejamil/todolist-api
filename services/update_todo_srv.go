package services

import "todolist-api/handlers/contracts"

// Update update todos
// Params id uint, todoDto contracts.TodoDTO
// Returns updated contracts.TodoDTO, error
func (s *Service) Update(id uint, todoDto contracts.TodoDTO) (contracts.TodoDTO, error) {
	todo := contracts.NewTodoIDTO(todoDto.Id, todoDto.Title, todoDto.Description, todoDto.DueDay)
	todoIDTO, err := s.UpdateTodo.UpdateTodo(id, todo)
	if err != nil {
		return contracts.TodoDTO{}, err
	}
	return contracts.ToTodoDTO(todoIDTO), nil
}
