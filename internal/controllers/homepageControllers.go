package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/terryhycheng/go-todo-list/internal/helpers"
	"github.com/terryhycheng/go-todo-list/internal/models"
	"github.com/terryhycheng/go-todo-list/web/pages"
)

func HomePageController(c echo.Context) error {
	todos := models.GetTodos()
	return helpers.Render(c, http.StatusOK, pages.Homepage(todos))
}
