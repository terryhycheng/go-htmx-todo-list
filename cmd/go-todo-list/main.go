package main

import (
	"github.com/labstack/echo/v4"
	"github.com/terryhycheng/go-todo-list/internal/controllers"
)

func main() {
	port := "3000"
	e := echo.New()	

	e.GET("/", controllers.HomePageController)
	e.POST("/add", controllers.AddTodoController)
	e.Static("/css", "build/css") 
	e.Static("/assets", "assets") 

	e.Logger.Fatal(e.Start(":" + port))
}