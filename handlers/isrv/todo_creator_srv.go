package isrv

import "todolist-api/handlers/contracts"

// TodoCreatorSrv todo creator service.
type TodoCreatorSrv interface {
	Create(todoDTO contracts.TodoDTO) (contracts.TodoDTO, error)
}
