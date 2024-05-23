package main

import (
	"fmt"

	tbot "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func GetListDash(dashboardlist []string) tbot.InlineKeyboardMarkup {
	var btns []tbot.InlineKeyboardButton

	for _, dashname := range dashboardlist {
		btn := tbot.NewInlineKeyboardButtonData(dashname, dashname)
		btns = append(btns, btn)
	}
	var rows [][]tbot.InlineKeyboardButton
	for i := 0; i < len(btns); i += 2 {
		if i+1 < len(btns) {
			rows = append(rows, tbot.NewInlineKeyboardRow(btns[i], btns[i+1]))
		} else {
			rows = append(rows, tbot.NewInlineKeyboardRow(btns[i]))
		}
	}

	dashList := tbot.NewInlineKeyboardMarkup(rows...)
	return dashList
}

func GetDrillDash(dashname string, mapList []string) tbot.InlineKeyboardMarkup {
	var btns []tbot.InlineKeyboardButton

	for _, mapping := range mapList {
		Payload := fmt.Sprintf("%s %s", dashname, mapping)
		btn := tbot.NewInlineKeyboardButtonData(mapping, Payload)
		btns = append(btns, btn)
	}
	var rows [][]tbot.InlineKeyboardButton
	for i := 0; i < len(btns); i += 2 {
		if i+1 < len(btns) {
			rows = append(rows, tbot.NewInlineKeyboardRow(btns[i], btns[i+1]))
		} else {
			rows = append(rows, tbot.NewInlineKeyboardRow(btns[i]))
		}
	}

	dashList := tbot.NewInlineKeyboardMarkup(rows...)
	return dashList
}
