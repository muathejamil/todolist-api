package contracts

import (
	"time"
)

type TodoDTO struct {
	Id          uint       `json:"Id"`
	Title       string     `json:"Title"`
	Description string     `json:"Description"`
	DueDay      *time.Time `json:"DueDay"`
}

func NewTodoDTO(id uint, title string, description string, dueDate *time.Time) TodoDTO {
	return TodoDTO{
		Id:          id,
		Title:       title,
		Description: description,
		DueDay:      dueDate,
	}
}

func ToTodoDTO(todo TodoIDTO) TodoDTO {
	return NewTodoDTO(todo.Id, todo.Title, todo.Description, todo.DueDay)
}
