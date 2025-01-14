package list_cmd

import (
	"fmt"
	"os"

	"github.com/4Noyis/basic_shell-go/models"
)

func Ls(command models.Cmd) {

	// flagler ile özellikler eklenicek -l -a ...
	// gizli dosyaları gizlemek için başında . var mı yok mu kontrol edilecek
	// fazla dosya konumu verilirse hata döndürücek

	if len(command.Args) == 0 { // iki değil sadece argüman verilmediyse bulunduğun yeri okuyacak
		file, _ := os.ReadDir("./")
		for _, file := range file {
			if file.Name()[0] == '.' {
				continue
			} else {
				fmt.Print(file.Name(), " ")
			}

		}
	} else {
		lastArg := command.Args[len(command.Args)-1]
		file, _ := os.ReadDir(lastArg)
		for _, file := range file {
			fmt.Print(file.Name(), " ")

		}
	}

	fmt.Println()
}

// func includeAll() {

// }
