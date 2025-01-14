package cdcmd

import (
	"main/models"
	"os"
)

func Cd(command models.Cmd) {

	if len(command.Args) == 0 {
		os.Chdir("./")
	} else {
		os.Chdir(command.Args[0])
	}

	//dir := command.Args[0]

}
