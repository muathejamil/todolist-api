package handlers

// TodoDeleterSrv todo deleter service interface.
type TodoDeleterSrv interface {
	Delete(id uint) error
}

// DeleterSrv todo deleter service interface.
type DeleterSrv struct {
	serv TodoDeleterSrv
}
