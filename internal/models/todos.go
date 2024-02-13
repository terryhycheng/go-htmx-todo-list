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

		todoModel := NewTodo()

		json.Unmarshal([]byte(todo), todoModel)

		*t = append(*t, todoModel)
	}

	return nil
}

func (t *Todos) GetByID(id uuid.UUID) (*Todo, int, error) {
	index := t.IndexOf(id)

	if index == -1 {
		return nil, index, errors.New("Todo not found")
	}

	return (*t)[index], index, nil
}

func (t *Todos) Add(todo *Todo) error {
	*t = append(*t, todo)

	_, pushErr := client.LPush(ctx, "todos", todo.Id.String()).Result()
	if pushErr != nil {
		return errors.New("Failed to push todo to Redis todos list")
	}

	return nil
}

func (t *Todos) Delete(id uuid.UUID) error {
	todo, index, err := t.GetByID(id)
	*t = append((*t)[:index], (*t)[index+1:]...)

	if err != nil {
		return err
	}

	_, delErr := client.LRem(ctx, "todos", 0, todo.Id.String()).Result()

	if delErr != nil {
		return errors.New("Failed to remove todo from Redis todos list")
	}
	

	return nil
}

func (t *Todos) IndexOf(id uuid.UUID) int {
	for i, todo := range *t {
		if todo.Id == id {
			return i
		}
	}

	return -1
}