package main

import (
	"net/http"
	"stack-gen-TODO/handlers"
	"stack-gen-TODO/service"
)

func main() {
	todoService := service.NewTodoService()
	todoHandler := handlers.NewTodoHandler(todoService)

	http.HandleFunc("/todos", todoHandler.Todos)
	http.HandleFunc("/todos/", todoHandler.Todos)


	http.ListenAndServe(":8080", nil)
}
