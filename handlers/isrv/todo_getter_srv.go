package isrv

import "todolist-api/contracts"

// TodoGetterSrv todo getter service.
type TodoGetterSrv interface {
	Get(id uint) (contracts.TodoDTO, error)
}
