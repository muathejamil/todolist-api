package common

import (
	contracts2 "todolist-api/contracts"
	"todolist-api/services/contracts"
)

// ToTodoDTO maps todoIdto to TodoDto
func ToTodoDTO(todo contracts.TodoIDTO) contracts2.TodoDTO {
	return contracts2.NewTodoDTO(todo.Id, todo.Title, todo.Description, todo.DueDay)
}
