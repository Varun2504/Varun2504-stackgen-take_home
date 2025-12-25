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
		
		case http.MethodPost:
			var todo models.Todo
			if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
				http.Error(w, "invalid request body", http.StatusBadRequest)
				return
			}

			h.service.Create(todo)
			w.WriteHeader(http.StatusCreated)

		case http.MethodGet:
			todos := h.service.GetAll()
			if err := json.NewEncoder(w).Encode(todos); err != nil {
				http.Error(w, "failed to encode response", http.StatusInternalServerError)
			}

		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}


}

func (h *TodoHandler) DeleteTodo(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/todos/") // /todos/{id}
	h.service.Delete(id)
	w.WriteHeader(http.StatusNoContent)
}
