package main

import (
	"log"

	"github.com/Raffy27/Hydra/api"
)

func main() {
	log.Println("Logged in as", api.Bot.Self.UserName)
	for u := range api.Updates {
		if u.Message == nil {
			continue
		}
		log.Printf("[%s] %s", u.Message.From.UserName, u.Message.Text)
		if u.Message.IsCommand() {
			switch u.Message.Command() {
			case "ping":
				api.SendMessage(u.Message, "Pong!")
			}
		}
	}
}
