package list_cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/4Noyis/basic_shell-go/models"
)

func Ls(command models.Cmd) {

	// flagler ile özellikler eklenicek -l -a ...
	// gizli dosyaları gizlemek için başında . var mı yok mu kontrol edilecek
	// fazla dosya konumu verilirse hata döndürücek

	if len(command.Args) == 0 {
		file, _ := os.ReadDir("./")
		wd, err := os.Getwd()
		if err != nil {
			fmt.Println("Error getting working directory:", err)
			return
		}
		wd = filepath.Base(wd)
		fmt.Println(wd + ": ")
		for _, file := range file {
			if file.Name()[0] == '.' { // checking for secret files
				continue
			} else {

				fmt.Fprintln(os.Stdout, "   └", file.Name())
			}

		}
	} else {
		lastArg := command.Args[len(command.Args)-1]
		file, _ := os.ReadDir(lastArg)
		for _, file := range file {
			fmt.Fprint(os.Stdout, file.Name()+"   ")
		}
	}

	fmt.Println()
}

// func includeAll() {

// }
