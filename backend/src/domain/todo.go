package domain

type Todo struct {
	ID         int    `json:"id" gorm:"primaryKey;autoIncrement"`
	Name       string `json:"name"`
	Done       bool   `json:"done"`
	TodoListID int    `json:"todo_list_id"`
}

type TodoRepository interface {
	NewTodo(todoListId int, do Todo) error
	ChangeTodo(do Todo) error
	LoadTodo(todoListId, id int) (*Todo, error)
	LoadAllTodo() (*[]Todo, error)
}
