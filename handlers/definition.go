package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"text/template"

	"github.com/go-telegram-bot-api/telegram-bot-api"
	"imam.miniproject/akimoto-ai/models"
)

func DefinitionHandler(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	word := update.Message.CommandArguments()
	result, err := requestDefinition(word)
	if err != nil {
		if err.Error() == "404" {
			msgText := fmt.Sprintf("The definition of %v is not found in my dictionary", word)
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, msgText)
			bot.Send(msg)
			return
		}

		log.Fatalln(err)
	}

	tpl := template.Must(template.New("msg").Parse(models.DefTmpl))
	sb := &strings.Builder{}
	for index, val := range result.Results[0].LexicalEntries {
		if index == 3 {
			break
		}

		val.LexicalCategory.Text = strings.ToLower(val.LexicalCategory.Text)
		if err := tpl.Execute(sb, val); err != nil {
			log.Fatalln(err.Error())
		}

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, sb.String())
		msg.ParseMode = "markdown"
		bot.Send(msg)
		sb.Reset()
	}
}

func requestDefinition(word string) (models.DictionaryEntry, error) {
	client := &http.Client{}
	url := fmt.Sprintf("%ventries/en-us/%v", models.OdBasePath, word)
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return models.DictionaryEntry{}, err
	}

	request.Header.Set("Accept", "application/json")
	request.Header.Add("app_id", os.Getenv("OD_APP_ID"))
	request.Header.Add("app_key", os.Getenv("OD_API_KEY"))

	q := request.URL.Query()
	q.Add("fields", models.OdFields)
	q.Add("strictMatch", "false")
	request.URL.RawQuery = q.Encode()

	response, err := client.Do(request)
	if err != nil {
		return models.DictionaryEntry{}, err
	}
	defer response.Body.Close()

	if response.StatusCode == http.StatusNotFound {
		return models.DictionaryEntry{}, fmt.Errorf("%v", http.StatusNotFound)
	}

	var data models.DictionaryEntry
	if err := json.NewDecoder(response.Body).Decode(&data); err != nil {
		return models.DictionaryEntry{}, err
	}

	return data, nil
}
