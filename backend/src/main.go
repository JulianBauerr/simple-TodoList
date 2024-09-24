package main

import (
	"Checklist/backend/src/infrastructure/database"
	"Checklist/backend/src/infrastructure/router"
	"Checklist/backend/src/interface/controllers"
	"Checklist/backend/src/interface/repository"
	"Checklist/backend/src/usecases"
	"fmt"
	"log"
	"net/http"
)

var (
	httpRouter     router.Router = router.NewMuxRouter()
	dbHandler, err               = database.NewDBHandler("postgres://julian:password@localhost:5432/postgres?sslmode=disable")
)

func getTodoController() controllers.TodoController {
	todoRepository := repository.NewTodoRepo(dbHandler)
	todoInteractor := usecases.NewTodoInteractor(todoRepository)
	todoController := controllers.NewTodoController(todoInteractor)
	return *todoController
}

func getTodoListController() controllers.TodoListController {
	todoListListRepository := repository.NewTodoListRepo(dbHandler)
	todoListInteractor := usecases.NewTodoListInteractor(todoListListRepository)
	todoListController := controllers.NewTodoListController(todoListInteractor)
	return *todoListController
}

func main() {
	if err != nil {
		log.Fatal("Database Error:", err)
	}

	httpRouter.GET("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("The Application is up and running...")
	})

	todoController := getTodoController()
	httpRouter.POST("/todo/add", todoController.AddTodo)
	httpRouter.POST("/todo/load", todoController.LoadTodo)
	httpRouter.POST("/todo/change", todoController.ChangeTodo)
	httpRouter.POST("/todo/loadAll", todoController.LoadAllTodos)

	todoListController := getTodoListController()
	httpRouter.POST("/todoList/add", todoListController.AddTodoList)
	httpRouter.POST("/todoList/change", todoListController.ChangeTodoList)
	httpRouter.POST("/todoList/load", todoListController.LoadTodoList)
	httpRouter.POST("/todoList/loadAll", todoListController.LoadAllTodoList)
	httpRouter.SERVE(":8000")
}
