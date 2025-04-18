package domain

type TodoList struct {
	ID    int    `json:"id" gorm:"primaryKey;autoIncrement"`
	Todos []Todo `json:"todos" gorm:"foreignKey:TodoListID"`
	Name  string `json:"name"`
}

type TodoListRepository interface {
	NewTodoList(list TodoList) error
	ChangeTodoList(list TodoList) error
	LoadTodoList(id int) (*TodoList, error)
	LoadAllTodoLists() (*[]TodoList, error)
}
