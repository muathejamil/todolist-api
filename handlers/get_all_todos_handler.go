package handlers

import (
	srvContracts "todolist-api/services/contracts"
)

// TodosGetterSrv get all todos service.
type TodosGetterSrv interface {
	GetAll() ([]srvContracts.TodoIDTO, error)
}

// GetterSrv get all todos service.
type GetterSrv struct {
	serv TodosGetterSrv
}

//// GetAll get all todos handler.
//// Params writer http.ResponseWriter, request *http.Request.
//// Returns slice of todos.
//func (g *GetterSrv) GetAll(writer http.ResponseWriter, _ *http.Request) {
//	todos, err := g.serv.GetAll()
//	if err != nil {
//		http.Error(writer, err.Error(), http.StatusBadRequest)
//		return
//	}
//	utils.WriteJsonTodosResponse(writer, todos)
//}
