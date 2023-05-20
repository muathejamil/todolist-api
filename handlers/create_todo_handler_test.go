package handlers

import (
	"fmt"
	"testing"
	srvContracts "todolist-api/services/contracts"
)

// MockTodoCreatorSrv is a mock implementation of TodoCreatorSrv interface for testing.
type MockTodoCreatorSrv struct{}

func (m *MockTodoCreatorSrv) Create(todoIDTO srvContracts.TodoIDTO) (srvContracts.TodoIDTO, error) {
	// Implement the mock logic here
	return srvContracts.TodoIDTO{}, nil
}

func TestCreateHandler(t *testing.T) {
	// Create a mock TodoCreatorSrv
	mockSrv := &MockTodoCreatorSrv{}

	// Create a new instance of the Create handler with the mock service
	handler := &CreatorSrv{
		serv: mockSrv,
	}

	todoIDTO := srvContracts.TodoIDTO{}

	createdTodo, err := handler.serv.Create(todoIDTO)

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	fmt.Println(createdTodo)
}
