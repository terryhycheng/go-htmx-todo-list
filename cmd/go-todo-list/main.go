package main

import (
	"github.com/labstack/echo/v4"
	"github.com/terryhycheng/go-todo-list/internal/controllers"
	"github.com/terryhycheng/go-todo-list/internal/helpers"
	"github.com/terryhycheng/go-todo-list/internal/models"
)

func main() {
	port := "3000"
	e := echo.New()

	// database connection
	helpers.Connect()
	db := helpers.GetDB()
	m := models.NewTodoRepository(db)
	c := controllers.NewTodoControllers(m)
	hc := controllers.NewHomepageControllers(m)

	// Migrate the schema
	db.AutoMigrate(&models.TodoGorm{})

	// Serve static files
	e.Static("/css", "build/css")
	e.Static("/assets", "assets")

	// General Routes
	e.GET("/", hc.HomePageController)

	// API Routes
	v1 := e.Group("/api/v1")
	{
		todo := v1.Group("/todo")
		{

			todo.POST("", c.AddTodoController)
			// e.PUT("/:id", controllers.ChangeTodoDetailsController)
			todo.PUT("/status/:id", c.ChangeTodoStatusController)
			todo.DELETE("/:id", c.DeleteTodoController)
		}
	}

	e.Logger.Fatal(e.Start(":" + port))
}
