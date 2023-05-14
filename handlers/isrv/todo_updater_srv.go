package isrv

import "todolist-api/handlers/contracts"

// TodoUpdaterSrv todo updater service.
type TodoUpdaterSrv interface {
	Update(id uint, todoDto contracts.TodoDTO) (contracts.TodoDTO, error)
}
