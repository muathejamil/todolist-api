package irepo

import "todolist-api/handlers/contracts"

type TodoUpdater interface {
	UpdateTodo(id uint, updatedTodo contracts.TodoIDTO) (contracts.TodoIDTO, error)
}
