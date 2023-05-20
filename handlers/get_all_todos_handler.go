package handlers

import (
	"errors"
	"todolist-api/handlers/contracts"
	srvContracts "todolist-api/services/contracts"
)

// TodosGetterSrv get all todos service.
type TodosGetterSrv interface {
	GetAll() ([]srvContracts.TodoIDTO, error)
}

// GetterSrv get all todos service.
type GetterSrv struct {
	serv TodosGetterSrv
}

// HandleGetAll handles get all todos.
// Params writer http.ResponseWriter, request *http.Request.
// Returns slice of todos.
func (g *GetterSrv) HandleGetAll() ([]contracts.TodoDTO, error) {
	todos, err := g.serv.GetAll()
	if err != nil {
		return nil, errors.New("unable to fetch todos from db")
	}
	todoDTOs := make([]contracts.TodoDTO, len(todos))
	for i, todo := range todos {
		todoDTOs[i] = contracts.ToTodoDTO(todo)
	}
	return todoDTOs, nil
}
