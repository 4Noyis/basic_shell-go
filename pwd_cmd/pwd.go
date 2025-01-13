package pwdcmd

import (
	"fmt"
	"os"
	"strings"
)

func Pwd() {
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Fprintln(os.Stdout, pwd)
}

func HeaderPwd() string {
	wd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}

	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error getting home directory:", err)
	}

	if strings.HasPrefix(wd, homeDir) {
		wd = strings.Replace(wd, homeDir, "~", 1)

	}
	return wd
}
