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
	var err error
	Bot, err = tgbotapi.NewBotAPI("1663927036:AAGiFpqOWXzFsBFVK6t7rWa-DpnlfZzIPYE")
	util.Handle(err)
	Bot.Debug = true

	upd := tgbotapi.NewUpdate(0)
	upd.Timeout = 30
	Updates, err = Bot.GetUpdatesChan(upd)
	util.Handle(err)
}

//SendMessage is a convenience function for replying to simple Telegram messages.
func SendMessage(org *tgbotapi.Message, txt string) tgbotapi.Message {
	m := tgbotapi.NewMessage(org.Chat.ID, txt)
	msg, _ := Bot.Send(m)
	return msg
}

func SendFragmented(msg string, sep string, prefix string, suffix string) tgbotapi.Message {
	var m tgbotapi.Message
	if len(msg) > msgSize {
		l := strings.Split(msg, sep)
		f := prefix
		for _, v := range l {
			if len(f)+len(v)+len(suffix) <= msgSize {
				f += v + sep
				continue
			}
			f += suffix
			cfg := tgbotapi.NewMessage(util.ChatID, f)
			cfg.ParseMode = "Markdown"
			m, _ = Bot.Send(cfg)
			f = prefix
		}
		return m
	}
	cfg := tgbotapi.NewMessage(util.ChatID, msg)
	cfg.ParseMode = "Markdown"
	m, _ = Bot.Send(cfg)
	return m
}
