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

	var senderName string
	if update.Message.Chat.IsSuperGroup() {
		senderName = update.Message.From.FirstName
	} else {
		senderName = update.Message.Chat.FirstName
	}

	intro := fmt.Sprintf("Hello %s ", senderName)
	msgDash := "Please select dashboard name below here ✅!"

	TimeMessage := time.Unix(int64(update.Message.Date), 0).Format("2006-01-02 15:04:05")

	fmt.Println(fmt.Sprintf("%s - %s : %s", TimeMessage, senderName, incomingChat))
	ReplyMsg := true

	switch {
	case strings.HasPrefix(incomingChat, "/start"):
		msg.Text = fmt.Sprintf("Welcome into %s %s !\nYou can access and request dashboard by typing /dashboard or click in menu !", bot.Self.FirstName, senderName)
	case strings.HasPrefix(incomingChat, "/help"):
		msg.Text = "You can tap /dashboard and select domain what you needed,\nthen the dashboard name button will showed up!,\nif u still need help feel free to contact me at https://t.me/fajarh2207"
	case strings.HasPrefix(incomingChat, "/about"):
		msg.Text = "Developed at May 2024\nThis bot developed using Go Language and\nusing library github.com/go-telegram-bot-api/telegram-bot-api/v5,\nBuild with love by https://t.me/fajarh2207"

	case strings.HasPrefix(incomingChat, "/dashboard"):
		msg.Text = fmt.Sprintf("%s\nPlease select domain dashboard below here ✅", intro)
		msg.ReplyMarkup = dashboardCommand
	case incomingChat == "ran 📶":
		msg.Text = msgDash
		msg.ReplyMarkup = setRequestDashboard("ran")
	case incomingChat == "core 📡":
		msg.Text = msgDash
		msg.ReplyMarkup = setRequestDashboard("core")
	case incomingChat == "netstat 🌎":
		msg.Text = msgDash
		msg.ReplyMarkup = setRequestDashboard("netstat")
	case incomingChat == "sales 🛒":
		msg.Text = msgDash
		msg.ReplyMarkup = setRequestDashboard("sales")
	default:
		ReplyMsg = false
	}
	if ReplyMsg {
		msg.ReplyToMessageID = update.Message.MessageID
		_, err := bot.Send(msg)
		isErrorMessage(err, update.Message.Chat.ID, bot)
	}
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
