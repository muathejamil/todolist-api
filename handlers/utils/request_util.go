package utils

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"time"
	"todolist-api/contracts"
)

// MapRequestToTodoDto maps the request body to todo struct.
// Params writer http.ResponseWriter, request *http.Request.
// Returns body, false.
func MapRequestToTodoDto(writer http.ResponseWriter, request *http.Request) (contracts.TodoDTO, bool) {
	var body struct {
		Id          uint
		Title       string
		Description string
		StartDay    *time.Time
		DueDay      *time.Time
	}
	err := json.NewDecoder(request.Body).Decode(&body)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return contracts.TodoDTO{}, true
	}
	dto := contracts.NewTodoDTO(body.Id, body.Title, body.Description, body.DueDay)
	return dto, false
}

// WriteJsonTodosResponse writes the response to the response body.
// Params writer http.ResponseWriter, todos []contracts.TodoDTO.
// Returns
func WriteJsonTodosResponse(writer http.ResponseWriter, todos []contracts.TodoDTO) {
	var err error
	if len(todos) == 1 {
		dto := todos[0]
		err = json.NewEncoder(writer).Encode(dto)

	} else {
		err = json.NewEncoder(writer).Encode(todos)
	}
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
}

// ExtractIdFromURL extracts the id from the request url.
// Params writer http.ResponseWriter, request *http.Request
// Returns id.
func ExtractIdFromURL(writer http.ResponseWriter, request *http.Request) uint64 {
	idStr := strings.TrimPrefix(request.URL.Path, "/api/v1/todos/")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		// handle error, such as returning an HTTP error response
		http.Error(writer, "Invalid ID", http.StatusBadRequest)
	}
	return id
}
