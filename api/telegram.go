package api

import (
	"strings"

	"github.com/Raffy27/Hydra/util"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

var (
	//Bot represents the API used for interacting with Telegram.
	Bot *tgbotapi.BotAPI
	//Updates is the channel interface for Telegram commands.
	Updates tgbotapi.UpdatesChannel
)

const msgSize = 4096

func init() {
	//This is here for a reason!
	//"Bot" absolutely needs global scope
	var err error
	Bot, err = tgbotapi.NewBotAPI(util.BotToken)
	util.Handle(err)
	//Bot.Debug = true

	upd := tgbotapi.NewUpdate(0)
	upd.Timeout = 30
	Updates, err = Bot.GetUpdatesChan(upd)
	util.Handle(err)
}

func SendFragmented(msg string, sep string, prefix string, suffix string) tgbotapi.Message {
	var m tgbotapi.Message
	cfg := tgbotapi.NewMessage(util.ChatID, prefix+msg+suffix)
	cfg.ParseMode = "Markdown"

	if len(msg) > msgSize {
		l := strings.Split(msg, sep)
		f := prefix
		for _, v := range l {
			if len(f)+len(v)+len(suffix) <= msgSize {
				f += v + sep
				continue
			}
			f += suffix
			cfg.Text = f
			Bot.Send(cfg)
			f = prefix
		}
		cfg.Text = f + suffix
	}

	m, _ = Bot.Send(cfg)
	return m
}
