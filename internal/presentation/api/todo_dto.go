package api

import "github.com/ShaimaaSabry/todo/internal/domain/model"

type todoDto struct {
	Id       int    `json:"id"`
	Task     string `json:"task"`
	Complete bool   `json:"complete"`
}

func convertTodoToDto(todo model.Todo) todoDto {
	return todoDto{
		Id:       todo.Id(),
		Task:     todo.Task(),
		Complete: todo.Complete(),
	}
}
