package main

import (
	"log"
	"os"

	tbot "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
)

var dashboardCommand = tbot.NewReplyKeyboard(
	tbot.NewKeyboardButtonRow(
		tbot.NewKeyboardButton("ran 📶"),
		tbot.NewKeyboardButton("core 📡"),
	),
	tbot.NewKeyboardButtonRow(
		tbot.NewKeyboardButton("netstat 🌎"),
		tbot.NewKeyboardButton("sales 🛒"),
	),
)

func setRequestDashboard(domain string) tbot.InlineKeyboardMarkup {
	list, err := process_list_dashboard_by_cat(domain)
	isError(err)
	return GetListDash(list)
}

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
		if update.Message != nil {
			handleMessage(update, bot)
		} else if update.CallbackQuery != nil {
			handleCallback(update, bot)
		}
	}

}
