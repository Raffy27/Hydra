package util

import (
	"fmt"
	"os/exec"
	"sort"
	"strconv"
	"strings"
	"syscall"

	"golang.org/x/sys/windows/registry"
)

const (
	upMask = 1 << iota
	enMask

	wmiCommand = "$i=(Get-CimInstance -n root/SecurityCenter2 -cl %s);foreach($v in $i){$v.displayName;$v.productState}"
	softKeys   = "SOFTWARE\\WOW6432Node\\Microsoft\\Windows\\CurrentVersion\\Uninstall"
)

type anti struct {
	name  string
	state byte
}

func parseAntiState(state int) byte {
	tmp := fmt.Sprintf("0%x", state)
	var r byte
	if tmp[2:4] == "11" || tmp[2:4] == "01" || tmp[2:4] == "10" {
		r |= enMask
	}
	if tmp[4:] == "00" {
		r |= upMask
	}
	if state == 393472 {
		r = 1
	}
	//fmt.Printf("%d\t%s [%s] [%s] --> %b\n", state, tmp, tmp[2:4], tmp[4:], r)
	return r
}

func parseAntisByClass(class string) []*anti {
	cmd := exec.Command("powershell", fmt.Sprintf(wmiCommand, class))
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	out, err := cmd.CombinedOutput()
	Handle(err)
	lines := strings.Split(string(out), "\n")

	var antis []*anti

	for i := 0; i < len(lines)-1; i += 2 {
		t, _ := strconv.Atoi(strings.TrimSpace(lines[i+1]))
		antis = append(antis, &anti{
			name:  strings.TrimSpace(lines[i]),
			state: parseAntiState(t),
		})
	}

	return antis
}

func condenseAntiList(antis []*anti) map[string]byte {
	so := map[string]byte{}

	for _, v := range antis {
		so[v.name] |= v.state
	}

	return so
}

//AVInfo returns information about installed antivirus products.
func AntiInfo() string {
	antis := append([]*anti{}, parseAntisByClass("AntiVirusProduct")...)
	antis = append(antis, parseAntisByClass("AntiSpywareProduct")...)

	so := condenseAntiList(antis)
	var info string
	for k, v := range so {
		info += k + " - "
		if v&enMask != 0 {
			info += "Enabled"
		} else {
			info += "Disabled"
		}
		info += ", "
		if v&upMask != 0 {
			info += "Updated"
		} else {
			info += "Outdated"
		}
		info += "\n"
	}
	return strings.TrimSpace(info)
}

func SoftwareInfo() string {
	s := []string{}

	uninst, err := registry.OpenKey(registry.LOCAL_MACHINE, softKeys, registry.ENUMERATE_SUB_KEYS)
	Handle(err)
	defer uninst.Close()
	keys, err := uninst.ReadSubKeyNames(0)
	Handle(err)
	for _, v := range keys {
		key, err := registry.OpenKey(uninst, v, registry.READ)
		Handle(err)
		name, _, err := key.GetStringValue("DisplayName")
		if err != nil {
			continue
		}
		key.Close()

		s = append(s, name)
	}

	sort.Strings(s)
	s = RemoveDuplicates(s)
	return strings.Join(s, "\n")
}
