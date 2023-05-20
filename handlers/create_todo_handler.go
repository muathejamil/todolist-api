package handlers

import (
	"errors"
	"todolist-api/handlers/contracts"
	srvContracts "todolist-api/services/contracts"
)

// TodoCreatorSrv todo creator service interface.
type TodoCreatorSrv interface {
	Create(todoIDTO srvContracts.TodoIDTO) (srvContracts.TodoIDTO, error)
}

// CreatorSrv todo creator service.
type CreatorSrv struct {
	serv TodoCreatorSrv
}

// HandleCreate handles create todo handler.
// Params writer http.ResponseWriter, request *http.Request.
// Returns created todo.
func (t *CreatorSrv) HandleCreate(todo contracts.TodoDTO) (contracts.TodoDTO, error) {
	// create new todo.
	response, err := t.serv.Create(srvContracts.ToTodoIDTO(todo))
	if err != nil {
		return contracts.TodoDTO{}, errors.New("unable to create todo")
	}
	return contracts.ToTodoDTO(response), nil
}
