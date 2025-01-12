package main

import (
	"bufio"
	"fmt"
	listls "main/list_ls"
	"os"
)

func main() {

	fmt.Fprint(os.Stdout, "$ ")

	// Wait for user input

	for {
		reader := bufio.NewReader(os.Stdin)
		// reads input from the user until it encounters the newline character
		command, err := reader.ReadString('\n')
		// removes newline character
		command = command[:len(command)-1]

		if err != nil {
			fmt.Fprintln(os.Stderr, "error reading input: ", err)
			os.Exit(1)
		}

		switch {
		case command == "ls":

			listls.Ls()

		case command == "exit":
			os.Exit(0)
		default:
			fmt.Println(command + ": command not found")
			fmt.Fprint(os.Stdout, "$ ")
		}

		fmt.Fprint(os.Stdout, "$ ")
	}

}
