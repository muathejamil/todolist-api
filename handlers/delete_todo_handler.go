package handlers

import (
	"net/http"
	"todolist-api/utils"
)

// Delete deletes todo handler.
// Params writer http.ResponseWriter, request *http.Request.
// Returns.
func (t *TodosHandler) Delete(writer http.ResponseWriter, request *http.Request) {
	id := utils.ExtractIdFromURL(writer, request)
	err := t.DeleteTodo.Delete(uint(id))
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}
}
