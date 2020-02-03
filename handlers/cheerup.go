package handlers

import (
	"bufio"
	"math/rand"
	"os"
	"time"

	"github.com/go-telegram-bot-api/telegram-bot-api"
)

func ParseQuotes(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var quotes []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		quotes = append(quotes, scanner.Text())
	}

	return quotes, nil
}

func CheermeupHandler(update tgbotapi.Update, bot *tgbotapi.BotAPI, quotes []string) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, getQuote(quotes))
	bot.Send(msg)
}

func getQuote(quotes []string) string {
	seed := time.Now().Unix()
	r := rand.New(rand.NewSource(seed))
	return quotes[r.Intn(len(quotes))]
}
