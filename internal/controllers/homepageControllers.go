package controllers

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/terryhycheng/go-todo-list/internal/helpers"
	"github.com/terryhycheng/go-todo-list/internal/models"
	"github.com/terryhycheng/go-todo-list/web/pages"
)

var ctx = context.Background()
var client = helpers.RedisClient()


func HomePageController(c echo.Context) error {
	todoIds, err := client.LRange(ctx, "todos", 0, -1).Result()

	if err != nil {
		return c.String(http.StatusInternalServerError, "Failed to get todos from Redis: " + err.Error())
	}

	todos := models.NewTodos()

	for _, todoId := range todoIds {
		todo, err := client.JSONGet(ctx, "todo:" + todoId, ".").Result()

		if err != nil {
			return c.String(http.StatusInternalServerError, "Failed to get todo details: " + err.Error())
		}

		todoModel := models.Todo{}

		json.Unmarshal([]byte(todo), &todoModel)

		_, addErr := todos.Add(todoModel.Title, todoModel.Content, todoModel.Status)

		if addErr != nil {
			return c.String(http.StatusInternalServerError, "Failed to add todo to the list: " + addErr.Error())
		}
	}

	return helpers.Render(c, http.StatusOK, pages.Homepage(todos))
}

func AddTodoController(c echo.Context) error {
	todo1 := models.NewTodo("Todo 1", "This is the first todo", "normal")
	todo2 := models.NewTodo("Todo 2", "This is the second todo", "normal")
	todo3 := models.NewTodo("Todo 3", "This is the third todo", "normal")

	todos := &models.Todos{todo1, todo2, todo3}

	for _, todo := range *todos {
		_, listPusherr := client.LPush(ctx, "todos", todo.Id.String()).Result()
		if listPusherr != nil {
			return c.String(http.StatusInternalServerError, "Failed to add todo to the list :" + listPusherr.Error())
		}

		json, jsonMarshalErr := json.Marshal(todo)

		if jsonMarshalErr != nil {
			return c.String(http.StatusInternalServerError, "Failed to marshal todo json :" + jsonMarshalErr.Error())
		}

		_, setJsonErr := client.JSONSet(ctx, "todo:" + todo.Id.String(), ".", json).Result()

		if setJsonErr != nil {
			return c.String(http.StatusInternalServerError, "Failed to set todo json :" + setJsonErr.Error())
		}
	}

	return c.String(http.StatusOK, "Added")
}