package api

import (
	"github.com/ShaimaaSabry/todo/internal/application/commands/addtask"
)

type createTodoRequest struct {
	Name string
}

func (r *createTodoRequest) convertToCommand() addtask.AddTodoCommand {
	return addtask.AddTodoCommand{
		Name: r.Name,
	}
}
