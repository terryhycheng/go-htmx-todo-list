package controllers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/terryhycheng/go-todo-list/internal/helpers"
	"github.com/terryhycheng/go-todo-list/internal/models"
	"github.com/terryhycheng/go-todo-list/web/partials"
)

func AddTodoController(c echo.Context) error {
	title := c.FormValue("title")
	description := c.FormValue("description")
	priority := c.FormValue("priority")

	newTodoGrom := &models.TodoGorm{
		Title:    title,
		Content:  description,
		Priority: priority,
		IsDone:   false,
	}

	newTodoGrom.AddTodo()

	return helpers.Render(c, http.StatusCreated, partials.Card(newTodoGrom))
}

func ChangeTodoStatusController(c echo.Context) error {
	id := c.Param("id")

	i, err := strconv.ParseInt(id, 10, 64)

	if err != nil {
		return c.String(http.StatusBadRequest, "Invalid ID")
	}

	todo := models.ChangeStatus(i)

	return helpers.Render(c, http.StatusOK, partials.Card(todo))
}

func DeleteTodoController(c echo.Context) error {
	id := c.Param("id")

	i, err := strconv.ParseInt(id, 10, 64)

	if err != nil {
		return c.String(http.StatusBadRequest, "Invalid ID")
	}

	models.DeleteTodoById(i)

	todos := models.GetTodos()

	return helpers.Render(c, http.StatusOK, partials.CardList(todos))
}
