package main

import (
	"bufio"
	"fmt"
	echocmd "main/echo_cmd"
	"main/models"
	typecmd "main/type_cmd"
	"os"
	"os/exec"
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

var cmd models.Cmd

func run(cmd models.Cmd) {
	switch cmd.Name {
	case "echo":
		echocmd.Echo(cmd)
	case "type":
		typecmd.Type(cmd)
	case "exit":
		// handleExit(cmd)
	default:
		command := exec.Command(cmd.Name, cmd.Args...)
		command.Stdout = os.Stdout
		command.Stderr = os.Stderr
		err := command.Run()
		if err != nil {
			fmt.Fprintf(os.Stdout, "%s: command not found\n", cmd.Name)
		}
	}
}
