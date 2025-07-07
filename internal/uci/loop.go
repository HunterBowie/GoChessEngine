package uci

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Starts the loop that communicates with the GUI using UCI
func LoopUCI() {
	for {
		message := recieveUCIMessage()

		switch message {
		case "uci":
			uci()
		case "isready":
			isready()
		case "quit":
			return
		default:
			if strings.HasPrefix(message, "position") {
				// pass
			}
			if strings.HasPrefix(message, "go") {

				if strings.Contains(message, "infinite") {
					goInfinite()
				} else {
					goTimed(message)
				}
			}
		}

	}
}

// Recieve and process a message from a chess GUI using UCI
func recieveUCIMessage() string {
	var raw string
	reader := bufio.NewReader(os.Stdin)
	raw, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("ERROR: " + err.Error())
	}
	message := strings.TrimSpace(raw)
	return message
}
