package services

// TodoDeleter delete todo interface.
type TodoDeleter interface {
	DeleteTodo(id uint) error
}

// Deleter delete todo struct.
type Deleter struct {
	Repo TodoDeleter
}

// Delete delete todo by id
// Params id uint
// returns error
func (d *Deleter) Delete(id uint) error {
	err := d.Repo.DeleteTodo(id)
	if err != nil {
		return err
	}
	return nil
}
