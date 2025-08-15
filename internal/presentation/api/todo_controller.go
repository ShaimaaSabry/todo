package api

import (
    "log"
	"encoding/json"
	"github.com/ShaimaaSabry/todo/internal/application/commands"
	"github.com/ShaimaaSabry/todo/internal/application/queries"
	"github.com/ShaimaaSabry/todo/internal/domain/contracts"
	"net/http"
	"strconv"
)

type TodoController struct {
	todoRepository contracts.TodoRepository
}

func NewTodoController(todoRepository contracts.TodoRepository) *TodoController {
	return &TodoController{todoRepository: todoRepository}
}

func (h *TodoController) AddTodoHandler(w http.ResponseWriter, request *http.Request) {
	var requestBody createTodoRequest
	err := json.NewDecoder(request.Body).Decode(&requestBody)
	if err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	todo, err := commands.Execute(
		h.todoRepository,
		requestBody.convertToCommand(),
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	response := convertTodoToDto(todo)

	responseBytes, err := json.Marshal(response)
    if err != nil {
        http.Error(w, "failed to encode response", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    if _, err := w.Write(responseBytes); err != nil {
        // At this point the status code is already sent, so we can't change it
        log.Printf("failed to write response: %v", err)
        return
    }
}

func (h *TodoController) GetTodosHandler(w http.ResponseWriter, request *http.Request) {
	var todos = queries.GetTasksHandler(h.todoRepository)

	var response []todoDto
	for _, t := range todos {
		response = append(
			response,
			convertTodoToDto(t),
		)
	}

	responseBytes, err := json.Marshal(response)
    if err != nil {
        http.Error(w, "failed to encode response", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    if _, err := w.Write(responseBytes); err != nil {
        // At this point the status code is already sent, so we can't change it
        log.Printf("failed to write response: %v", err)
        return
    }
}

func (h *TodoController) CheckTodoHandler(w http.ResponseWriter, request *http.Request, taskIdStr string) {
	taskId, err := strconv.Atoi(taskIdStr)
	if err != nil {
		http.Error(w, "invalid task ID", http.StatusBadRequest)
		return
	}

	commands.CheckTaskHandler(h.todoRepository, taskId)
}

func (h *TodoController) UncheckTodoHandler(w http.ResponseWriter, request *http.Request, taskIdStr string) {
	taskId, err := strconv.Atoi(taskIdStr)
	if err != nil {
		http.Error(w, "invalid task ID", http.StatusBadRequest)
		return
	}

	commands.UncheckTaskHandler(h.todoRepository, taskId)
}
