package handlers

import (
	"net/http"
	"todolist-api/utils"
)

// GetAll get all todos handler.
// Params writer http.ResponseWriter, request *http.Request.
// Returns slice of todos.
func (t *TodosHandler) GetAll(writer http.ResponseWriter, _ *http.Request) {
	todos, err := t.GetAllTodos.GetAll()
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}
	utils.WriteJsonTodosResponse(writer, todos)
}
