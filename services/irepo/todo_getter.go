package irepo

import "todolist-api/contracts"

type TodoGetter interface {
	GetTodo(id uint) (contracts.TodoIDTO, error)
}