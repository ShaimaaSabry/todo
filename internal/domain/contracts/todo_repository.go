package contracts

import "github.com/ShaimaaSabry/todo/internal/domain/model"

type TodoRepository interface {
	GetTasks(includeComplete bool) []model.Todo
	GetTask(taskId int) *model.Todo
	SaveTask(todo model.Todo)
}
