package cdcmd

import (
	"fmt"
	"os"

	"github.com/4Noyis/basic_shell-go/models"
)

func Cd(command models.Cmd) {
	reset := models.Reset
	red := models.Red

	if len(command.Args) == 0 {
		os.Chdir("./")
	} else if len(command.Args) == 1 {
		err := os.Chdir(command.Args[0])
		if err != nil {
			fmt.Fprintln(os.Stderr, red, err, reset)
			return
		}
	} else {
		fmt.Fprintln(os.Stderr, red+" Syntax error: To many arguments"+reset)
	}

}
