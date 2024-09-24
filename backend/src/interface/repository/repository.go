package repository

import "Checklist/backend/src/domain"

type DBHandler interface {
	SaveTodo(todoListId int, do domain.Todo) error
	LoadTodo(todoListId, id int) (*domain.Todo, error)
	ChangeTodo(do domain.Todo) error
	LoadAllTodos() (*[]domain.Todo, error)
	SaveTodoList(list domain.TodoList) error
	ChangeTodoList(list domain.TodoList) error
	LoadTodoList(id int) (*domain.TodoList, error)
	LoadAllTodoLists() (*[]domain.TodoList, error)
}
