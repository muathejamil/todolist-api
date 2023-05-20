package contracts

import (
	"google.golang.org/protobuf/reflect/protoreflect"
	"time"
	"todolist-api/services/contracts"
)

type TodoDTO struct {
	Id          uint       `json:"Id"`
	Title       string     `json:"Title"`
	Description string     `json:"Description"`
	DueDay      *time.Time `json:"DueDay"`
}

func (t TodoDTO) ProtoReflect() protoreflect.Message {
	//TODO implement me
	panic("implement me")
}

// NewTodoDTO creates new todoDto
func NewTodoDTO(id uint, title string, description string, dueDate *time.Time) TodoDTO {
	return TodoDTO{
		Id:          id,
		Title:       title,
		Description: description,
		DueDay:      dueDate,
	}
}

// ToTodoDTO maps todoIdto to TodoDto
func ToTodoDTO(todo contracts.TodoIDTO) TodoDTO {
	return NewTodoDTO(todo.Id, todo.Title, todo.Description, todo.DueDay)
}
