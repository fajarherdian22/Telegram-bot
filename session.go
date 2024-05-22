package main

import (
	"strings"

	tbot "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// var DashRan = tbot.NewInlineKeyboardMarkup(
// 	tbot.NewInlineKeyboardRow(
// 		tbot.NewInlineKeyboardButtonData("netstat", "netstat"),
// 		tbot.NewInlineKeyboardButtonData("list", "list"),
// 		tbot.NewInlineKeyboardButtonData("dashboard 217", "dashboard 217"),
// 	),
// )

// func GetListDash(list []string) tbot.InlineKeyboardMarkup {

// 	var DashRan = tbot.NewInlineKeyboardMarkup(
// 		tbot.NewInlineKeyboardRow(
// 			for _, element := range list {
// 				if !strings.Contains(element, " ") {
// 					tbot.NewInlineKeyboardButtonData(element, element)
// 				}
// 			}
// 		),
// 	)

// 	return DashRan

// }

func GetListDash(list []string) tbot.InlineKeyboardMarkup {
	var btns []tbot.InlineKeyboardButton

	// Create buttons for each element without spaces
	for _, element := range list {
		if !strings.Contains(element, " ") {
			btn := tbot.NewInlineKeyboardButtonData(element, element)
			btns = append(btns, btn)
		}
	}

	var rows [][]tbot.InlineKeyboardButton

	// Group buttons into rows of two
	for i := 0; i < len(btns); i += 2 {
		if i+1 < len(btns) {
			rows = append(rows, tbot.NewInlineKeyboardRow(btns[i], btns[i+1]))
		} else {
			rows = append(rows, tbot.NewInlineKeyboardRow(btns[i]))
		}
	}

	dashRan := tbot.NewInlineKeyboardMarkup(rows...)
	return dashRan
}

// func GetListDash(list []string) tbot.InlineKeyboardMarkup {

// 	var btns []tbot.InlineKeyboardButton

// 	for _, element := range list {
// 		if !strings.Contains(element, " ") {
// 			btn := tbot.NewInlineKeyboardButtonData(element, element)
// 			btns = append(btns, btn)
// 		}
// 	}

// 	var rows []tbot.InlineKeyboardButton

// 	for index, _ := range list {
// 		index += 1
// 		if index%2 == 0 {
// 			fmt.Println(index)
// 		}
// 	}

// 	dashRan := tbot.NewInlineKeyboardMarkup(tbot.NewInlineKeyboardRow(btns...))
// 	return dashRan
// }

var DashCore = tbot.NewInlineKeyboardMarkup(
	tbot.NewInlineKeyboardRow(
		tbot.NewInlineKeyboardButtonData("core_apn", "core_apn"),
		tbot.NewInlineKeyboardButtonData("core_ps", "core_ps"),
		tbot.NewInlineKeyboardButtonData("core_ytd", "core_ytd"),
	),
)

var DashOther = tbot.NewInlineKeyboardMarkup(
	tbot.NewInlineKeyboardRow(
		tbot.NewInlineKeyboardButtonData("1", "1"),
		tbot.NewInlineKeyboardButtonData("2", "2"),
	),
)
