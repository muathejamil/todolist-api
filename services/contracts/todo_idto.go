package contracts

import (
	"time"
	"todolist-api/handlers/contracts"
	"todolist-api/models"
)

type TodoIDTO struct {
	Id          uint       `json:"Id"`
	Title       string     `json:"Title"`
	Description string     `json:"Description"`
	DueDay      *time.Time `json:"DueDay"`
}

// NewTodoIDTO creates new TodoIDTO
func NewTodoIDTO(id uint, title string, description string, dueDate *time.Time) TodoIDTO {
	return TodoIDTO{
		Id:          id,
		Title:       title,
		Description: description,
		DueDay:      dueDate,
	}
}

// ToTodoIDTO Map from Todo model to TodoIDTO
// params todo type models.Todo
// returns TodoIDTO
func ToTodoIDTO(todo contracts.TodoDTO) TodoIDTO {
	return NewTodoIDTO(todo.Id, todo.Title, todo.Description, todo.DueDay)
}

// ToIDTO Map from TodoIDTO to Todo model
// params todo type TodoIDTO
// returns TodoIDTO
func ToIDTO(todo models.Todo) TodoIDTO {
	return NewTodoIDTO(todo.ID, todo.Title, todo.Description, todo.DueDay)
}
