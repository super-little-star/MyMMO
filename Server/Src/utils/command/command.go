package command

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Run() {
	isRunning := true

	for isRunning == true {

		reader := bufio.NewReader(os.Stdin)
		input, err := reader.ReadString('\r')
		if err != nil {
			fmt.Println("Input err : ", err)
			continue
		}
		cmd := strings.Split(input, " ")
		switch cmd[0] {
		case "exit\r":
			isRunning = false
			break
		case "help\r":
			help()
			break
		default:
			help()
			break

		}
	}
}

func help() {
	s := "Help:\n" +
		"exit --- Stop Game Server\n" +
		"help --- Show Help\n"
	print(s)
}
