package models

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/google/uuid"
)

type Todo struct {
	Id        uuid.UUID `json:"id" redis:"id"`
	Title     string    `json:"title" redis:"title"`
	Content   string    `json:"content" redis:"content"`
	Priority  string    `json:"priority" redis:"priority"`
	IsDone    bool      `json:"isDone" redis:"isDone"`
	CreatedAt time.Time `json:"createdAt" redis:"createdAt"`
}

func NewTodo() *Todo {
	return &Todo{
		Id:        uuid.New(),
		IsDone:    false,
		CreatedAt: time.Now(),
	}
}

func (t *Todo) Add(title string, description string, priority string) error {
	if title == "" || description == "" || priority == "" {
		return errors.New("title, description and priority are required")
	}

	t.Title = title
	t.Content = description
	t.Priority = priority

	addErr := NewTodos().Add(t)
	if addErr != nil {
		return addErr
	}

	json, jsonMarshalErr := json.Marshal(t)
	if jsonMarshalErr != nil {
		return errors.New("failed to marshal todo")
	}

	_, setJsonErr := client.JSONSet(ctx, "todo:"+t.Id.String(), ".", json).Result()
	if setJsonErr != nil {
		return errors.New("failed to set todo in Redis")
	}

	return nil
}

func (t *Todo) ChangeStatus() error {
	t.IsDone = !t.IsDone

	_, setJsonErr := client.JSONSet(ctx, "todo:"+t.Id.String(), ".isDone", t.IsDone).Result()
	if setJsonErr != nil {
		return errors.New("Failed to set todo in Redis")
	}

	return nil
}

func (t *Todo) Update(title string, discription string, isDone bool) error {
	t.Title = title
	t.Content = discription
	t.IsDone = isDone

	return nil
}
