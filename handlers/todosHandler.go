package handlers

import (
	"net/http"
	"regexp"
	"todolist-api/services"
)

var (
	getAllTodosReg = regexp.MustCompile(`^\/todos[\/]*$`)
	getTodoReg     = regexp.MustCompile(`^\/todos\/(\d+)$`)
	createTodoReg  = regexp.MustCompile(`^\/todos[\/]*$`)
	updateTodoReg  = regexp.MustCompile(`^\/todos\/(\d+)$`)
	deleteTodoReg  = regexp.MustCompile(`^\/todos\/(\d+)$`)
)

type TodosHandler struct {
	todolistService services.TodoListService
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

func (t *TodosHandler) Create(writer http.ResponseWriter, request *http.Request) {
	//TODO: extract the body from the request and send it to the service
	t.todolistService.Create()
}

func (t *TodosHandler) Get(writer http.ResponseWriter, request *http.Request) {
	//TODO: extract the body from the request and send it to the service
	t.todolistService.Get()
}

func (t *TodosHandler) GetAll(writer http.ResponseWriter, request *http.Request) {
	//TODO: extract the body from the request and send it to the service
	t.todolistService.GetAll()
}

func (t *TodosHandler) Update(writer http.ResponseWriter, request *http.Request) {
	//TODO: extract the body from the request and send it to the service
	t.todolistService.Update()
}

func (t *TodosHandler) Delete(writer http.ResponseWriter, request *http.Request) {
	//TODO: extract the body from the request and send it to the service
	t.todolistService.Delete()
}

func NotFound(writer http.ResponseWriter, request *http.Request) {
	writer.WriteHeader(http.StatusNotFound)
	writer.Write([]byte(`{"Error" : "Not supported server http method!"}`))
}
