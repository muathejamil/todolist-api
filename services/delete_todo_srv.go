package services

// Delete delete todo by id
// Params id uint
// returns error
func (s *Service) Delete(id uint) error {
	err := s.DeleteTodo.DeleteTodo(id)
	if err != nil {
		return err
	}
	return nil
}
