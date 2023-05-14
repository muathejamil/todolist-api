package isrv

import "todolist-api/contracts"

// TodoUpdaterSrv todo updater service.
type TodoUpdaterSrv interface {
	Update(id uint, todoDto contracts.TodoDTO) (contracts.TodoDTO, error)
}
