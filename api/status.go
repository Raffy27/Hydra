package api

import (
	"fmt"
	"time"

	"github.com/Raffy27/Hydra/util"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

var (
	//StopHeartbeat signals the associated goroutine to stop when it is closed.
	StopHeartbeat = make(chan struct{})
)

//Heartbeat is the function that provides status updates.
func Heartbeat() {
	ticker := time.NewTicker(util.Interval * time.Second)
	//Set caption
	caption := tgbotapi.NewEditMessageCaption(util.ChatID, util.Genesis, "Genesis")
	Bot.Send(caption)
	//Create edit struct
	status := tgbotapi.NewEditMessageText(util.ChatID, util.Genesis, "")
	for {
		select {
		case <-ticker.C:
			t := time.Now()
			status.Text = fmt.Sprintf("Beep boop!\n%s", t.Format(time.RFC3339))
			_, err := Bot.Send(status)
			util.Handle(err)
		case <-StopHeartbeat:
			ticker.Stop()
			return
		}
	}
}
