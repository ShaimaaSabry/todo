package commands

import (
	"fmt"
	"github.com/ShaimaaSabry/todo/internal/domain/contracts"
)

func UncheckTaskHandler(todoRepository contracts.TodoRepository, taskId int) {
	fmt.Println("Fetching task with ID:", taskId)

	var todo = todoRepository.GetTask(taskId)
	fmt.Println("Unchecking task with ID:", todo.Id())
	todo.Uncheck()
}
