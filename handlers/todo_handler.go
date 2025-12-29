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
	switch r.Method {

	case http.MethodGet:
		
		if r.URL.Path != "/todos" && r.URL.Path != "/todos/" {
			http.NotFound(w, r)
			return
		}

		todos := h.service.GetAll()
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(todos)

	case http.MethodPost:
		var todo models.Todo
		if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
			http.Error(w, "invalid request body", http.StatusBadRequest)
			return
		}

		h.service.Create(todo)
		w.WriteHeader(http.StatusCreated)

	case http.MethodDelete:
		id := strings.TrimPrefix(r.URL.Path, "/todos/")
		if id == "" || id == "todos" {
			http.Error(w, "missing todo id", http.StatusBadRequest)
			return
		}

		h.service.Delete(id)
		w.WriteHeader(http.StatusNoContent)

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}