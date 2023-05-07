package models

import (
	"gorm.io/gorm"
	"time"
)

type Todo struct {
	gorm.Model
	Title       string     `json:"Title"`
	Description string     `json:"Description"`
	DueDay      *time.Time `json:"DueDay"`
}

func NewTodo(title string, description string, dueDate *time.Time) Todo {
	return Todo{
		Title:       title,
		Description: description,
		DueDay:      dueDate,
	}
}
