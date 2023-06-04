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
