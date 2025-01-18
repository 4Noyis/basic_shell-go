package app

import (
	"fmt"
	"os"
	"os/exec"

	catcmd "github.com/4Noyis/basic_shell-go/commands/cat_cmd"
	cdcmd "github.com/4Noyis/basic_shell-go/commands/cd_cmd"
	clearcmd "github.com/4Noyis/basic_shell-go/commands/clear_cmd"
	echocmd "github.com/4Noyis/basic_shell-go/commands/echo_cmd"
	exitcmd "github.com/4Noyis/basic_shell-go/commands/exit_cmd"
	listcmd "github.com/4Noyis/basic_shell-go/commands/list_cmd"
	pwdcmd "github.com/4Noyis/basic_shell-go/commands/pwd_cmd"
	typecmd "github.com/4Noyis/basic_shell-go/commands/type_cmd"
	"github.com/4Noyis/basic_shell-go/models"
)

func Run(cmd models.Cmd) {
	reset := models.Reset
	red := models.Red

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
	case "ls":
		listcmd.Ls(cmd)
	case "cat":
		catcmd.Cat(cmd)
	default:
		//fmt.Fprintln(os.Stdout, cmd.Name+": not found")
		command := exec.Command(cmd.Name, cmd.Args[0:]...)
		command.Stderr = os.Stderr
		command.Stdout = os.Stdout
		err := command.Run()
		if err != nil {
			fmt.Fprintln(os.Stderr, red, cmd.Name, ": command not found", reset)
		}

	}
}
