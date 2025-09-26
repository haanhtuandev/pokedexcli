package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	cf := Config{Next: url, Previous: ""}

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		cmd_list := handleCommand(scanner.Text())
		cmd := cmd_list[0]
		value, ok := commandMap[cmd]
		if !ok {
			fmt.Println("Unknown command")
		} else {
			if len(cmd_list) < 2 {
				value.callback(&cf)
			} else {
				value.callback(&cf, cmd_list[1])
			}

		}

	}
}
