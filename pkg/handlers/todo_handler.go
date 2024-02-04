package handlers

import (
	service "github.com/satoshi-tahara-st/todo_backend/pkg/service/todo"
)

type TodoHandler struct {
	service service.ITodoService
}

func NewTodoHandler(service service.ITodoService) TodoHandler {
	return TodoHandler{service: service}
}

func (h TodoHandler) Get() error {
	return nil
}
