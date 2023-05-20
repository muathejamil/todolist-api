package handlers

import (
	srvContracts "todolist-api/services/contracts"
)

// TodoCreatorSrv todo creator service interface.
type TodoCreatorSrv interface {
	Create(todoIDTO srvContracts.TodoIDTO) (srvContracts.TodoIDTO, error)
}

// CreatorSrv todo creator service.
type CreatorSrv struct {
	serv TodoCreatorSrv
}

//// Create creates todo handler.
//// Params writer http.ResponseWriter, request *http.Request.
//// Returns created todo.
//func (t *CreatorSrv) Create(writer http.ResponseWriter, request *http.Request) {
//	// map request to todo struct.
//	body, done := utils.MapRequestToTodoDto(writer, request)
//	if done {
//		// add the status code to the writer.
//		http.Error(writer, "Bad request", http.StatusBadRequest)
//		return
//	}
//	// creates new todoDto.
//	todo := contracts.NewTodoDTO(0, body.Title, body.Description, body.DueDay)
//	response, err := t.serv.Create(todo)
//	if err != nil {
//		http.Error(writer, err.Error(), http.StatusBadRequest)
//		return
//	}
//	// make new todoDto slice.
//	todos := make([]contracts.TodoDTO, 0)
//	// append the value.
//	dtos := append(todos, response)
//	// write response json to writer.
//	utils.WriteJsonTodosResponse(writer, dtos)
//}
