package listls

import (
	"fmt"
	"os"
)

func Ls() {

	// flagler ile özellikler eklenicek -l -a ...
	// gizli dosyaları gizlemek için başında . var mı yok mu kontrol edilecek
	// fazla dosya konumu verilirse hata döndürücek

	if len(os.Args) < 2 { // iki değil sadece argüman verilmediyse bulunduğun yeri okuyacak
		file, _ := os.ReadDir("./")
		for _, file := range file {
			fmt.Print(file.Name(), " ")
		}
	} else {
		file, _ := os.ReadDir(os.Args[1]) // dosya konumu alınırken son arg alınıcak 1 değil (flaglarden)
		for _, file := range file {
			fmt.Print(file.Name(), " ")

		}
	}

	fmt.Println()
}
