package models

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/google/uuid"
	h "github.com/terryhycheng/go-todo-list/internal/helpers"
)

type Todos []*Todo
var ctx = context.Background()
var client = h.RedisClient()

func NewTodos() *Todos {
	return &Todos{}
}

func (t *Todos) GetAll() error {

	todoIds, err := client.LRange(ctx, "todos", 0, -1).Result()

	if err != nil {
		return errors.New("Failed to get todos from Redis: " + err.Error())
	}

	for _, todoId := range todoIds {
		todo, err := client.JSONGet(ctx,"todo:" + todoId, ".").Result()

		if err != nil {
			return errors.New("Failed to get todo details: " + err.Error())
		}

		todoModel := &Todo{}

		json.Unmarshal([]byte(todo), todoModel)

		addErr := t.Add(todoModel.Title, todoModel.Content, todoModel.Priority)

		if addErr != nil {
			return errors.New("Failed to add todo to the homepage list: " + addErr.Error())
		}
	}

	return nil
}

func (t *Todos) GetByID(id uuid.UUID) (*Todo, int, error) {
	index := t.IndexOf(id)

	if index == -1 {
		return nil, index, h.Error(h.ErrTodoNotFound)
	}

	return (*t)[index], index, nil
}

func (t *Todos) Add(title string, description string, status string) error {
	if title == "" {
		return h.Error(h.ErrTitleRequired)
	}

	if description == "" {
		return h.Error(h.ErrDescriptionRequired)
	}

	if status == "" {
		return h.Error(h.ErrStatusRequired)
	}

	newTodo := NewTodo()

	addErr := newTodo.Add(title, description, status)


	if addErr != nil {
		return errors.New("Failed to add todo details:" + addErr.Error())
	}

	*t = append(*t, newTodo)

	return nil
}

func (t *Todos) Update(id uuid.UUID, title string, description string, isDone bool) (*Todo, error) {
	todo, _ , err := t.GetByID(id)

	if err != nil {
		return nil, err
	}

	todo.Update(title, description, isDone)

	return todo, nil
}

func (t *Todos) Delete(id uuid.UUID) (*Todo, error) {
	todo, index, err := t.GetByID(id)

	if err != nil {
		return nil, err
	}
	
	*t = append((*t)[:index], (*t)[index+1:]...)

	return todo, nil
}

func (t *Todos) IndexOf(id uuid.UUID) int {
	for i, todo := range *t {
		if todo.Id == id {
			return i
		}
	}

	return -1
}