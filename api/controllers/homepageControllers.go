package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/terryhycheng/go-todo-list/internal"
	"github.com/terryhycheng/go-todo-list/web/templates"
)

func HomePageController(c echo.Context) error {
	return internal.Render(c, http.StatusOK, templates.Homepage())
}