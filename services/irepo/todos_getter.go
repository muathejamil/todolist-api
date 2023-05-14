package irepo

import (
	"todolist-api/services/contracts"
)

type TodosGetter interface {
	GetAllTodos() ([]contracts.TodoIDTO, error)
}
