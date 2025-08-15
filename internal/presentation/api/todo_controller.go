package api

import (
	"encoding/json"
	"github.com/ShaimaaSabry/todo/internal/application/commands/addtask"
	"github.com/ShaimaaSabry/todo/internal/domain/model"
	"net/http"
	"strconv"
)

type addTaskInteractor interface {
	Execute(command addtask.AddTodoCommand) (model.Todo, error)
}

type getTasksInteractor interface {
	Execute() []model.Todo
}

type checkTaskInteractor interface {
	Execute(taskId int) error
}

type uncheckTaskInteractor interface {
	Execute(taskId int) error
}

type TodoController struct {
	addTaskInteractor     addTaskInteractor
	getTasksInteractor    getTasksInteractor
	checkTaskInteractor   checkTaskInteractor
	uncheckTaskInteractor uncheckTaskInteractor
}

func NewTodoController(
	addTaskInteractor addTaskInteractor,
	getTasksInteractor getTasksInteractor,
	checkTaskInteractor checkTaskInteractor,
	uncheckTaskInteractor uncheckTaskInteractor,
) *TodoController {
	return &TodoController{
		addTaskInteractor:     addTaskInteractor,
		getTasksInteractor:    getTasksInteractor,
		checkTaskInteractor:   checkTaskInteractor,
		uncheckTaskInteractor: uncheckTaskInteractor,
	}
}

// AddTodoHandler godoc
// @Tags        To-Do
// @Router      /todo [post]
// @Summary     Create a To-Do item.
// @Param       payload body createTodoRequest true "Create To-Do item payload"
// @Success     201 {object}  todoDto
func (h *TodoController) AddTodoHandler(w http.ResponseWriter, request *http.Request) {
	var requestBody createTodoRequest
	err := json.NewDecoder(request.Body).Decode(&requestBody)
	if err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	todo, err := h.addTaskInteractor.Execute(
		requestBody.convertToCommand(),
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	response := convertTodoToDto(todo)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

// GetTodosHandler godoc
// @Tags        To-Do
// @Router      /todo [get]
// @Summary     Get the To-Do list.
// @Success     200 {array}  todoDto
func (h *TodoController) GetTodosHandler(w http.ResponseWriter, request *http.Request) {
	var todos = h.getTasksInteractor.Execute()

	var response []todoDto
	for _, t := range todos {
		response = append(
			response,
			convertTodoToDto(t),
		)
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

// CheckTodoHandler godoc
// @Tags        To-Do
// @Router      /todo/:id/check [put]
// @Summary     Check a To-Do item.
// @Success     200
func (h *TodoController) CheckTodoHandler(w http.ResponseWriter, request *http.Request, taskIdStr string) {
	taskId, err := strconv.Atoi(taskIdStr)
	if err != nil {
		http.Error(w, "invalid task ID", http.StatusBadRequest)
		return
	}

	err = h.checkTaskInteractor.Execute(taskId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
}

// UncheckTodoHandler godoc
// @Tags        To-Do
// @Router      /todo/:id/uncheck [put]
// @Summary     Uncheck a To-Do item.
// @Success     200
func (h *TodoController) UncheckTodoHandler(w http.ResponseWriter, request *http.Request, taskIdStr string) {
	taskId, err := strconv.Atoi(taskIdStr)
	if err != nil {
		http.Error(w, "invalid task ID", http.StatusBadRequest)
		return
	}

	err = h.uncheckTaskInteractor.Execute(taskId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
}
