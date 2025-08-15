package model

import "errors"

var (
	ErrInvalidTaskName = errors.New("task name must not be empty")
)

type Todo struct {
	Id       int
	Task     string
	Complete bool
}

func NewTodo(name string) (*Todo, error) {
	if name == "" {
		return nil, ErrInvalidTaskName
	}

	return &Todo{
		Id:       0,
		Task:     name,
		Complete: false,
	}, nil
}

func (todo *Todo) Check() {
	todo.Complete = true
}

func (todo *Todo) Uncheck() {
	todo.Complete = false
}
