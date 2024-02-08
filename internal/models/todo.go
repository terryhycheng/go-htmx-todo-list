package models

import (
	"time"

	"github.com/google/uuid"
)

type Todo struct {
	Id uuid.UUID `json:"id"`
	Title string `json:"title"`
	Content string `json:"content"`
	Status string `json:"status"`
	IsDone bool `json:"isDone"`
	CreatedAt time.Time `json:"createdAt"`
}

func NewTodo(title string, discription string, status string) *Todo {
	return &Todo{
		Id: uuid.New(),
		Title: title,
		Content: discription,
		Status: status,
		IsDone: false,
		CreatedAt: time.Now(),
	}
}

func (t *Todo) Update(title string, discription string, isDone bool) *Todo {
	t.Title = title
	t.Content = discription
	t.IsDone = isDone

	return t
}