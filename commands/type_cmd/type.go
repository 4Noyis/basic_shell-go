package typecmd

import (
	"fmt"
	"os"

	"github.com/4Noyis/basic_shell-go/models"
)

func Type(command models.Cmd) {

	reset := models.Reset
	red := models.Red

	if len(command.Args) == 0 {
		fmt.Fprintln(os.Stdout, "Error: No arguments provided")
		return
	} else {
		for _, arg := range command.Args {
			switch arg {
			case "echo":
				fmt.Fprintln(os.Stdout, "echo is a shell builtin")
			case "type":
				fmt.Fprintln(os.Stdout, "type is a shell builtin")
			case "exit":
				fmt.Fprintln(os.Stdout, "exit is a shell builtin")
			case "pwd":
				fmt.Fprintln(os.Stdout, "pwd is a shell builtin")
			case "ls":
				fmt.Fprintln(os.Stdout, "ls is a shell builtin")
			default:
				fmt.Fprintln(os.Stderr, red+arg+": not found"+reset)
			}
		}
	}

}
