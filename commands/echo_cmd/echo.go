package echocmd

import (
	"fmt"

	"github.com/4Noyis/basic_shell-go/models"
)

func Echo(command models.Cmd) {

	for i := 0; i < len(command.Args); i++ {

		fmt.Printf("%s", command.Args[i])
	}
	fmt.Printf("\n")

}
