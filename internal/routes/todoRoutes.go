package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/terryhycheng/go-todo-list/internal/controllers"
)

func TodoSubRoutes(e *echo.Group){
	e.POST("", controllers.AddTodoController)
	// e.PUT("/:id", controllers.ChangeTodoDetailsController)
	e.PUT("/status/:id", controllers.ChangeTodoStatusController)	
	e.DELETE("/:id", controllers.DeleteTodoController)
}