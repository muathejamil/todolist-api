package handlers

import (
	"errors"
)

// TodoDeleterSrv todo deleter service interface.
type TodoDeleterSrv interface {
	Delete(id uint) error
}

// DeleterSrv todo deleter service interface.
type DeleterSrv struct {
	serv TodoDeleterSrv
}

// HandleDelete handle deletes todo handler.
// Params writer http.ResponseWriter, request *http.Request.
// Returns.
func (d *DeleterSrv) HandleDelete(id uint) error {
	err := d.serv.Delete(uint(id))
	if err != nil {
		return errors.New("unable to delete todo")
	}
	return nil
}
