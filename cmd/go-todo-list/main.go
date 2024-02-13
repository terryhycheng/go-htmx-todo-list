package main

import (
	"github.com/labstack/echo/v4"
	"github.com/terryhycheng/go-todo-list/internal/controllers"
	"github.com/terryhycheng/go-todo-list/internal/routes"
)

func main() {
	port := "3000"
	e := echo.New()	

	// Serve static files
	e.Static("/css", "build/css") 
	e.Static("/assets", "assets") 
	
	// General Routes
	e.GET("/", controllers.HomePageController)
	
	// Todo routes
	todoGroup := e.Group("/todo")
	routes.TodoSubRoutes(todoGroup)

	e.Logger.Fatal(e.Start(":" + port))
}