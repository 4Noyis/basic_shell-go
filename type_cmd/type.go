package typecmd

import (
	"fmt"
	"main/models"
	"os"
)

func Type(command models.Cmd) {
	var arg string
	if command.Args != nil {
		arg = command.Args[0]
	} else {
		fmt.Printf("wrong usage")
		os.Exit(1)
	}

	switch arg {
	case "echo":
		fmt.Fprintln(os.Stdout, "echo is a shell builtin")
	case "type":
		fmt.Fprintln(os.Stdout, "type is a shell builtin")
	case "exit":
		fmt.Fprintln(os.Stdout, "exit is a shell builtin")
	default:
		fmt.Fprintln(os.Stdout, "how to use: ")
	}
}
