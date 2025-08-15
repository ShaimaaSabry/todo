package commands

import (
	"fmt"
	"github.com/ShaimaaSabry/todo/internal/domain/contracts"
)

func CheckTaskHandler(todoRepository contracts.TodoRepository, taskId int) {
	fmt.Println("Fetching task with ID:", taskId)

	var todo = todoRepository.GetTask(taskId)
	fmt.Println("Checking task with ID:", todo.Id())
	todo.Check()
}
