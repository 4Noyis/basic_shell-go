package pwdcmd

import (
	"fmt"
	"os"
)

func Pwd() {
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Fprintln(os.Stdout, pwd)
}
