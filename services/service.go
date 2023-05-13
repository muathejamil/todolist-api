package services

import "todolist-api/services/irepo"

type Service struct {
	CreateTodo  irepo.TodoCreator
	DeleteTodo  irepo.TodoDeleter
	GetTodo     irepo.TodoGetter
	UpdateTodo  irepo.TodoUpdater
	GetAllTodos irepo.TodosGetter
}
