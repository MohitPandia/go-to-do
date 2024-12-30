package todo

import "go-to-do/services/todosvc"

type todoHandler struct {
	todoSvc todosvc.Interface
}

func Handler(todoSvc todosvc.Interface) *todoHandler {
	return &todoHandler{
		todoSvc: todoSvc,
	}
}
