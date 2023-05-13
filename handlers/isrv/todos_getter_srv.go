package isrv

import "todolist-api/contracts"

type TodosGetterSrv interface {
	GetAll() ([]contracts.TodoDTO, error)
}
