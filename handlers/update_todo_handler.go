package handlers

import (
	srvContracts "todolist-api/services/contracts"
)

// TodoUpdaterSrv todo updater service interface.
type TodoUpdaterSrv interface {
	Update(id uint, todoIDto srvContracts.TodoIDTO) (srvContracts.TodoIDTO, error)
}

// UpdaterSrv update todo service.
type UpdaterSrv struct {
	serv TodoUpdaterSrv
}

// Update updates todo handler.
// Params writer http.ResponseWriter, request *http.Request
// Returns updated todo.
//func (u *UpdaterSrv) Update(writer http.ResponseWriter, request *http.Request) {
//	id := utils.ExtractIdFromURL(writer, request)
//	body, done := utils.MapRequestToTodo(writer, request)
//	if done {
//		return
//	}
//	todo := contracts.NewTodoDTO(uint(id), body.Title, body.Description, body.DueDay)
//	updatedTodo, err := u.serv.Update(uint(id), todo)
//	if err != nil {
//		http.Error(writer, err.Error(), http.StatusBadRequest)
//		return
//	}
//	// make new todoDto slice.
//	todos := make([]contracts.TodoDTO, 0)
//	// append value to todoDto slice.
//	dtos := append(todos, updatedTodo)
//	// write response as json to writer.
//	utils.WriteJsonTodosResponse(writer, dtos)
//}
