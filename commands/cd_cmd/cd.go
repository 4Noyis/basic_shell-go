package cdcmd

import (
	"os"

	"github.com/4Noyis/basic_shell-go/models"
)

func Cd(command models.Cmd) {

	if len(command.Args) == 0 {
		os.Chdir("./")
	} else {
		os.Chdir(command.Args[0])
	}

	//dir := command.Args[0]

}
