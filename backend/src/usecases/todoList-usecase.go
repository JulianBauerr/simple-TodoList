package usecases

import "Checklist/backend/src/domain"

type TodoListInteractor struct {
	TodoListRepository domain.TodoListRepository
}

func NewTodoListInteractor(repository domain.TodoListRepository) TodoListInteractor {
	return TodoListInteractor{repository}
}

func (interactor TodoListInteractor) CreateTodoList(list domain.TodoList) error {
	err := interactor.TodoListRepository.NewTodoList(list)
	if err != nil {
		return err
	}
	return nil
}

func (interactor TodoListInteractor) ChangeTodoList(list domain.TodoList) error {
	err := interactor.TodoListRepository.ChangeTodoList(list)
	if err != nil {
		return err
	}
	return nil
}

func (interactor TodoListInteractor) LoadTodoList(id int) (*domain.TodoList, error) {
	list, err := interactor.TodoListRepository.LoadTodoList(id)
	if err != nil {
		return nil, err
	}
	return list, err
}

func (interactor TodoListInteractor) LoadAllTodoLists() (*[]domain.TodoList, error) {
	lists, err := interactor.TodoListRepository.LoadAllTodoLists()
	if err != nil {
		return nil, err
	}
	return lists, nil
}
