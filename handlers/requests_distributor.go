package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"regexp"
	commonContracts "todolist-api/contracts"
	"todolist-api/handlers/common"
	"todolist-api/handlers/utils"
	"todolist-api/services/contracts"
	srvContracts "todolist-api/services/contracts"
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
	CreateTodoSrv  TodoCreatorSrv
	DeleteTodoSrv  TodoDeleterSrv
	GetTodoSrv     TodoGetterSrv
	GetAllTodosSrv TodosGetterSrv
	UpdateTodoSrv  TodoUpdaterSrv
}

// ServeHTTP serve the http requests and redirect them to the desired service.
func (t *TodosHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("content-type", "application/json")
	path := request.URL.Path
	switch {
	case request.Method == http.MethodPost && createTodoReg.MatchString(path):
		t.CreateTodoHandler(writer, request)
		return
	case request.Method == http.MethodGet && getAllTodosReg.MatchString(path):
		_, err := t.HandleGetAll(writer, request)
		if err != nil {
			if HandleError(writer, "Unable to get all todos!") {
				return
			}
		}
		return
	case request.Method == http.MethodGet && getTodoReg.MatchString(path):
		_, err := t.HandleGetOne(writer, request)
		if err != nil {
			if HandleError(writer, "Unable to get todo!") {
				return
			}
		}
		return
	case request.Method == http.MethodPut && updateTodoReg.MatchString(path):
		_, err := t.HandleUpdate(writer, request)
		if err != nil {
			if HandleError(writer, "Unable to update todo!") {
				return
			}
		}
		return
	case request.Method == http.MethodDelete && deleteTodoReg.MatchString(path):
		err := t.HandleDelete(writer, request)
		if err != nil {
			if HandleError(writer, "Unable to delete todo!") {
				return
			}
		}
		return
	default:
		NotFound(writer, request)
	}

}

// HandleError handle unexpected errors
func HandleError(writer http.ResponseWriter, msg string) bool {
	s := fmt.Sprintf(`{"Error" : "%s!"}`, msg)
	_, err := writer.Write([]byte(s))
	if err != nil {
		return true
	}
	return false
}

// NotFound returns response error when not found in json.
func NotFound(writer http.ResponseWriter, _ *http.Request) {
	writer.WriteHeader(http.StatusNotFound)
	_, err := writer.Write([]byte(`{"Error" : "Not supported server http method!"}`))
	if err != nil {
		return
	}
}

// CreateTodoHandler handles create todo handler.
// Params writer http.ResponseWriter, request *http.Request.
// Returns created todo.
func (t *TodosHandler) CreateTodoHandler(writer http.ResponseWriter, request *http.Request) {
	// extracts the todo body form request.
	dto, _ := utils.MapRequestToTodoDto(writer, request)
	// map to dto request idto.
	idto := srvContracts.ToTodoIDTO(dto)
	createdTodo, err := t.CreateTodoSrv.Create(idto)
	if err != nil {

	}
	// write json to the response.
	utils.WriteJsonTodosResponse(writer, []commonContracts.TodoDTO{commonContracts.TodoDTO(createdTodo)})
}

// HandleGetAll handles get all todos.
// Params writer http.ResponseWriter, request *http.Request.
// Returns slice of todos.
func (t *TodosHandler) HandleGetAll(_ http.ResponseWriter, _ *http.Request) ([]commonContracts.TodoDTO, error) {
	todos, err := t.GetAllTodosSrv.GetAll()
	if err != nil {
		return nil, errors.New("unable to fetch todos from db")
	}
	todoDTOs := make([]commonContracts.TodoDTO, len(todos))
	for i, todo := range todos {
		todoDTOs[i] = common.ToTodoDTO(todo)
	}
	return todoDTOs, nil
}

// HandleGetOne get todo handler.
// Params writer http.ResponseWriter, request *http.Request.
// Returns todo if exist.
func (t *TodosHandler) HandleGetOne(writer http.ResponseWriter, request *http.Request) (commonContracts.TodoDTO, error) {
	// extract id from the request url
	id := utils.ExtractIdFromURL(writer, request)
	// get todo from service
	todoIdto, err := t.GetTodoSrv.Get(uint(id))
	if err != nil {
		return commonContracts.TodoDTO{}, errors.New("unable to find todo")
	}
	// convert to todo dto
	todoDTO := common.ToTodoDTO(todoIdto)
	return todoDTO, nil
}

// HandleUpdate updates todo handler.
// Params writer http.ResponseWriter, request *http.Request
// Returns updated todo.
func (t *TodosHandler) HandleUpdate(writer http.ResponseWriter, request *http.Request) (commonContracts.TodoDTO, error) {
	// map request to todo dto
	dto, _ := utils.MapRequestToTodoDto(writer, request)
	// convert to todo idto
	idto := contracts.ToTodoIDTO(dto)
	// extract id from url
	id := utils.ExtractIdFromURL(writer, request)
	// update request
	update, err := t.UpdateTodoSrv.Update(uint(id), idto)
	if err != nil {
		return commonContracts.TodoDTO{}, errors.New("unable to update todo")
	}
	return common.ToTodoDTO(update), nil
}

// HandleDelete handle deletes todo handler.
// Params writer http.ResponseWriter, request *http.Request.
// Returns.
func (t *TodosHandler) HandleDelete(writer http.ResponseWriter, request *http.Request) error {
	id := utils.ExtractIdFromURL(writer, request)
	err := t.DeleteTodoSrv.Delete(uint(id))
	if err != nil {
		return errors.New("unable to delete todo")
	}
	return nil
}
