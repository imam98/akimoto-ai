package handlers

import (
	"strings"

	"github.com/go-telegram-bot-api/telegram-bot-api"
)

func HelpHandler(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	arg := update.Message.CommandArguments()
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")

	switch arg {
	case "":
		msg.Text = `Currently available command:

		/weather used to get current weather forecast in your location
		/define [word] 

		That's all I can do right now, I'm pretty sure my paisen will add more features soon`
		msg.Text = strings.ReplaceAll(msg.Text, "\t", "")
	case "weather":
		msg.Text = "To get weather report, use command /weather then send me your location using telegram share location feature so I could process the report"
	}

	bot.Send(msg)
}
