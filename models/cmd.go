package models

import "strings"

// Cmd is an exported struct to represent a command
type Cmd struct {
	Name string   // Exported field (starts with uppercase)
	Args []string // Exported field (starts with uppercase)
}

func ParseCmd(str string) Cmd {
	parts := strings.Fields(str)
	name := parts[0]

	var args []string
	if parts[1:] != nil {
		args = parts[1:]
	} else {
		args = nil
	}

	return Cmd{name, args}
}
