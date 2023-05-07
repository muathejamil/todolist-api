package services

import (
	"todolist-api/models"
	"todolist-api/repo"
)

type TodoListService struct {
	Repo repo.TodoRepo
}

func (t *TodoListService) Create(todo models.Todo) models.Todo {
	return t.Repo.Create(todo)
}

func (t *TodoListService) Get(id uint) models.Todo {
	return t.Repo.Get(id)
}

func (t *TodoListService) GetAll() []models.Todo {
	return t.Repo.GetAll()
}

func (t *TodoListService) Update(id uint, todo models.Todo) models.Todo {
	return t.Repo.Update(id, todo)
}

func (t *TodoListService) Delete(id uint) {
	t.Repo.Delete(id)
}
