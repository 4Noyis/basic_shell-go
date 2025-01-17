package catcmd

import (
	"fmt"
	"os"
)

func Cat() {
	content, err := os.ReadFile("main.go")
	mystring := string(content[:])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	fmt.Fprintln(os.Stdout, mystring)
}
