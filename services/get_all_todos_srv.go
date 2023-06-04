package services

import (
	"todolist-api/models"
	"todolist-api/services/contracts"
)

// TodosGetter gets all todos interface.
type TodosGetter interface {
	GetAllTodos() ([]models.Todo, error)
}

// AllGetter gets all todos struct.
type AllGetter struct {
	Repo TodosGetter
}

// GetAll get all todos
// Params
// returns slice of contracts.TodoDTO
func (g *AllGetter) GetAll() ([]contracts.TodoIDTO, error) {
	todoIDTOS, err := g.Repo.GetAllTodos()
	if err != nil {
		return nil, err
	}
	todoIDTOs := make([]contracts.TodoIDTO, len(todoIDTOS))
	for i, todo := range todoIDTOS {
		todoIDTOs[i] = contracts.ToIDTO(todo)
	}
	return todoIDTOs, nil
}
