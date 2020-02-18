package main

import (
	"log"
	"os"

	"github.com/go-telegram-bot-api/telegram-bot-api"
	"imam.miniproject/akimoto-ai/handlers"
	"imam.miniproject/akimoto-ai/models"
	"imam.miniproject/akimoto-ai/storage"
)

var isAskingWeather = false

func main() {
	mux, err := models.NewCommandRouter(os.Getenv("TG_BOT_TOKEN"), true)
	if err != nil {
		log.Fatalln(err.Error())
	}

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	err = storage.Prepare()
	if err != nil {
		log.Fatalln("Error parsing file:", err.Error())
	}

	mux.Handle("help", handlers.HelpHandler)
	mux.Handle("cheermeup", handlers.CheermeupHandler)
	mux.Handle("weather", handlers.WeatherReportHandler)
	mux.Serve(u)

	// updates, err := bot.GetUpdatesChan(u)
	// for update := range updates {
	// 	if update.Message == nil {
	// 		continue
	// 	}

	// 	if update.Message.IsCommand() {
	// 		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
	// 		switch update.Message.Command() {
	// 		case "weather":
	// 			if !isAskingWeather {
	// 				msg.Text = "Alright I'm gonna process the data after you tell me your location"
	// 				isAskingWeather = true
	// 			} else {
	// 				msg.Text = "You've already asked that a moment ago, please be kind and send me your location so I can start my work 😇"
	// 			}
	// 			bot.Send(msg)
	// 		case "help":
	// 			handlers.HelpHandler(update, bot)
	// 		case "cheermeup":
	// 			handlers.CheermeupHandler(update, bot, quotes)
	// 		default:
	// 			msg.Text = "Sorry, I don't acknowledge that command. Please try another one."
	// 			bot.Send(msg)
	// 		}
	// 	} else {
	// 		if isAskingWeather {
	// 			if loc := update.Message.Location; loc != nil {
	// 				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
	// 				msg.Text = "Processing data, may take a while..."
	// 				bot.Send(msg)

	// 				msg.Text = handlers.GenWeatherMsg(loc)
	// 				msg.ParseMode = "markdown"
	// 				bot.Send(msg)

	// 				isAskingWeather = false
	// 			} else {
	// 				isAskingWeather = false
	// 			}
	// 		}
	// 	}
	// }
}
