package database

import (
	"Checklist/backend/src/domain"
	"database/sql"
	"errors"
	"fmt"
	"gorm.io/driver/postgres"
	_ "gorm.io/driver/postgres"
	"gorm.io/gorm"
	_ "gorm.io/gorm"
	"log"
	"time"
)

type DBHandler struct {
	SqlDB  *sql.DB
	GormDB *gorm.DB
}

func NewDBHandler(connectString string) (DBHandler, error) {
	sqlDB, err := sql.Open("pgx", connectString)
	if err != nil {
		log.Fatal(err)
		return DBHandler{}, err
	}

	err = sqlDB.Ping()
	if err != nil {
		return DBHandler{}, err
	}

	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: sqlDB,
	}), &gorm.Config{})
	if err != nil {
		return DBHandler{}, err
	}

	err = gormDB.AutoMigrate(&domain.TodoList{}, &domain.Todo{})
	if err != nil {
		return DBHandler{}, err
	}

	return DBHandler{sqlDB, gormDB}, nil
}

func (dbHandler DBHandler) SaveTodo(todoListId int, do domain.Todo) error {
	var todoList domain.TodoList
	if err := dbHandler.GormDB.First(&todoList, todoListId).Error; err != nil {
		return fmt.Errorf("todo list with ID %d does not exist", todoListId)
	}
	do.TodoListID = todoListId
	if err := dbHandler.GormDB.Save(&do).Error; err != nil {
		return err
	}
	fmt.Println(fmt.Sprintf("%s: Added a todo\t\t", time.Now().Format(time.TimeOnly)), do)
	return nil
}

func (dbHandler DBHandler) LoadTodo(todoListId, id int) (*domain.Todo, error) {
	todoList := domain.TodoList{}
	dbHandler.GormDB.Preload("Todos").Find(&todoList, "ID = ?", todoListId)
	for i := 0; i < len(todoList.Todos); i++ {
		if todoList.Todos[i].ID == id {
			do := todoList.Todos[i]
			fmt.Println(fmt.Sprintf("%s: Loaded Todo(%d, %d)\t", time.Now().Format(time.TimeOnly), todoListId, id), do)
			return &do, nil
		}
	}
	return nil, errors.New(fmt.Sprintf("Todo with todoListid: %d and id: %d not found", todoListId, id))
}

func (dbHandler DBHandler) ChangeTodo(do domain.Todo) error {
	var todo domain.Todo
	if err := dbHandler.GormDB.Find(&todo, do.ID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("todo with ID %d not found", do.ID)
		}
		return fmt.Errorf("failed to find todo: %v", err)
	}

	if err := dbHandler.GormDB.Save(&do).Error; err != nil {
		return fmt.Errorf("failed to update todo: %v", err)
	}

	fmt.Println(fmt.Sprintf("%s: Upadted Todo\t", time.Now().Format(time.TimeOnly)), do)
	return nil
}

func (dbHandler DBHandler) LoadAllTodos() (*[]domain.Todo, error) {
	return nil, nil
}

func (dbHandler DBHandler) SaveTodoList(list domain.TodoList) error {
	if err := dbHandler.GormDB.Create(&list).Error; err != nil {
		return fmt.Errorf("failed to insert data: %v", err)
	}
	return nil
}

func (dbHandler DBHandler) ChangeTodoList(list domain.TodoList) error {
	var todoList domain.TodoList
	if err := dbHandler.GormDB.Find(&todoList, list.ID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("todoList with ID %d not found", list.ID)
		}
		return fmt.Errorf("failed to find todo: %v", err)
	}

	if err := dbHandler.GormDB.Create(&list).Error; err != nil {
		return fmt.Errorf("failed to insert data: %v", err)
	}

	fmt.Println(fmt.Sprintf("%s: Upadted TodoList\t", time.Now().Format(time.TimeOnly)), list)
	return nil
}

func (dbHandler DBHandler) LoadTodoList(id int) (*domain.TodoList, error) {
	todoList := domain.TodoList{}
	var todos []domain.Todo
	if err := dbHandler.GormDB.Where("todo_list_id = ?", id).Find(&todos).Error; err != nil {
		return nil, fmt.Errorf("failed to retrieve todos for todo list ID %d: %w", id, err)
	}
	todoList = domain.TodoList{
		ID:    id,
		Todos: todos,
	}

	fmt.Println(fmt.Sprintf("%s: Loaded TodoList\t", time.Now().Format(time.TimeOnly)), todoList)
	return &todoList, nil
}

func (dbHandler DBHandler) LoadAllTodoLists() (*[]domain.TodoList, error) {
	var todoLists []domain.TodoList
	if err := dbHandler.GormDB.Find(&todoLists).Error; err != nil {
		return nil, fmt.Errorf("failed to retrieve todo lists: %w", err)
	}

	for i := range todoLists {
		var todos []domain.Todo
		if err := dbHandler.GormDB.Where("todo_list_id = ?", todoLists[i].ID).Find(&todos).Error; err != nil {
			return nil, fmt.Errorf("failed to retrieve todos for todo list ID %d: %w", todoLists[i].ID, err)
		}
		todoLists[i].Todos = todos
	}
	fmt.Println(fmt.Sprintf("%s: All TodoLists loaded\t", time.Now().Format(time.TimeOnly)))
	return &todoLists, nil
}

func Initialize(dbHandler DBHandler) error {
	exampleData := []domain.TodoList{
		{
			Todos: []domain.Todo{
				{Name: "Buy groceries", Done: false},
				{Name: "Write Go code", Done: true},
			},
		},
		{
			Todos: []domain.Todo{
				{Name: "Read a book", Done: false},
				{Name: "Exercise", Done: true},
			},
		},
	}

	for i := range exampleData {
		if err := dbHandler.GormDB.Create(&exampleData[i]).Error; err != nil {
			return fmt.Errorf("failed to insert example data: %v", err)
		}
	}

	log.Println("Database initialized with example data")
	return nil
}
