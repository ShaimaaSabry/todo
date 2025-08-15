package addtask

import (
	"github.com/ShaimaaSabry/todo/internal/domain/model"
)

type todoRepository interface {
	SaveTask(todo model.Todo)
}

type Interactor struct {
	todoRepository todoRepository
}

func NewInteractor(todoRepository todoRepository) *Interactor {
	return &Interactor{
		todoRepository: todoRepository,
	}
}

type AddTodoCommand struct {
	Name string
}

func (c *Interactor) Execute(command AddTodoCommand) (model.Todo, error) {
	todo, err := model.NewTodo(command.Name)
	if err != nil {
		return model.Todo{}, err
	}

	c.todoRepository.SaveTask(*todo)

	return *todo, nil
}
