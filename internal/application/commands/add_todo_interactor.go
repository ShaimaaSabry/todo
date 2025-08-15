package commands

import (
	"github.com/ShaimaaSabry/todo/internal/domain/model"
)

type todoRepository interface {
	SaveTask(todo model.Todo)
}

type AddTodoCommand struct {
	Name string
}

func Execute(todoRepository todoRepository, command AddTodoCommand) (model.Todo, error) {
	todo, err := model.NewTodo(command.Name)
	if err != nil {
		return model.Todo{}, err
	}

	todoRepository.SaveTask(*todo)

	return *todo, nil
}
