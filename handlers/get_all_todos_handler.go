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
