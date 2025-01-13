package main

import (
	"bufio"
	"fmt"
	clearcmd "main/clear_cmd"
	echocmd "main/echo_cmd"
	exitcmd "main/exit_cmd"
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

func run(cmd models.Cmd) {
	switch cmd.Name {
	case "echo":
		echocmd.Echo(cmd)
	case "type":
		typecmd.Type(cmd)
	case "exit":
		exitcmd.Exit(cmd)
	case "clear":
		clearcmd.Clear()
	default:
		//fmt.Fprintln(os.Stdout, cmd.Name+": not found")
		command := exec.Command(cmd.Name, cmd.Args[0:]...)
		command.Stderr = os.Stderr
		command.Stdout = os.Stdout
		err := command.Run()
		if err != nil {
			fmt.Printf("%s: command not found\n", cmd.Name)
		}

	}
}
