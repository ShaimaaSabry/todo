package gettasks

import (
	"github.com/ShaimaaSabry/todo/internal/domain/model"
)

type todoRepository interface {
	GetTasks(includeComplete bool) []model.Todo
}

type Interactor struct {
	todoRepository todoRepository
}

func NewInteractor(todoRepository todoRepository) *Interactor {
	return &Interactor{
		todoRepository: todoRepository,
	}
}

func (c *Interactor) Execute() []model.Todo {
	return c.todoRepository.GetTasks(true)
}
