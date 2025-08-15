package model

import "errors"

var (
	ErrInvalidTaskName = errors.New("task name must not be empty")
)

type Todo struct {
	id       int
	task     string
	complete bool
}

func NewTodo(name string) (*Todo, error) {
	if name == "" {
		return nil, ErrInvalidTaskName
	}

	return &Todo{
		id:       0,
		task:     name,
		complete: false,
	}, nil
}

func Of(id int, name string, complete bool) Todo {
	return Todo{
		id:       id,
		task:     name,
		complete: complete,
	}
}

func (todo *Todo) Id() int {
	return todo.id
}

func (todo *Todo) Task() string {
	return todo.task
}

func (todo *Todo) Complete() bool {
	return todo.complete
}

func (todo *Todo) Check() {
	todo.complete = true
}

func (todo *Todo) Uncheck() {
	todo.complete = false
}
