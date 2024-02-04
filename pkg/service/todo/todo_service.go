package service

import (
	domain "github.com/satoshi-tahara-st/todo_backend/pkg/domain/todo"
)

type ITodoService interface {
	Fetch()
}

func NewTodoService(tr domain.ITodoRepository) ITodoService {
	return &todoServiceImpl{
		todoRepository: tr,
	}
}

type todoServiceImpl struct {
	todoRepository domain.ITodoRepository
}

func (s *todoServiceImpl) Fetch() {

}
