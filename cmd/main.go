package main

import (
	"github.com/ShaimaaSabry/todo/internal/infrastructure/repository"
	"github.com/ShaimaaSabry/todo/internal/presentation/api"
	"log"
	"net/http"
	"strings"
)

func main() {
	// dependency injection
	todoRepository := &repository.TodoInMemoryRepository{}
	todoController := api.NewTodoController(todoRepository)

	// routing
	http.HandleFunc(
		"/todos",
		func(w http.ResponseWriter, request *http.Request) {
			switch request.Method {
			case http.MethodGet:

				todoController.GetTodosHandler(w, request)
			case http.MethodPost:
				todoController.AddTodoHandler(w, request)
			default:
				http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			}
		},
	)

	http.HandleFunc(
		"/todos/",
		func(w http.ResponseWriter, request *http.Request) {
			if request.Method != http.MethodPut {
				http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
				return
			}

			path := strings.TrimPrefix(request.URL.Path, "/todos/")
			parts := strings.Split(path, "/")

			if len(parts) != 2 {
				http.NotFound(w, request)
				return
			}

			taskID := parts[0]
			action := parts[1]

			switch action {
			case "check":

				todoController.CheckTodoHandler(w, request, taskID)
			case "uncheck":
				todoController.UncheckTodoHandler(w, request, taskID)
			default:
				http.NotFound(w, request)
			}
		},
	)

	log.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
