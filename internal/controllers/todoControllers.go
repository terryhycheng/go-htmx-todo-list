package controllers

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/terryhycheng/go-todo-list/internal/helpers"
	"github.com/terryhycheng/go-todo-list/internal/models"
	"github.com/terryhycheng/go-todo-list/web/partials"
)

func AddTodoController(c echo.Context) error {
	title := c.FormValue("title")
	description := c.FormValue("description")
	priority := c.FormValue("priority")

	newTodo := models.NewTodo()

	createNewTodoErr := newTodo.Add(title, description, priority)
	if createNewTodoErr != nil {
		return c.String(http.StatusBadRequest, "Failed to create new todo: "+createNewTodoErr.Error())
	}

	return helpers.Render(c, http.StatusCreated, partials.Card(newTodo))
}

func ChangeTodoStatusController(c echo.Context) error {
	id := c.Param("id")

	ParsedUuid, idErr := uuid.Parse(id)
	if idErr != nil {
		return c.String(http.StatusBadRequest, "Invalid UUID")
	}

	todos := models.NewTodos()

	getAllErr := todos.GetAll()
	if getAllErr != nil {
		return c.String(http.StatusInternalServerError, "Failed to get todos from Redis: "+getAllErr.Error())
	}

	todo, _, getTodoErr := todos.GetByID(ParsedUuid)
	if getTodoErr != nil {
		return c.String(http.StatusNotFound, "Failed to get todo: "+getTodoErr.Error())
	}

	changeStatusErr := todo.ChangeStatus()
	if changeStatusErr != nil {
		return c.String(http.StatusInternalServerError, "Failed to change status: "+changeStatusErr.Error())
	}

	return helpers.Render(c, http.StatusOK, partials.Card(todo))
}

func DeleteTodoController(c echo.Context) error {
	id := c.Param("id")

	ParsedUuid, idErr := uuid.Parse(id)
	if idErr != nil {
		return c.String(http.StatusBadRequest, "Invalid UUID")
	}

	todos := models.NewTodos()

	getAllErr := todos.GetAll()
	if getAllErr != nil {
		return c.String(http.StatusInternalServerError, "Failed to get todos from Redis: "+getAllErr.Error())
	}

	deleteErr := todos.Delete(ParsedUuid)
	if deleteErr != nil {
		return c.String(http.StatusInternalServerError, "Failed to delete todo: "+deleteErr.Error())
	}

	return helpers.Render(c, http.StatusOK, partials.CardList(todos))
}
