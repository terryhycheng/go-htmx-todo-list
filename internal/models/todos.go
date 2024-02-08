package models

import (
	"github.com/google/uuid"
	h "github.com/terryhycheng/go-todo-list/internal/helpers"
)

type Todos []*Todo

func NewTodos() *Todos {
	return &Todos{}
}

func (t *Todos) GetAll() *Todos {
	return t
}

func (t *Todos) GetByID(id uuid.UUID) (*Todo, int, error) {
	index := t.IndexOf(id)

	if index == -1 {
		return nil, index, h.Error(h.ErrTodoNotFound)
	}

	return (*t)[index], index, nil
}

func (t *Todos) Add(title string, description string, status string) (*Todo, error) {
	if title == "" {
		return nil, h.Error(h.ErrTitleRequired)
	}

	if description == "" {
		return nil, h.Error(h.ErrDescriptionRequired)
	}

	if status == "" {
		return nil, h.Error(h.ErrStatusRequired)
	}

	newTodo := NewTodo(title, description, status)
	*t = append(*t, newTodo)

	return newTodo, nil
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