package api

import "github.com/ShaimaaSabry/todo/internal/application/commands"

type createTodoRequest struct {
	Name string
}

func (r *createTodoRequest) convertCreateTodoRequestToCommand() commands.AddTodoCommand {
	return commands.AddTodoCommand{
		Name: r.Name,
	}
}
