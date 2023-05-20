package handlers

import (
	"testing"
)

type MockTodoDeleterSrv struct{}

func (m *MockTodoDeleterSrv) Delete(id uint) error {
	return nil
}

func TestDeleteHandler(t *testing.T) {
	mockSrv := &MockTodoDeleterSrv{}

	handler := &DeleterSrv{
		serv: mockSrv,
	}

	id := uint(123)

	err := handler.serv.Delete(id)

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

}
