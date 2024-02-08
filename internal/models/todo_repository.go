package models

import "github.com/google/uuid"

type TodoRepository interface {
	GetAll() *[]*Todo
	GetByID(id uuid.UUID) (*Todo, error)
	Add(title string, description string, status string) (*Todo, error)
	Update(id uuid.UUID, title string, description string, isDone bool) (*Todo, error)
	Delete(id uuid.UUID) (*Todo, error)
	IndexOf(id uuid.UUID) int
}