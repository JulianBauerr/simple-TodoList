package controllers

import (
	"Checklist/backend/src/domain"
	"Checklist/backend/src/usecases"
	"encoding/json"
	"net/http"
)

type TodoController struct {
	todoInteractor usecases.TodoInteractor
}

func NewTodoController(todoInteractor usecases.TodoInteractor) *TodoController {
	return &TodoController{todoInteractor}
}

func (controller *TodoController) AddTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	if r.Method != http.MethodPost {
		http.Error(w, "Invalid Request Method", http.StatusMethodNotAllowed)
		return
	}

	type todoRequest struct {
		TodoListId int    `json:"todoListId"`
		Name       string `json:"name"`
		Done       bool   `json:"done"`
	}

	var todoReq todoRequest
	err := json.NewDecoder(r.Body).Decode(&todoReq)
	if err != nil {
		http.Error(w, "Error parsing request Body", http.StatusBadRequest)
		return
	}

	var todo = domain.Todo{
		Name: todoReq.Name,
		Done: todoReq.Done,
	}

	err = controller.todoInteractor.CreateToDo(todoReq.TodoListId, todo)
	if err != nil {
		http.Error(w, "Error saving to Database", http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(map[string]string{"message": "Todo added successfully"})
	if err != nil {
		return
	}
}

func (controller *TodoController) LoadTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	if r.Method != http.MethodPost {
		http.Error(w, "Invalid Request Method", http.StatusMethodNotAllowed)
		return
	}

	type idRequest struct {
		TodoListId int `json:"todoListId"`
		Id         int `json:"id"`
	}

	var idReq idRequest
	err := json.NewDecoder(r.Body).Decode(&idReq)
	if err != nil {
		http.Error(w, "Error parsing body"+err.Error(), http.StatusBadRequest)
		return
	}

	todo, err := controller.todoInteractor.LoadToDo(idReq.TodoListId, idReq.Id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	js, err := json.Marshal(todo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, err = w.Write(js)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (controller *TodoController) ChangeTodo(w http.ResponseWriter, r *http.Request) {
	//todo
}

func (controller *TodoController) LoadAllTodos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	todos, err := controller.todoInteractor.LoadAllTodo()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	js, err := json.Marshal(*todos)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(js)
}
