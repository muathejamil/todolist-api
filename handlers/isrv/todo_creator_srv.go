package isrv

import "todolist-api/contracts"

type TodoCreatorSrv interface {
	Create(todoDTO contracts.TodoDTO) (contracts.TodoDTO, error)
}
