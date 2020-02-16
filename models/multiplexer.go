package models

import (
	"log"

	"github.com/go-telegram-bot-api/telegram-bot-api"
)

type Multiplexer struct {
	bot    *tgbotapi.BotAPI
	routes map[string]Handler
}

type Handler func(bot *tgbotapi.BotAPI, update tgbotapi.Update)

func NewCommandRouter(token string, debugMode bool) (Multiplexer, error) {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return Multiplexer{}, err
	}

	bot.Debug = debugMode
	mux := Multiplexer{
		bot:    bot,
		routes: map[string]Handler{},
	}

	log.Printf("Authorized on account %v\n", bot.Self.UserName)
	return mux, nil
}

func (mux *Multiplexer) HandleCommand(command string, handler Handler) {
	mux.routes[command] = handler
}

func (mux *Multiplexer) Serve(updateConfig tgbotapi.UpdateConfig) error {
	updates, err := mux.bot.GetUpdatesChan(updateConfig)
	if err != nil {
		return err
	}

	for update := range updates {
		if update.Message == nil {
			continue
		}

		if update.Message.IsCommand() {
			if handle, ok := mux.routes[update.Message.Command()]; ok {
				handle(mux.bot, update)
			} else {
				id := update.Message.Chat.ID
				msg := tgbotapi.NewMessage(id, "Sorry, I don't acknowledge that command. Please try another one.")
				mux.bot.Send(msg)
			}
		}
	}

	return nil
}
