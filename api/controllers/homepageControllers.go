package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/terryhycheng/go-todo-list/internal/helpers"
	"github.com/terryhycheng/go-todo-list/internal/types"
	"github.com/terryhycheng/go-todo-list/web/pages"
)


func HomePageController(c echo.Context) error {
	card1 := &types.Card{
		Status: "normal",
	}

	card2 := &types.Card{
		Status: "finished",
		IsDone: true,
		Title: "Wash dishes is quite annoying",
		Content: "I hate washing dishes",
	}

	card3 := &types.Card{
		Status: "urgent",
	}
	
	card4 := &types.Card{
		Status: "random",
	}

	cards := &[]*types.Card{card1, card2, card3, card4}
	
	return helpers.Render(c, http.StatusOK, pages.Homepage(cards))
}