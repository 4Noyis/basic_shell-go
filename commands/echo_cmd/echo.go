package echocmd

import (
	"fmt"
	"main/models"
)

func Echo(command models.Cmd) {

	for i := 0; i < len(command.Args); i++ {

		fmt.Printf("%s aaaaaaaa", command.Args[i])
	}
	fmt.Printf("\n")

}
