package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/terryhycheng/go-todo-list/internal/controllers"
)

func TodoSubRoutes(e *echo.Group, c controllers.TodoControllers) {
	e.POST("", c.AddTodoController)
	// e.PUT("/:id", controllers.ChangeTodoDetailsController)
	e.PUT("/status/:id", c.ChangeTodoStatusController)
	e.DELETE("/:id", c.DeleteTodoController)
}
