package handlers

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"imam.miniproject/akimoto-ai/storage"
)

func LocationHandler(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	loc := update.Message.Location
	user := storage.GetUser(update.Message.Chat.ID)

	if user.IsAskingWeather {
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
		msg.Text = "Processing data, may take a while..."
		bot.Send(msg)

		msg.Text = genWeatherMsg(loc)
		msg.ParseMode = "markdown"
		bot.Send(msg)

		user.IsAskingWeather = false
	}
}
