package isrv

// TodoDeleterSrv todo deleter service.
type TodoDeleterSrv interface {
	Delete(id uint) error
}
