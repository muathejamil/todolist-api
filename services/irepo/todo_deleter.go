package irepo

type TodoDeleter interface {
	DeleteTodo(id uint) error
}
