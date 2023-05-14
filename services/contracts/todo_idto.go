package contracts

import (
	"time"
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
func ToTodoIDTO(todo models.Todo) TodoIDTO {
	return NewTodoIDTO(todo.ID, todo.Title, todo.Description, todo.DueDay)
}

// ToTodo Map from TodoIDTO to Todo model
// params todo type TodoIDTO
// returns models.Todo
func ToTodo(todo TodoIDTO) models.Todo {
	return models.NewTodo(todo.Id, todo.Title, todo.Description, todo.DueDay)
}
