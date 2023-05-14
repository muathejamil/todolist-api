package irepo

import (
	"todolist-api/services/contracts"
)

// TodoCreator create todo interface.
type TodoCreator interface {
	CreateTodo(contracts.TodoIDTO) (contracts.TodoIDTO, error)
}
