package queries

import (
	"github.com/ShaimaaSabry/todo/internal/domain/contracts"
	"github.com/ShaimaaSabry/todo/internal/domain/model"
	"github.com/ShaimaaSabry/todo/internal/infrastructure/repository"
)

func GetTasksHandler(todoRepository contracts.TodoRepository) []model.Todo {
	return todoRepository.GetTasks(true)
}

func new() contracts.TodoRepository {
	return &repository.TodoInMemoryRepository{}
}
