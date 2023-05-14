package isrv

import "todolist-api/contracts"

// TodosGetterSrv get all todos service.
type TodosGetterSrv interface {
	GetAll() ([]contracts.TodoDTO, error)
}
