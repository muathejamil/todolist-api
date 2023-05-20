package handlers

import (
	"todolist-api/handlers/contracts"
	srvContracts "todolist-api/services/contracts"
)

// TodoUpdaterSrv todo updater service interface.
type TodoUpdaterSrv interface {
	Update(id uint, todoIDto srvContracts.TodoIDTO) (srvContracts.TodoIDTO, error)
}

// UpdaterSrv update todo service.
type UpdaterSrv struct {
	serv TodoUpdaterSrv
}

// Update updates todo handler.
// Params writer http.ResponseWriter, request *http.Request
// Returns updated todo.
func (u *UpdaterSrv) Update(id uint, todo contracts.TodoDTO) (contracts.TodoDTO, error) {
	updatedTodo, err := u.serv.Update(id, srvContracts.ToTodoIDTO(todo))
	if err != nil {
		return contracts.TodoDTO{}, nil
	}
	return contracts.ToTodoDTO(updatedTodo), nil
}
