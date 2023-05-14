package irepo

import (
	"todolist-api/services/contracts"
)

type TodoUpdater interface {
	UpdateTodo(id uint, updatedTodo contracts.TodoIDTO) (contracts.TodoIDTO, error)
}
