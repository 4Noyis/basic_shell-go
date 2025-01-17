package list_cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/4Noyis/basic_shell-go/models"
)

func Ls(command models.Cmd) {
	magenta := models.Magenta
	reset := models.Reset
	red := models.Red
	// flagler ile özellikler eklenicek -l -a ...

	if len(command.Args) == 0 {
		file, _ := os.ReadDir("./")
		run(file, "", magenta, reset)
	} else {
		for _, dirName := range command.Args {
			files, err := os.ReadDir(dirName)
			if err != nil {
				fmt.Fprint(os.Stdout, red, err, reset+"\n")
				return
			}

			run(files, dirName, magenta, reset)
		}
	}

	fmt.Println()
}

func run(file []os.DirEntry, lastArg string, color string, reset string) {
	wd, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting working directory:", err)
		return
	}

	if lastArg != "" {
		wd = filepath.Base(wd)
		fmt.Println(wd + "/" + lastArg + ": ")
	} else {
		wd = filepath.Base(wd)
		fmt.Println(wd + ": ")
	}

	for _, file := range file {
		if file.Name()[0] == '.' { // checking for secret files
			continue
		} else {

			if file.IsDir() {
				fmt.Fprintln(os.Stdout, color+"   └", file.Name()+reset)
			} else {
				fmt.Fprintln(os.Stdout, "   └", file.Name())
			}

		}
	}
}
