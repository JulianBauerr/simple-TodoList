package repository

import "Checklist/backend/src/domain"

type TodoListRepo struct {
	handler DBHandler
}

func NewTodoListRepo(handler DBHandler) TodoListRepo {
	return TodoListRepo{handler}
}

func (repo TodoListRepo) NewTodoList(list domain.TodoList) error {
	err := repo.handler.SaveTodoList(list)
	if err != nil {
		return err
	}
	return nil
}

func (repo TodoListRepo) ChangeTodoList(list domain.TodoList) error {
	err := repo.handler.ChangeTodoList(list)
	if err != nil {
		return err
	}
	return nil
}

func (repo TodoListRepo) LoadTodoList(id int) (*domain.TodoList, error) {
	list, err := repo.handler.LoadTodoList(id)
	if err != nil {
		return nil, err
	}
	return list, nil
}

func (repo TodoListRepo) LoadAllTodoLists() (*[]domain.TodoList, error) {
	lists, err := repo.handler.LoadAllTodoLists()
	if err != nil {
		return nil, err
	}
	return lists, nil
}
