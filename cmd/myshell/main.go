package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"strconv"
	"strings"
)

func main() {
	for {
		fmt.Fprint(os.Stdout, "$ ")

		input, _ := bufio.NewReader(os.Stdin).ReadString('\n')
		cmd := strings.Fields(input)
		path := strings.Split(os.Getenv("PATH"), ":")
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
				break
			}

			found := false
			for _, p := range path {
				fullpath := filepath.Join(p, cmd[1])
				if _, err := os.Stat(fullpath); err == nil {
					found = true
					fmt.Printf("%s is %s\n", cmd[1], fullpath)
					break
				}
			}

			if !found {
				fmt.Fprintf(os.Stdout, "%s: not found\n", cmd[1])
			}
		default:
			fmt.Fprintf(os.Stdout, "%s: command not found\n", cmd[0])
		}
	}
}
