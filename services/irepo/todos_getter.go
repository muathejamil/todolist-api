package irepo

import "todolist-api/handlers/contracts"

type TodosGetter interface {
	GetAllTodos() ([]contracts.TodoIDTO, error)
}
