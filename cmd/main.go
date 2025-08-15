package main

import (
	"github.com/ShaimaaSabry/todo/internal/application/commands/addtask"
	"github.com/ShaimaaSabry/todo/internal/application/commands/checktask"
	"github.com/ShaimaaSabry/todo/internal/application/queries/gettasks"
	"github.com/ShaimaaSabry/todo/internal/infrastructure/repositories"
	"github.com/ShaimaaSabry/todo/internal/presentation/api"
	httpSwagger "github.com/swaggo/http-swagger/v2"
	"log"
	"net/http"
	"strings"

	_ "github.com/ShaimaaSabry/todo/docs"
)

// @title        To-Do API
// @version      1.0
// @description  Simple APIs for a To-Do app.
func main() {
	todoController := dependencyInjection()

	setupRoutes(todoController)
	setupSwagger()
	startServer()
}

func dependencyInjection() *api.TodoController {
	todoRepository := &repositories.TodoInMemoryRepository{}

	addTasksInteractor := addtask.NewInteractor(todoRepository)
	getTasksInteractor := gettasks.NewInteractor(todoRepository)
	checkTaskInteractor := checktask.NewInteractor(todoRepository)
	uncheckTaskInteractor := checktask.NewInteractor(todoRepository)

	todoController := api.NewTodoController(
		addTasksInteractor,
		getTasksInteractor,
		checkTaskInteractor,
		uncheckTaskInteractor,
	)

	return todoController
}

func setupRoutes(todoController *api.TodoController) {
	http.HandleFunc(
		"/v1/todos",
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
		"/v1/todos/",
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
}

func setupSwagger() {
	http.Handle("/swagger/", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8080/swagger/doc.json"),
	))
}

func startServer() {
	log.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
