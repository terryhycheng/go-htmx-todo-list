package main

import (
	"github.com/labstack/echo/v4"
	"github.com/terryhycheng/go-todo-list/api/controllers"
)

func main() {
	port := "3000"
	e := echo.New()

	e.GET("/", controllers.HomePageController)
	e.Static("/css", "build/css") 

	e.Logger.Fatal(e.Start(":" + port))
}