package controllers

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/terryhycheng/go-todo-list/internal/helpers"
	"github.com/terryhycheng/go-todo-list/internal/models"
	"github.com/terryhycheng/go-todo-list/web/pages"
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