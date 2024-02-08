package helpers

import "github.com/terryhycheng/go-todo-list/internal/types"


func GetButtonColour(card *types.Card) string {
	if card.IsDone {
		return "bg-[#102916]"
	} else {
		if card.Status == "urgent" {
			return "bg-[#2B2511]"
		} else {
			return "bg-[#11292B]"
		}
	}
}

func GetCardBorder(isDone bool) string {
	if isDone {
		return "border border-[#5CA05A]"
	}

	return ""
}

func GetCardIconPath(card *types.Card) string {
	var iconName string
	
	if card.IsDone {
		iconName = "CheckOutline"
	} else {

		if card.Status == "urgent" {
			iconName = "LightningBoltOutline"
	} else {
		iconName = "CalendarOutline"
	}

	}

	return "/assets/" + iconName + ".svg"	
}

func GetCardIconBgColour(card *types.Card) string {
	if card.IsDone {
		return "bg-[#5CA05A]"
	}

	if card.Status == "urgent" {
		return "bg-[#B6A155]"
	}

	return "bg-[#112C2E]"
}