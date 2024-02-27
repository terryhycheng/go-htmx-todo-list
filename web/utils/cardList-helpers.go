package utils

import "github.com/terryhycheng/go-todo-list/internal/models"

func GetButtonColour(card *models.TodoGorm) string {
	if card.IsDone {
		return "bg-darkGreen"
	} else {
		if card.Priority == "urgent" {
			return "bg-darkYellow"
		} else {
			return "bg-dark"
		}
	}
}

func GetCardBorder(isDone bool) string {
	if isDone {
		return "border border-green"
	}

	return ""
}

func GetCardIconPath(card *models.TodoGorm) string {
	var iconName string

	if card.IsDone {
		iconName = "CheckOutline"
	} else {

		if card.Priority == "urgent" {
			iconName = "LightningBoltOutline"
		} else {
			iconName = "CalendarOutline"
		}

	}

	return "/assets/" + iconName + ".svg"
}

func GetCardIconBgColour(card *models.TodoGorm) string {
	if card.IsDone {
		return "bg-green"
	}

	if card.Priority == "urgent" {
		return "bg-yellow"
	}

	return "bg-dark"
}
