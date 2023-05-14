package irepo

import "todolist-api/handlers/contracts"

type TodoCreator interface {
	CreateTodo(contracts.TodoIDTO) (contracts.TodoIDTO, error)
}
