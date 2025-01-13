package exitcmd

import (
	"main/models"
	"os"
)

func Exit(command models.Cmd) {
	os.Exit(0)
}
