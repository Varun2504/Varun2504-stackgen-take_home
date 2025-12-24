package handlers

import (
	"encoding/json"
	"net/http"
	"strings"

	"stack-gen-TODO/models"
	"stack-gen-TODO/service"
)

type TodoHandler struct {
	service *service.TodoService
}

func NewTodoHandler(s *service.TodoService) *TodoHandler {
	return &TodoHandler{service: s}
}

func (h *TodoHandler) Todos(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodPost {
	var todo models.Todo
	json.NewDecoder(r.Body).Decode(&todo)
	h.service.Create(todo)
	w.WriteHeader(http.StatusCreated)

} else if r.Method == http.MethodGet {
	todos := h.service.GetAll()
	json.NewEncoder(w).Encode(todos)

} else {
	w.WriteHeader(http.StatusMethodNotAllowed)
}

}

func (h *TodoHandler) DeleteTodo(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/todos/") // /todos/{id}
	h.service.Delete(id)
	w.WriteHeader(http.StatusNoContent)
}
