package controllers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/terryhycheng/go-todo-list/internal/helpers"
	"github.com/terryhycheng/go-todo-list/internal/models"
	"github.com/terryhycheng/go-todo-list/web/partials"
)

//go:generate mockery --name TodoControllers
type TodoControllers interface {
	AddTodoController(c echo.Context) error
	ChangeTodoStatusController(c echo.Context) error
	DeleteTodoController(c echo.Context) error
}

type todoControllers struct {
	models models.TodoRepository
}

func NewTodoControllers(m models.TodoRepository) TodoControllers {
	return &todoControllers{
		models: m,
	}
}

func (tc *todoControllers) AddTodoController(c echo.Context) error {
	title := c.FormValue("title")
	description := c.FormValue("description")
	priority := c.FormValue("priority")

	newTodoGrom := &models.TodoGorm{
		Title:    title,
		Content:  description,
		Priority: priority,
		IsDone:   false,
	}

	todo := tc.models.AddTodo(*newTodoGrom)

	return helpers.Render(c, http.StatusCreated, partials.Card(todo))
}

func (tc *todoControllers) ChangeTodoStatusController(c echo.Context) error {
	id := c.Param("id")

	i, err := strconv.ParseInt(id, 10, 64)

	if err != nil {
		return c.String(http.StatusBadRequest, "Invalid ID")
	}

	todo := tc.models.ChangeStatus(i)

	return helpers.Render(c, http.StatusOK, partials.Card(todo))
}

func (tc *todoControllers) DeleteTodoController(c echo.Context) error {
	id := c.Param("id")

	i, err := strconv.ParseInt(id, 10, 64)

	if err != nil {
		return c.String(http.StatusBadRequest, "Invalid ID")
	}

	tc.models.DeleteTodoById(i)

	todos := tc.models.GetTodos()

	return helpers.Render(c, http.StatusOK, partials.CardList(todos))
}
