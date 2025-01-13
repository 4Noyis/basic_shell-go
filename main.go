package main

import (
	"bufio"
	"fmt"
	echocmd "main/echo_cmd"
	exitcmd "main/exit_cmd"
	"main/models"
	typecmd "main/type_cmd"
	"os"
)

func main() {
	for {
		fmt.Fprint(os.Stdout, "$ ")
		// Wait for user input
		input, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			fmt.Println(err.Error())
		}
		cmd := models.ParseCmd(input[:len(input)-1])
		run(cmd)
	}
}

func run(cmd models.Cmd) {
	switch cmd.Name {
	case "echo":
		echocmd.Echo(cmd)
	case "type":
		typecmd.Type(cmd)
	case "exit":
		exitcmd.Exit(cmd)
	default:
		fmt.Fprintln(os.Stdout, cmd.Name+": not found")

	}
}
