package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	for {
		fmt.Fprint(os.Stdout, "$ ")

		input, _ := bufio.NewReader(os.Stdin).ReadString('\n')
		cmd := strings.Fields(input)

		commands := []string{"exit", "echo", "type"}

		switch cmd[0] {
		case "exit":
			code, _ := strconv.Atoi(cmd[1])
			os.Exit(code)
		case "echo":
			fmt.Fprintln(os.Stdout, strings.Join(cmd[1:], " "))
		case "type":
			if slices.Contains(commands, cmd[1]) {
				fmt.Fprintf(os.Stdout, "%s is a shell builtin\n", cmd[1])
			} else {
				fmt.Fprintf(os.Stdout, "%s: not found\n", cmd[1])
			}
		default:
			fmt.Fprintf(os.Stdout, "%s: command not found\n", cmd[0])
		}
	}
}
