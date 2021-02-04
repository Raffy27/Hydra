package api

import (
	"github.com/Raffy27/Hydra/util"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

var (
	//Bot represents the API used for interacting with Telegram.
	Bot *tgbotapi.BotAPI
	//Updates is the channel interface for Telegram commands.
	Updates tgbotapi.UpdatesChannel
)

func init() {
	var err error
	Bot, err = tgbotapi.NewBotAPI("1663927036:AAGiFpqOWXzFsBFVK6t7rWa-DpnlfZzIPYE")
	util.Handle(err)
	//Bot.Debug = true

	upd := tgbotapi.NewUpdate(0)
	upd.Timeout = 30
	Updates, err = Bot.GetUpdatesChan(upd)
	util.Handle(err)
}

//SendMessage is a convenience function for replying to simple Telegram messages.
func SendMessage(org *tgbotapi.Message, msg string) {
	m := tgbotapi.NewMessage(org.Chat.ID, msg)
	Bot.Send(m)
}
