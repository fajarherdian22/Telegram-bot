package main

import (
	"fmt"

	tbot "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func isError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func isErrorMessage(err error, targetInline int64, bot *tbot.BotAPI) {
	if err != nil {
		bot.Send(tbot.NewMessage(targetInline, "Error : "+err.Error()))
	}
}
