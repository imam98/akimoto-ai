package handlers

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"imam.miniproject/akimoto-ai/storage"
)

func CheermeupHandler(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, storage.GetQuote())
	msg.ParseMode = "markdown"
	bot.Send(msg)
}
