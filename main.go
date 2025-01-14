package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/4Noyis/basic_shell-go/cmd/app"
	wd "github.com/4Noyis/basic_shell-go/commands/pwd_cmd"
	"github.com/4Noyis/basic_shell-go/models"
)

func main() {
	for {
		wd := wd.WorkingDir()
		fmt.Fprint(os.Stdout, wd+" $ ")
		// Wait for user input
		input, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			fmt.Println(err.Error())
		}
		cmd := models.ParseCmd(input[:len(input)-1])
		app.Run(cmd)
	}
}
