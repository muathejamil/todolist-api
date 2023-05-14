package handlers

import (
	"net/http"
	"todolist-api/contracts"
	"todolist-api/utils"
)

// Get get todo handler.
// Params writer http.ResponseWriter, request *http.Request.
// Returns todo if exist.
func (t *TodosHandler) Get(writer http.ResponseWriter, request *http.Request) {
	id := utils.ExtractIdFromURL(writer, request)
	// Check if the value exist in the cache.
	todo, err := t.TodoCache.Get(id)
	if todo.Id == 0 {
		// Get the value from the service
		todo, err = t.GetTodo.Get(uint(id))
		if err != nil {
			http.Error(writer, err.Error(), http.StatusBadRequest)
			return
		}
		//store the value in the cache.
		t.TodoCache.Set(id, todo)
	}
	// create new todoDTO slice.
	todos := make([]contracts.TodoDTO, 0)
	// append value
	dtos := append(todos, todo)
	// write response as to writer.
	utils.WriteJsonTodosResponse(writer, dtos)
}
