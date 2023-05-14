package isrv

import "todolist-api/handlers/contracts"

// TodosGetterSrv get all todos service.
type TodosGetterSrv interface {
	GetAll() ([]contracts.TodoDTO, error)
}
