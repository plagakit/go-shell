package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var builtins map[string]func([]string)

func main() {

	stdin := bufio.NewReader(os.Stdin)
	builtins = map[string]func([]string) {
		"exit": exit,
		"echo": echo,
		"type": type_fn,
	}

	for {
		fmt.Print("$ ")

		input, err := stdin.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error reading input:", err)
			os.Exit(1)
		}
		
		args, valid_args := parse_input(input)
		if !valid_args {
			continue
		}

		if cmd, ok := builtins[args[0]]; ok {
			cmd(args)
		} else {
			fmt.Printf("%v: command not found", args[0])
		}

		fmt.Println()
	}
}

func parse_input(input string) ([]string, bool) {
	input = strings.TrimRight(input, "\n\r")
	input = strings.TrimSpace(input)

	if len(input) == 0 {
		return nil, false
	}

	args := strings.Split(input, " ")

	if len(args) == 0 {
		return nil, false
	} else {
		return args, true
	}
}

func exit(args []string) {
	os.Exit(0)
}

func echo(args []string) {
	fmt.Print(strings.Join(args[1:], " "))
}

func type_fn(args []string) {
	cmd := args[1]
	if _, is_builtin := builtins[cmd]; is_builtin {
		fmt.Printf("%v is a shell builtin", cmd)
	} else {
		fmt.Printf("%v: not found", cmd)
	}
}