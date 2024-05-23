package main

import (
	"fmt"
	"strings"

	tbot "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func handleMessage(update tbot.Update, bot *tbot.BotAPI) {
	incomingChat := strings.ToLower(update.Message.Text)
	msg := tbot.NewMessage(update.Message.Chat.ID, update.Message.Text)
	senderName := update.Message.Chat.FirstName
	intro := fmt.Sprintf("Hello %s ", senderName)
	msgDash := "Please select dashboard name below here ✅!"

	fmt.Println(senderName, ":", incomingChat)

	switch incomingChat {
	case "/start":
		msg.Text = fmt.Sprintf("Welcome into %s %s !", bot.Self.FirstName, senderName)
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
	if _, err := bot.Send(msg); err != nil {
		fmt.Println(err)
	}
}

func handleCallback(update tbot.Update, bot *tbot.BotAPI) {
	waitMsg := "Please wait while we are getting your request !"
	requestInline := update.CallbackQuery.Data
	targetInline := update.CallbackQuery.Message.Chat.ID

	callback := tbot.NewCallback(update.CallbackQuery.ID, requestInline)

	msg := tbot.NewMessage(targetInline, requestInline)

	if _, err := bot.Request(callback); err != nil {
		fmt.Println(err)
	}

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

		dashList := GetDrillDash(requestInline, list)
		msg.ReplyMarkup = dashList

		if _, err := bot.Send(msg); err != nil {
			fmt.Println(err)
		}

	} else {
		bot.Send(tbot.NewMessage(targetInline, waitMsg))
		result, cType, err := process_show_dashboard(requestInline)
		isError(err)
		img := tbot.FileBytes{
			Name:  cType,
			Bytes: result,
		}

		_, err = bot.Send(tbot.NewPhoto(targetInline, img))

		if err != nil {
			bot.Send(tbot.NewMessage(targetInline, "Command unrecognized!"))
		}

	}
}
