package main

import (
	"bufio"
	"fmt"
	"main/cmd/app"
	wd "main/commands/pwd_cmd"
	"main/models"
	"os"
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
