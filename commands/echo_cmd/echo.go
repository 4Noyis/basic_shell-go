package echocmd

import (
	"fmt"
	"strings"

	"github.com/4Noyis/basic_shell-go/models"
)

func Echo(command models.Cmd) {

	for i := 0; i < len(command.Args); i++ {

		arg := strings.Trim(command.Args[i], `"'`)
		fmt.Printf("%s ", arg)
	}
	fmt.Printf("\n")

}
