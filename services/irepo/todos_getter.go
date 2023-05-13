package irepo

import "todolist-api/contracts"

type TodosGetter interface {
	GetAllTodos() ([]contracts.TodoIDTO, error)
}
