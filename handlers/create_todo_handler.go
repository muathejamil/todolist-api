package handlers

import (
	"net/http"
	"todolist-api/contracts"
	"todolist-api/utils"
)

// Create creates todo handler.
// Params writer http.ResponseWriter, request *http.Request.
// Returns created todo.
func (t *TodosHandler) Create(writer http.ResponseWriter, request *http.Request) {
	body, done := utils.MapRequestToTodo(writer, request)
	if done {
		// add the status code to the writer.
		http.Error(writer, "Bad request", http.StatusBadRequest)
		return
	}
	todo := contracts.NewTodoDTO(0, body.Title, body.Description, body.DueDay)
	response, err := t.CreateTodo.Create(todo)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}
	todos := make([]contracts.TodoDTO, 0)
	dtos := append(todos, response)
	utils.WriteJsonTodosResponse(writer, dtos)
}
