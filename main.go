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
	mux.LocationMsgHandler = handlers.LocationHandler
	mux.Serve(u)
}
