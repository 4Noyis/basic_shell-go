package catcmd

import (
	"fmt"
	"os"

	"github.com/4Noyis/basic_shell-go/models"
)

func Cat(command models.Cmd) {
	reset := models.Reset
	red := models.Red

	if len(command.Args) == 0 {
		fmt.Fprintln(os.Stderr, red, "Wrong usage", reset)
	} else {
		for i, fileName := range command.Args {
			content, err := os.ReadFile(fileName)
			if err != nil {
				fmt.Fprintln(os.Stderr, red, err, reset)
				return
			}
			fmt.Println(command.Args[i] + ":")
			mystring := string(content[:])
			fmt.Fprintln(os.Stdout, mystring)

		}
	}

}
