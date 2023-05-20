package handlers

import (
	"net/http"
	"regexp"
	"todolist-api/services/contracts"
	"todolist-api/utils"
)

var (
	getAllTodosReg = regexp.MustCompile(`^/api/v1/todos/*$`)
	getTodoReg     = regexp.MustCompile(`^/api/v1/todos/(\d+)$`)
	createTodoReg  = regexp.MustCompile(`^/api/v1/todos/*$`)
	updateTodoReg  = regexp.MustCompile(`^/api/v1/todos/(\d+)$`)
	deleteTodoReg  = regexp.MustCompile(`^/api/v1/todos/(\d+)$`)
)

// TodosHandler todoHandler.
// CreateTodo create todo service.
// DeleteTodo delete todo service.
// GetTodo get one todo service.
// GetAllTodos get all todos service.
// UpdateTodo update todo service.
// TodoCache todo redis cache.
type TodosHandler struct {
	CreateTodo  TodoCreatorSrv
	DeleteTodo  TodoDeleterSrv
	GetTodo     TodoGetterSrv
	GetAllTodos TodosGetterSrv
	UpdateTodo  TodoUpdaterSrv
}

// ServeHTTP serve the http requests and redirect them to the desired service.
func (t *TodosHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("content-type", "application/json")
	path := request.URL.Path
	switch {
	case request.Method == http.MethodPost && createTodoReg.MatchString(path):
		// extracts the todo body form request.
		dto, _ := utils.MapRequestToTodoDto(writer, request)
		// map to dto request idto.
		idto := contracts.ToTodoIDTO(dto)
		t.CreateTodo.Create(idto)
		// Get the value from handle method.
		// write the response to the writer.
		return
	case request.Method == http.MethodGet && getAllTodosReg.MatchString(path):
		t.GetAllTodos.GetAll()
		return
	case request.Method == http.MethodGet && getTodoReg.MatchString(path):
		id := utils.ExtractIdFromURL(writer, request)
		t.GetTodo.Get(uint(id))
		return
	case request.Method == http.MethodPut && updateTodoReg.MatchString(path):
		dto, _ := utils.MapRequestToTodoDto(writer, request)
		idto := contracts.ToTodoIDTO(dto)
		id := utils.ExtractIdFromURL(writer, request)
		t.UpdateTodo.Update(uint(id), idto)
		return
	case request.Method == http.MethodDelete && deleteTodoReg.MatchString(path):
		id := utils.ExtractIdFromURL(writer, request)
		t.DeleteTodo.Delete(uint(id))
		return
	default:
		NotFound(writer, request)
	}

}

// NotFound returns response error when not found in json.
func NotFound(writer http.ResponseWriter, _ *http.Request) {
	writer.WriteHeader(http.StatusNotFound)
	_, err := writer.Write([]byte(`{"Error" : "Not supported server http method!"}`))
	if err != nil {
		return
	}
}
