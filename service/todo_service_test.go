package service

import (
	"testing"
	"stack-gen-TODO/models"
)

func TestCreateAndGetAllTodos(t *testing.T) {
	service := NewTodoService()

	todo := models.Todo{
		ID:        "1",
		Title:     "Learn Go",
		Completed: false,
	}

	err := service.Create(todo)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	todos := service.GetAll()
	if len(todos) != 1 {
		t.Fatalf("expected 1 todo, got %d", len(todos))
	}
}
