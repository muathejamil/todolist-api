package irepo

import (
	"todolist-api/services/contracts"
)

type TodoGetter interface {
	GetTodo(id uint) (contracts.TodoIDTO, error)
}
