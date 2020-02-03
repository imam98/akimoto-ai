package main

import (
	"log"
	"os"

	"github.com/go-telegram-bot-api/telegram-bot-api"
	"imam.miniproject/akimoto-ai/handlers"
)

var isAskingWeather = false
var quotes []string

func main() {
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TG_BOT_TOKEN"))
	if err != nil {
		log.Fatalln(err.Error())
	}

	bot.Debug = true
	log.Printf("Authorized on account %v\n", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	quotes, err = handlers.ParseQuotes("./assets/quotes")
	if err != nil {
		log.Fatalln("Error parsing file:", err.Error())
	}

	updates, err := bot.GetUpdatesChan(u)
	for update := range updates {
		if update.Message == nil {
			continue
		}

		if update.Message.IsCommand() {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
			switch update.Message.Command() {
			case "weather":
				if !isAskingWeather {
					msg.Text = "Alright I'm gonna process the data after you tell me your location"
					isAskingWeather = true
				} else {
					msg.Text = "You've already asked that a moment ago, please be kind and send me your location so I can start my work 😇"
				}
				bot.Send(msg)
			case "help":
				handlers.HelpHandler(update, bot)
			case "cheermeup":
				handlers.CheermeupHandler(update, bot, quotes)
			default:
				msg.Text = "Sorry, I don't acknowledge that command. Please try another one."
				bot.Send(msg)
			}
		} else {
			if isAskingWeather {
				if loc := update.Message.Location; loc != nil {
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
					msg.Text = "Processing data, may take a while..."
					bot.Send(msg)

					msg.Text = handlers.GenWeatherMsg(loc)
					msg.ParseMode = "markdown"
					bot.Send(msg)

					isAskingWeather = false
				} else {
					isAskingWeather = false
				}
			}
		}
	}
}
