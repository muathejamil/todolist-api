package irepo

import "todolist-api/handlers/contracts"

type TodoGetter interface {
	GetTodo(id uint) (contracts.TodoIDTO, error)
}
