package handlers

import (
	"encoding/json"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"
	"todolist-api/models"
	"todolist-api/services"
)

var (
	getAllTodosReg = regexp.MustCompile(`^/api/v1/todos/*$`)
	getTodoReg     = regexp.MustCompile(`^/api/v1/todos/(\d+)$`)
	createTodoReg  = regexp.MustCompile(`^/api/v1/todos/*$`)
	updateTodoReg  = regexp.MustCompile(`^/api/v1/todos/(\d+)$`)
	deleteTodoReg  = regexp.MustCompile(`^/api/v1/todos/(\d+)$`)
)

type TodosHandler struct {
	TodolistService services.TodoListService
}

func (t *TodosHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("content-type", "application/json")
	path := request.URL.Path
	switch {
	case request.Method == http.MethodPost && createTodoReg.MatchString(path):
		t.Create(writer, request)
		return
	case request.Method == http.MethodGet && getAllTodosReg.MatchString(path):
		t.GetAll(writer, request)
		return
	case request.Method == http.MethodGet && getTodoReg.MatchString(path):
		t.Get(writer, request)
		return
	case request.Method == http.MethodPut && updateTodoReg.MatchString(path):
		t.Update(writer, request)
		return
	case request.Method == http.MethodDelete && deleteTodoReg.MatchString(path):
		t.Delete(writer, request)
		return
	default:
		NotFound(writer, request)
	}

}

//TODO; Add docs to methods

func (t *TodosHandler) Create(writer http.ResponseWriter, request *http.Request) {
	body, done := mapRequestToTodo(writer, request)
	if done {
		// add the status code to the writer.
		return
	}
	todo := models.NewTodo(body.Title, body.Description, body.DueDay)
	response := t.TodolistService.Create(todo)
	writeJsonTodoResponse(writer, response)
}

func (t *TodosHandler) Get(writer http.ResponseWriter, request *http.Request) {
	id := extractIdFromURL(writer, request)
	//Handle error Db connection as example
	todo := t.TodolistService.Get(uint(id))
	writeJsonTodoResponse(writer, todo)
}

func (t *TodosHandler) GetAll(writer http.ResponseWriter, _ *http.Request) {
	todos := t.TodolistService.GetAll()
	writeJsonTodosResponse(writer, todos)
}

func (t *TodosHandler) Update(writer http.ResponseWriter, request *http.Request) {
	id := extractIdFromURL(writer, request)
	body, done := mapRequestToTodo(writer, request)
	if done {
		return
	}
	todo := models.NewTodo(body.Title, body.Description, body.DueDay)
	updatedTodo := t.TodolistService.Update(uint(id), todo)
	writeJsonTodoResponse(writer, updatedTodo)
}

func (t *TodosHandler) Delete(writer http.ResponseWriter, request *http.Request) {
	id := extractIdFromURL(writer, request)
	t.TodolistService.Delete(uint(id))
}

func NotFound(writer http.ResponseWriter, _ *http.Request) {
	writer.WriteHeader(http.StatusNotFound)
	_, err := writer.Write([]byte(`{"Error" : "Not supported server http method!"}`))
	if err != nil {
		return
	}
}

func mapRequestToTodo(writer http.ResponseWriter, request *http.Request) (struct {
	Title       string
	Description string
	StartDay    *time.Time
	DueDay      *time.Time
}, bool) {
	var body struct {
		Title       string
		Description string
		StartDay    *time.Time
		DueDay      *time.Time
	}
	err := json.NewDecoder(request.Body).Decode(&body)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return struct {
			Title       string
			Description string
			StartDay    *time.Time
			DueDay      *time.Time
		}{}, true
	}
	return body, false
}

func extractIdFromURL(writer http.ResponseWriter, request *http.Request) uint64 {
	idStr := strings.TrimPrefix(request.URL.Path, "/api/v1/todos/")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		// handle error, such as returning an HTTP error response
		http.Error(writer, "Invalid ID", http.StatusBadRequest)
	}
	return id
}

// We can get rid of one of these methods and use slice len to make sure 1 or multiple.
func writeJsonTodoResponse(writer http.ResponseWriter, todo models.Todo) {
	if err := json.NewEncoder(writer).Encode(todo); err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
}

func writeJsonTodosResponse(writer http.ResponseWriter, todos []models.Todo) {
	if err := json.NewEncoder(writer).Encode(todos); err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
}
