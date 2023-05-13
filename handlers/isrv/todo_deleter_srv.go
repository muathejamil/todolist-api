package isrv

type TodoDeleterSrv interface {
	Delete(id uint) error
}
