package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
	"text/template"
	"time"

	"github.com/go-telegram-bot-api/telegram-bot-api"

	"imam.miniproject/akimoto-ai/models"
	"imam.miniproject/akimoto-ai/storage"
)

func WeatherReportHandler(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
	user := storage.GetUser(update.Message.Chat.ID)
	if !user.IsAskingWeather {
		msg.Text = "Alright I'm gonna process the data after you tell me your location"
		user.IsAskingWeather = true
	} else {
		msg.Text = "You've already asked that a moment ago, please be kind and send me your location so I can start my work 😇"
	}
	bot.Send(msg)
}

func GenWeatherMsg(loc *tgbotapi.Location) string {
	report, err := requestWeatherReport(loc.Latitude, loc.Longitude)
	if err != nil {
		return err.Error()
	}

	humidity := report.Weather.Humidity * 100.0
	precipProb := report.Weather.PrecipProb * 100.0

	report.Weather.Humidity = humidity
	report.Weather.PrecipProb = precipProb

	funcs := template.FuncMap{
		"ic2emot": convertIconIntoEmoji,
		"time":    timeFromUnix,
	}

	tpl := template.Must(template.New("msg").Funcs(funcs).Parse(models.MsgTmpl))
	sb := &strings.Builder{}
	if err := tpl.Execute(sb, report); err != nil {
		return err.Error()
	}

	return sb.String()
}

func requestWeatherReport(latitude float64, longitude float64) (models.Report, error) {
	client := &http.Client{}
	url := fmt.Sprintf("%v%v/%v,%v?exclude=%v&units=si", models.DsBasePath, os.Getenv("DS_API_KEY"), latitude, longitude, models.Excluded)

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return models.Report{}, err
	}

	response, err := client.Do(request)
	if err != nil {
		return models.Report{}, err
	}
	defer response.Body.Close()

	var data models.Report
	if err := json.NewDecoder(response.Body).Decode(&data); err != nil {
		return models.Report{}, err
	}

	return data, nil
}

func convertIconIntoEmoji(icon string) string {
	var emoji string
	switch icon {
	case "clear-day", "clear-night":
		emoji = models.IconClear
	case "partly-cloudy-day", "partly-cloudy-night":
		emoji = models.IconPartialCloud
	case "cloudy":
		emoji = models.IconHumid
	case "rain":
		emoji = models.IconRain
	case "thunderstorm":
		emoji = models.IconThunderstorm
	default:
		emoji = ""
	}

	return emoji
}

func timeFromUnix(unix int64) string {
	t := time.Unix(unix, 0)
	return t.Format("15:04")
}
