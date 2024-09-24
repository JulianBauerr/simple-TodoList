package repository

import "Checklist/backend/src/domain"

type TodoRepo struct {
	handler DBHandler
}

func NewTodoRepo(handler DBHandler) TodoRepo {
	return TodoRepo{handler}
}

func (repo TodoRepo) NewTodo(todoListId int, do domain.Todo) error {
	err := repo.handler.SaveTodo(todoListId, do)
	if err != nil {
		return err
	}
	return nil
}

func (repo TodoRepo) LoadTodo(todoListId, id int) (*domain.Todo, error) {
	do, err := repo.handler.LoadTodo(todoListId, id)
	if err != nil {
		return nil, err
	}
	return do, nil
}

func (repo TodoRepo) ChangeTodo(do domain.Todo) error {
	err := repo.handler.ChangeTodo(do)
	if err != nil {
		return err
	}
	return nil
}

func (repo TodoRepo) LoadAllTodo() (*[]domain.Todo, error) {
	todos, err := repo.handler.LoadAllTodos()
	if err != nil {
		return nil, err
	}
	return todos, nil
}
