package models

import (
	"github.com/terryhycheng/go-todo-list/internal/helpers"
	"gorm.io/gorm"
)

var db *gorm.DB

type TodoGorm struct {
	gorm.Model
	Title    string `json:"title"`
	Content  string `json:"content"`
	Priority string `json:"priority"`
	IsDone   bool   `json:"is_done"`
}

func init() {
	helpers.Connect()
	db = helpers.GetDB()
	db.AutoMigrate(&TodoGorm{})
}

func (t *TodoGorm) AddTodo() *TodoGorm {
	db.Create(&t)
	return t
}

func GetTodos() []TodoGorm {
	var todos []TodoGorm
	db.Order("created_at DESC").Find(&todos)
	return todos
}

func GetTodoById(id int64) *TodoGorm {
	var todo TodoGorm
	db.First(&todo, id)
	return &todo
}

func DeleteTodoById(id int64) TodoGorm {
	var todo TodoGorm
	db.Delete(&todo, id)
	return todo
}

func ChangeStatus(id int64) *TodoGorm {
	todo := GetTodoById(id)
	todo.IsDone = !todo.IsDone
	db.Save(&todo)
	return todo
}
