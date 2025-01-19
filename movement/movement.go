package movement

import (
	"fmt"
	"os"
)

func Move() {

	buf := make([]byte, 3) // Buffer to read arrow key sequences
	for {
		_, err := os.Stdin.Read(buf)
		if err != nil {
			fmt.Println("Error reading input:", err)
			break
		}

		// Handle arrow keys
		if buf[0] == 27 && buf[1] == 91 { // Arrow key prefix
			switch buf[2] {
			case 65: // Up Arrow
				fmt.Print("\033[A")
			case 66: // Down Arrow
				fmt.Print("\033[B")
			case 67: // Right Arrow
				fmt.Print("\033[C")
			case 68: // Left Arrow
				fmt.Print("\033[D") // Move cursor left
			}
		}
	}
}
