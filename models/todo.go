package models

import (
	"gorm.io/gorm"
	"time"
)

// Todo todo struct.
type Todo struct {
	gorm.Model
	Title       string     `json:"Title"`
	Description string     `json:"Description"`
	DueDay      *time.Time `json:"DueDay"`
}

// NewTodo creates new Todo.
func NewTodo(id uint, title string, description string, dueDate *time.Time) Todo {
	return Todo{
		Title:       title,
		Description: description,
		DueDay:      dueDate,
	}
}
