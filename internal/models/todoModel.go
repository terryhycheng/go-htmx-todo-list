package models

import (
	"gorm.io/gorm"
)

//go:generate mockery --name TodoRepository
type TodoRepository interface {
	AddTodo(todo TodoGorm) *TodoGorm
	GetTodos() []TodoGorm
	GetTodoById(id int64) *TodoGorm
	DeleteTodoById(id int64) TodoGorm
	ChangeStatus(id int64) *TodoGorm
}

type todoRepository struct {
	db *gorm.DB
}

type TodoGorm struct {
	gorm.Model
	Title    string `json:"title"`
	Content  string `json:"content"`
	Priority string `json:"priority"`
	IsDone   bool   `json:"is_done"`
}

func NewTodoRepository(db *gorm.DB) TodoRepository {
	return &todoRepository{
		db: db,
	}
}

func (t *todoRepository) AddTodo(todo TodoGorm) *TodoGorm {
	t.db.Create(&todo)
	return &todo
}

func (t *todoRepository) GetTodos() []TodoGorm {
	var todos []TodoGorm
	t.db.Order("created_at DESC").Find(&todos)
	return todos
}

func (t *todoRepository) GetTodoById(id int64) *TodoGorm {
	var todo TodoGorm
	t.db.First(&todo, id)
	return &todo
}

func (t *todoRepository) DeleteTodoById(id int64) TodoGorm {
	var todo TodoGorm
	t.db.Delete(&todo, id)
	return todo
}

func (t *todoRepository) ChangeStatus(id int64) *TodoGorm {
	todo := t.GetTodoById(id)
	todo.IsDone = !todo.IsDone
	t.db.Save(&todo)
	return todo
}
