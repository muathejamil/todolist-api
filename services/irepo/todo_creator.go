package irepo

import "todolist-api/contracts"

type TodoCreator interface {
	CreateTodo(contracts.TodoIDTO) (contracts.TodoIDTO, error)
}
