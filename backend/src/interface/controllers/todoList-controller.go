package controllers

import (
	"Checklist/backend/src/domain"
	"Checklist/backend/src/usecases"
	"encoding/json"
	"net/http"
)

type TodoListController struct {
	todoListInteractor usecases.TodoListInteractor
}

func NewTodoListController(todoListInteractor usecases.TodoListInteractor) *TodoListController {
	return &TodoListController{todoListInteractor}
}

func (controller TodoListController) AddTodoList(w http.ResponseWriter, r *http.Request) {
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

	var todoList domain.TodoList
	err := json.NewDecoder(r.Body).Decode(&todoList)
	if err != nil {
		http.Error(w, "Error parsing request Body", http.StatusBadRequest)
		return
	}

	err = controller.todoListInteractor.CreateTodoList(todoList)
	if err != nil {
		http.Error(w, "Error saving to Database", http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(map[string]string{"message": "TodoList added successfully"})
	if err != nil {
		return
	}
}

func (controller TodoListController) ChangeTodoList(w http.ResponseWriter, r *http.Request) {
}

func (controller TodoListController) LoadTodoList(w http.ResponseWriter, r *http.Request) {
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
		Id int `json:"id"`
	}

	var idReq idRequest
	err := json.NewDecoder(r.Body).Decode(&idReq)
	if err != nil {
		http.Error(w, "Error parsing request Body", http.StatusBadRequest)
		return
	}

	list, err := controller.todoListInteractor.LoadTodoList(idReq.Id)
	if err != nil {
		http.Error(w, "Error saving to Database", http.StatusInternalServerError)
	}

	js, err := json.Marshal(list)
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

func (controller TodoListController) LoadAllTodoList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	lists, err := controller.todoListInteractor.LoadAllTodoLists()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	js, err := json.Marshal(*lists)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	_, err = w.Write(js)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
