package services

import (
	"todolist-api/repo"
)

type TodoListService struct {
	Repo repo.TodoRepo
}
