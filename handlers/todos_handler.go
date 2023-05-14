package handlers

import (
	"net/http"
	"regexp"
	"todolist-api/cache"
	"todolist-api/handlers/isrv"
)

var (
	getAllTodosReg = regexp.MustCompile(`^/api/v1/todos/*$`)
	getTodoReg     = regexp.MustCompile(`^/api/v1/todos/(\d+)$`)
	createTodoReg  = regexp.MustCompile(`^/api/v1/todos/*$`)
	updateTodoReg  = regexp.MustCompile(`^/api/v1/todos/(\d+)$`)
	deleteTodoReg  = regexp.MustCompile(`^/api/v1/todos/(\d+)$`)
)

// TodosHandler todoHandler.
type TodosHandler struct {
	CreateTodo  isrv.TodoCreatorSrv
	DeleteTodo  isrv.TodoDeleterSrv
	GetTodo     isrv.TodoGetterSrv
	GetAllTodos isrv.TodosGetterSrv
	UpdateTodo  isrv.TodoUpdaterSrv
	TodoCache   cache.TodoCache
}

// ServeHTTP serve the http requests and redirect them to the desired service.
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

// NotFound returns response error when not found in json.
func NotFound(writer http.ResponseWriter, _ *http.Request) {
	writer.WriteHeader(http.StatusNotFound)
	_, err := writer.Write([]byte(`{"Error" : "Not supported server http method!"}`))
	if err != nil {
		return
	}
}
