package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	stdin := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("$ ")

		input, err := stdin.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error reading input:", err)
			os.Exit(1)
		}
		input = strings.TrimRight(input, "\n\r")

		fmt.Printf("%v: command not found", input)

		fmt.Println()
	}
}
