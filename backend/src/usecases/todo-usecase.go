package usecases

import (
	"Checklist/backend/src/domain"
	"log"
)

type TodoInteractor struct {
	TodoRepository domain.TodoRepository
}

func NewTodoInteractor(repository domain.TodoRepository) TodoInteractor {
	return TodoInteractor{repository}
}

func (interactor *TodoInteractor) CreateToDo(todoListId int, do domain.Todo) error {
	err := interactor.TodoRepository.NewTodo(todoListId, do)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (interactor *TodoInteractor) LoadToDo(todoListId, id int) (*domain.Todo, error) {
	do, err := interactor.TodoRepository.LoadTodo(todoListId, id)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return do, nil
}

func (interactor *TodoInteractor) ChangeToDo(do domain.Todo) error {
	err := interactor.TodoRepository.ChangeTodo(do)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (interactor *TodoInteractor) LoadAllTodo() (*[]domain.Todo, error) {
	todos, err := interactor.TodoRepository.LoadAllTodo()
	if err != nil {
		return nil, err
	}
	return todos, nil
}
