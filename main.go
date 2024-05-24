package main

import (
	"log"
	"os"

	tbot "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	isError(err)

	bot, err := tbot.NewBotAPI(os.Getenv("TOKEN"))
	isError(err)

	bot.Debug = false

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tbot.NewUpdate(1)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		go Messagefunc(update, bot)
	}
}

func Messagefunc(update tbot.Update, bot *tbot.BotAPI) {
	if update.Message != nil {
		handleMessage(update, bot)
	} else if update.CallbackQuery != nil {
		handleCallback(update, bot)
	}
}
