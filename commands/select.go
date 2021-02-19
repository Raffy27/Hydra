package commands

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func Perform(message *tgbotapi.Message) {
	switch message.Command() {
	case "ping":
		Ping()
	case "reset":
		Reset()
	case "info":
		Info()
	case "software":
		Software()
	case "sh":
		Shell(message.CommandArguments())
	case "file":
		UploadFile(message.CommandArguments())
	}
}
