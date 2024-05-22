package main

import (
	"fmt"
	"log"
	"os"
	"strings"

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

	dashList := GetListDash(list)
	return dashList
}

func main() {

	err := godotenv.Load()
	isError(err)

	bot, err := tbot.NewBotAPI(os.Getenv("TOKEN"))
	isError(err)

	bot.Debug = false

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tbot.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil {
			incomingChat := strings.ToLower(update.Message.Text)

			msg := tbot.NewMessage(update.Message.Chat.ID, update.Message.Text)
			senderName := update.Message.Chat.FirstName

			intro := fmt.Sprintf("Hello %s ", senderName)

			msgSelectDash := "Please select dashboard name below here ✅ !"

			switch incomingChat {
			case "/start":
				msg.Text = fmt.Sprintf("Welcome %s bot %s !", bot.Self.FirstName, senderName)
			case "/hi":
				msg.Text = intro + "i'm " + bot.Self.FirstName
			case "/help":
				msg.Text = intro + "this help"
			case "/about":
				msg.Text = intro + "this about"
			case "/dashboard":
				msg.Text = fmt.Sprintf("%s\n Please select domain dashboard below here ✅", intro)
				msg.ReplyMarkup = dashboardCommand

			case "ran 📶":
				msg.Text = msgSelectDash
				dashList := setRequestDashboard("ran")
				msg.ReplyMarkup = dashList

			case "core 📡":
				msg.Text = msgSelectDash
				dashList := setRequestDashboard("core")
				msg.ReplyMarkup = dashList

			case "netstat 🌎":
				msg.Text = msgSelectDash
				dashList := setRequestDashboard("netstat")
				msg.ReplyMarkup = dashList

			case "sales 🛒":
				msg.Text = msgSelectDash
				dashList := setRequestDashboard("sales")
				msg.ReplyMarkup = dashList

			default:
				msg.Text = "I don't know your command :("
			}

			if incomingChat != "/dashboard" {
				msg.ReplyToMessageID = update.Message.MessageID
			}
			if _, err := bot.Send(msg); err != nil {
				fmt.Println(err)
			}

		} else if update.CallbackQuery != nil {
			waitMsg := "Please wait while we are getting your request !"
			requestInline := update.CallbackQuery.Data
			targetInline := update.CallbackQuery.Message.Chat.ID

			callback := tbot.NewCallback(update.CallbackQuery.ID, requestInline)

			if _, err := bot.Request(callback); err != nil {
				fmt.Println(err)
			}
			bot.Send(tbot.NewMessage(targetInline, waitMsg))

			result, err := process_show_dashboard(requestInline)
			isError(err)

			img := tbot.FileBytes{
				Name:  "picture",
				Bytes: result,
			}

			_, err = bot.Send(tbot.NewPhoto(targetInline, img))

			if err != nil {
				bot.Send(tbot.NewMessage(targetInline, "Command unrecognized!"))
			}

		}
	}
}
