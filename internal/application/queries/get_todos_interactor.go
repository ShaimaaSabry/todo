package queries

import (
	"github.com/ShaimaaSabry/todo/internal/domain/contracts"
	"github.com/ShaimaaSabry/todo/internal/domain/model"
)

func GetTasksHandler(todoRepository contracts.TodoRepository) []model.Todo {
	return todoRepository.GetTasks(true)
}
