package main

import (
	"github.com/labstack/echo/v4"
	"github.com/terryhycheng/go-todo-list/internal/controllers"
	"github.com/terryhycheng/go-todo-list/internal/helpers"
	"github.com/terryhycheng/go-todo-list/internal/models"
	"github.com/terryhycheng/go-todo-list/internal/routes"
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

	// Todo routes
	todoGroup := e.Group("/todo")
	routes.TodoSubRoutes(todoGroup, c)

	e.Logger.Fatal(e.Start(":" + port))
}
