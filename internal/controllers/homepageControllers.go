package controllers

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/terryhycheng/go-todo-list/internal/helpers"
	"github.com/terryhycheng/go-todo-list/internal/models"
	"github.com/terryhycheng/go-todo-list/web/pages"
	"github.com/terryhycheng/go-todo-list/web/partials"
)

var ctx = context.Background()
var client = helpers.RedisClient()


func HomePageController(c echo.Context) error {
	todos := models.NewTodos()

	getAllErr := todos.GetAll()

	if getAllErr != nil {
		return c.String(http.StatusInternalServerError, "Failed to get todos from Redis: " + getAllErr.Error())
	}

	return helpers.Render(c, http.StatusOK, pages.Homepage(todos))
}

func AddTodoController(c echo.Context) error {
	title := c.FormValue("title")
	description := c.FormValue("description")
	priority := c.FormValue("priority")

	newTodo := models.NewTodo()

	createNewTodoErr := newTodo.Add(title, description, priority)
	if createNewTodoErr != nil {
		return c.String(http.StatusBadRequest, "Failed to create new todo: " + createNewTodoErr.Error())
	}

	return helpers.Render(c, http.StatusCreated, partials.Card(newTodo))
}