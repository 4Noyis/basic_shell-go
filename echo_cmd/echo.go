package echocmd

import (
	"fmt"
	"main/models"
)

func Echo(command models.Cmd) {

	for i := 0; i < len(command.Args); i++ {

		fmt.Printf("%s ", command.Args[i])
	}
	fmt.Printf("\n")

}
