package unchecktask

import (
	"errors"
	"fmt"
	"github.com/ShaimaaSabry/todo/internal/domain/model"
)

type todoRepository interface {
	GetTask(taskId int) *model.Todo
}

type Interactor struct {
	todoRepository todoRepository
}

func NewInteractor(todoRepository todoRepository) *Interactor {
	return &Interactor{
		todoRepository: todoRepository,
	}
}

func (c *Interactor) Execute(taskId int) error {
	fmt.Println("Fetching task with ID:", taskId)
	var todo = c.todoRepository.GetTask(taskId)
	if todo == nil {
		return errors.New(
			fmt.Sprintf("task with ID %d not found", taskId),
		)
	}

	fmt.Println("Unchecking task with ID:", todo.Id())
	todo.Uncheck()
	return nil
}
