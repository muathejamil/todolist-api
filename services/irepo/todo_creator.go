package irepo

import (
	"todolist-api/services/contracts"
)

type TodoCreator interface {
	CreateTodo(contracts.TodoIDTO) (contracts.TodoIDTO, error)
}
