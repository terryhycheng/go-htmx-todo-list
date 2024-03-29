package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/terryhycheng/go-todo-list/internal/helpers"
	"github.com/terryhycheng/go-todo-list/internal/models"
	"github.com/terryhycheng/go-todo-list/web/pages"
)

//go:generate mockery --name PageControllers
type PageControllers interface {
	HomePageController(c echo.Context) error
}

type pageControllers struct {
	models models.TodoRepository
}

func NewHomepageControllers(m models.TodoRepository) PageControllers {
	return &pageControllers{
		models: m,
	}
}

func (pc *pageControllers) HomePageController(c echo.Context) error {
	todos := pc.models.GetTodos()
	return helpers.Render(c, http.StatusOK, pages.Homepage(todos))
}
