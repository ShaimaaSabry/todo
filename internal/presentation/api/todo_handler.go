package api

import (
	"encoding/json"
	"github.com/ShaimaaSabry/todo/internal/application/commands"
	"github.com/ShaimaaSabry/todo/internal/application/queries"
	"github.com/ShaimaaSabry/todo/internal/domain/contracts"
	"net/http"
	"strconv"
)

type TodoApiHandler struct {
	todoRepository contracts.TodoRepository
}

func NewTodoApiHandler(todoRepository contracts.TodoRepository) *TodoApiHandler {
	return &TodoApiHandler{todoRepository: todoRepository}
}

func (h *TodoApiHandler) AddTodoHandler(w http.ResponseWriter, request *http.Request) {
	var requestBody createTodoRequest
	err := json.NewDecoder(request.Body).Decode(&requestBody)
	if err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	todo, err := commands.AddTaskHandler(
		h.todoRepository,
		requestBody.convertCreateTodoRequestToCommand(),
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	response := convertTodoToDto(todo)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

func (h *TodoApiHandler) GetTodosHandler(w http.ResponseWriter, request *http.Request) {
	var todos = queries.GetTasksHandler(h.todoRepository)

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

func (h *TodoApiHandler) CheckTodoHandler(w http.ResponseWriter, request *http.Request, taskIdStr string) {
	taskId, err := strconv.Atoi(taskIdStr)
	if err != nil {
		http.Error(w, "invalid task ID", http.StatusBadRequest)
		return
	}

	commands.CheckTaskHandler(h.todoRepository, taskId)
}

func (h *TodoApiHandler) UncheckTodoHandler(w http.ResponseWriter, request *http.Request, taskIdStr string) {
	taskId, err := strconv.Atoi(taskIdStr)
	if err != nil {
		http.Error(w, "invalid task ID", http.StatusBadRequest)
		return
	}

	commands.UncheckTaskHandler(h.todoRepository, taskId)
}
