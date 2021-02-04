package main

import (
	"fmt"

	"github.com/Raffy27/Hydra/util"
)

func main() {
	elevated := util.RunningAsAdmin()
	admin := util.IsUserAdmin()
	fmt.Println("Elevated?", elevated)
	fmt.Println("Admin?", admin)
}
