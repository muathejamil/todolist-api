package isrv

import "todolist-api/contracts"

type TodoGetterSrv interface {
	Get(id uint) (contracts.TodoDTO, error)
}
