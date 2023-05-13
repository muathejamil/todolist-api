package handlers

import (
	"net/http"
	"todolist-api/contracts"
	"todolist-api/utils"
)

// Update updates todo handler.
// Params writer http.ResponseWriter, request *http.Request
// Returns updated todo.
func (t *TodosHandler) Update(writer http.ResponseWriter, request *http.Request) {
	id := utils.ExtractIdFromURL(writer, request)
	body, done := utils.MapRequestToTodo(writer, request)
	if done {
		return
	}
	todo := contracts.NewTodoDTO(uint(id), body.Title, body.Description, body.DueDay)
	updatedTodo, err := t.UpdateTodo.Update(uint(id), todo)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}
	todos := make([]contracts.TodoDTO, 0)
	dtos := append(todos, updatedTodo)
	utils.WriteJsonTodosResponse(writer, dtos)
}
