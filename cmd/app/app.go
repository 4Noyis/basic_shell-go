package app

import (
	"fmt"
	cdcmd "main/commands/cd_cmd"
	clearcmd "main/commands/clear_cmd"
	echocmd "main/commands/echo_cmd"
	exitcmd "main/commands/exit_cmd"
	pwdcmd "main/commands/pwd_cmd"
	typecmd "main/commands/type_cmd"
	"main/models"
	"os"
	"os/exec"
)

func Run(cmd models.Cmd) {
	switch cmd.Name {
	case "echo":
		echocmd.Echo(cmd)
	case "type":
		typecmd.Type(cmd)
	case "exit":
		exitcmd.Exit()
	case "clear":
		clearcmd.Clear()
	case "pwd":
		pwdcmd.Pwd()
	case "cd":
		cdcmd.Cd(cmd)
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
