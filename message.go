package main

import (
	"fmt"
	"strings"
	"time"

	tbot "github.com/go-telegram-bot-api/telegram-bot-api/v5"
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

func handleMessage(update tbot.Update, bot *tbot.BotAPI) {
	incomingChat := strings.ToLower(update.Message.Text)
	msg := tbot.NewMessage(update.Message.Chat.ID, update.Message.Text)
	senderName := update.Message.Chat.FirstName
	intro := fmt.Sprintf("Hello %s ", senderName)
	msgDash := "Please select dashboard name below here ✅!"

	TimeMessage := time.Unix(int64(update.Message.Date), 0).Format("2006-01-02 15:04:05")
	fmt.Println(fmt.Sprintf("%s - %s : %s", TimeMessage, senderName, incomingChat))

	switch incomingChat {

	case "/start":
		msg.Text = fmt.Sprintf("Welcome into %s %s !\nYou can access and request dashboard by typing /dashboard or click in menu !", bot.Self.FirstName, senderName)
	case "/hi":
		msg.Text = intro + "i'm " + bot.Self.FirstName
	case "/help":
		msg.Text = intro + "this help"
	case "/about":
		msg.Text = intro + "this about"
	case "/dashboard":
		msg.Text = fmt.Sprintf("%s\nPlease select domain dashboard below here ✅", intro)
		msg.ReplyMarkup = dashboardCommand
	case "ran 📶":
		msg.Text = msgDash
		msg.ReplyMarkup = setRequestDashboard("ran")
	case "core 📡":
		msg.Text = msgDash
		msg.ReplyMarkup = setRequestDashboard("core")
	case "netstat 🌎":
		msg.Text = msgDash
		msg.ReplyMarkup = setRequestDashboard("netstat")
	case "sales 🛒":
		msg.Text = msgDash
		msg.ReplyMarkup = setRequestDashboard("sales")
	default:
		msg.Text = "I don't know your command :("
	}
	msg.ReplyToMessageID = update.Message.MessageID
	_, err := bot.Send(msg)
	isErrorMessage(err, update.Message.Chat.ID, bot)
}

func handleCallback(update tbot.Update, bot *tbot.BotAPI) {
	waitMsg := "Please wait while we are getting your request !"
	requestInline := update.CallbackQuery.Data
	targetInline := update.CallbackQuery.Message.Chat.ID
	TimeMessage := time.Unix(int64(update.CallbackQuery.Message.Date), 0).Format("2006-01-02 15:04:05")

	callback := tbot.NewCallback(update.CallbackQuery.ID, requestInline)

	msg := tbot.NewMessage(targetInline, requestInline)

	_, err := bot.Request(callback)
	isErrorMessage(err, targetInline, bot)

	if strings.Count(requestInline, " ") == 1 {
		var list []string
		switch requestInline {
		case "core_perf area":
			list = neArea
		case "core_perf ne":
			list = ne
		default:
			if strings.Contains(requestInline, "circle") {
				list = Circle
			} else if strings.Contains(requestInline, "region") {
				list = region
			} else if strings.Contains(requestInline, "area") {
				list = area
			} else if strings.Contains(requestInline, "ggsn") {
				list = ggsn
			} else if strings.Contains(requestInline, "location") {
				list = ggsnArea
			} else {
				list = nil
			}
		}
		msg.ReplyMarkup = GetDrillDash(requestInline, list)

		_, err := bot.Send(msg)
		isErrorMessage(err, targetInline, bot)

	} else {
		bot.Send(tbot.NewMessage(targetInline, waitMsg))
		result, cType, err := process_show_dashboard(requestInline)
		isErrorMessage(err, targetInline, bot)

		fmt.Println(TimeMessage, "-", "Request", ":", requestInline)

		// Sending ImageHandler
		img := tbot.FileBytes{
			Name:  cType,
			Bytes: result,
		}
		msg := tbot.NewPhoto(targetInline, img)
		msg.Caption = requestInline

		_, err = bot.Send(msg)
		isErrorMessage(err, targetInline, bot)

	}
}
