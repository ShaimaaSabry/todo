package api

import "github.com/ShaimaaSabry/todo/internal/domain/model"

type todoDto struct {
	Id       int
	Task     string
	Complete bool
}

func convertTodoToDto(todo model.Todo) todoDto {
	return todoDto{
		Id:       todo.Id(),
		Task:     todo.Task(),
		Complete: todo.Complete(),
	}
}
