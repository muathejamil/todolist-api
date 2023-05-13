package isrv

import "todolist-api/contracts"

type TodoUpdaterSrv interface {
	Update(id uint, todoDto contracts.TodoDTO) (contracts.TodoDTO, error)
}
