package repository

import (
	"fmt"
	"github.com/ShaimaaSabry/todo/internal/domain/model"
)

var db = []model.Todo{
	model.Of(1, "Play piano", false),
	model.Of(2, "Cook", false),
	model.Of(3, "Yoga", false),
}

type TodoInMemoryRepository struct {
}

func (r *TodoInMemoryRepository) GetTasks(includeComplete bool) []model.Todo {
	return db
}

func (r *TodoInMemoryRepository) GetTask(taskId int) *model.Todo {
	for i := range db {
		if db[i].Id() == taskId {
			fmt.Println("Returning task with ID:", db[i].Id())
			return &db[i]
		}
	}
	return nil
}

func (r *TodoInMemoryRepository) SaveTask(todo model.Todo) {
	db = append(db, todo)
}
