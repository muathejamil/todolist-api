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
