package repo

import (
	"gorm.io/gorm"
	"log"
	"todolist-api/models"
)

type TodoRepo struct {
	DB *gorm.DB
}

// All methods must return errors.
func (t *TodoRepo) Create(todo models.Todo) models.Todo {
	// Id is our responsibility not DB.
	dbTodo := t.DB.Create(&todo)
	if dbTodo.Error != nil {
		//We should not fatal it will stop the service.
		log.Fatal("Error while creating a new todo!")
	}
	return todo
}

func (t *TodoRepo) Get(id uint) models.Todo {
	var todo models.Todo
	t.DB.Find(&todo, id)
	return todo
}

func (t *TodoRepo) GetAll() []models.Todo {
	//TODO: support pagination and sorting
	var todos []models.Todo
	t.DB.Find(&todos)
	return todos
}

func (t *TodoRepo) Update(id uint, updatedTodo models.Todo) models.Todo {
	//TODO: We can use patch here if we want to update single fields.
	// Get old todo from db
	var todo models.Todo
	t.DB.Find(&todo, id)
	//Check if todo not exist
	if &todo == nil {
		log.Fatal("Todo not found in the database.")
	}

	// update multiple columns
	t.DB.Model(&todo).Updates(
		models.Todo{Title: updatedTodo.Title, Description: updatedTodo.Description, DueDay: updatedTodo.DueDay})
	return todo
}

func (t *TodoRepo) Delete(id uint) {
	t.DB.Delete(&models.Todo{}, id)
}
