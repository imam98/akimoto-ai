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

		/weather to get current weather forecast in your location
		/cheermeup use this when you are feeling down
		/define [word] to get the definition(s) of such [word]
		/help [command] to get more detailed information of such [command]

		That's all I can do right now, I'm pretty sure my paisen will add more features soon`
		msg.Text = strings.ReplaceAll(msg.Text, "\t", "")
	case "weather":
		msg.Text = "To get weather report, use command /weather then send me your location using telegram share location feature so I could process the report"
	case "cheermeup":
		msg.Text = "Whenever you are feeling down or unmotivated, use this command and I will send you some words that maybe could cheer you up even for a little bit :)"
	case "define":
		msg.Text = "To use this command, type /define [word] where [word] is the word to search"
		bot.Send(msg)
		msg.Text = "For example: /define hello"
		bot.Send(msg)
		msg.Text = "If the word is listed in my dictionary, I will give you maximum 3 definitions of that word"
	case "help":
		msg.Text = "Help command is used as the manual page of how you could utilize my capabilities which implemented by my paisen"
		bot.Send(msg)
		msg.Text = "type /help without any arguments to get the list of all my capabilities and its short description"
		bot.Send(msg)
		msg.Text = "type /help [command] to get an insight description about such command, where [command] is the command that you want to search"
		bot.Send(msg)
		msg.Text = "[command] should be typed in lowercase without the backslash (/) prefix"
		bot.Send(msg)
		msg.Text = "For example: /help weather"
	default:
		msg.Text = "I'm sorry, I don't think that I recognized that command"
		bot.Send(msg)
		msg.Text = "If you are confused about the usage of help command, I suggest you to type /help help"
	}

	bot.Send(msg)
}
