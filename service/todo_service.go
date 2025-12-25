package service

import (
	"stack-gen-TODO/models";
	"errors"
	)

type TodoService struct {
	todos map[string]models.Todo // In-memory store for todos
}

func NewTodoService() *TodoService {
	return &TodoService{
		todos: make(map[string]models.Todo), 
	}
}

func (s *TodoService) Create(todo models.Todo) error {
	if todo.ID == "" {
		return errors.New("id cannot be empty")
	}
	if todo.Title == "" {
		return errors.New("title cannot be empty")
	}

	s.todos[todo.ID] = todo
	
	return nil
} // With validation


func (s *TodoService) GetAll() []models.Todo {
	result := []models.Todo{}
	for _, todo := range s.todos {
		result = append(result, todo)
	}
	return result
}

func (s *TodoService) Delete(id string) error {
	if _, exists := s.todos[id]; !exists {
		return errors.New("invalid id")
	}
	delete(s.todos, id)
	return nil
}
