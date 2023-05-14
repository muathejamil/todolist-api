package repos

import (
	"todolist-api/models"
)

// DeleteTodo delete todo by id
// Params id uint
// returns error
func (t *TodoRepo) DeleteTodo(id uint) error {
	result := t.DB.Delete(&models.Todo{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
