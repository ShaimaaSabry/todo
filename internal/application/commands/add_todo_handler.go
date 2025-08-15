package commands

import (
	"github.com/ShaimaaSabry/todo/internal/domain/contracts"
	"github.com/ShaimaaSabry/todo/internal/domain/model"
	"github.com/ShaimaaSabry/todo/internal/infrastructure/repository"
)

type AddTodoCommand struct {
	Name string
}

func AddTaskHandler(todoRepository contracts.TodoRepository, command AddTodoCommand) (model.Todo, error) {
	todo, err := model.NewTodo(command.Name)
	if err != nil {
		return model.Todo{}, err
	}

	todoRepository.SaveTask(*todo)

	return *todo, nil
}

func newx() contracts.TodoRepository {
	return &repository.TodoInMemoryRepository{}
}
