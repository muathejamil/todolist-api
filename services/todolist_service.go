package services

import (
	"todolist-api/repos"
)

type TodoListService struct {
	Repo repos.TodoRepo
}
