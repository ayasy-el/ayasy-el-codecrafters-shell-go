package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"slices"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	commands := []string{"exit", "echo", "type"}

	for {
		fmt.Fprint(os.Stdout, "$ ")
		if !scanner.Scan() {
			break
		}

		cmd := strings.Fields(scanner.Text())
		if len(cmd) == 0 {
			continue
		}

		switch cmd[0] {
		case "exit":
			code := 0
			if len(cmd) > 1 {
				code, _ = strconv.Atoi(cmd[1])
			}
			os.Exit(code)
		case "echo":
			fmt.Println(strings.Join(cmd[1:], " "))
		case "type":
			if len(cmd) < 2 {
				fmt.Println("type: missing argument")
				continue
			}

			if slices.Contains(commands, cmd[1]) {
				fmt.Printf("%s is a shell builtin\n", cmd[1])
			} else if foundCommand := findCommand(cmd[1]); foundCommand != "" {
				fmt.Printf("%s is %s\n", cmd[1], foundCommand)
			} else {
				fmt.Printf("%s: not found\n", cmd[1])
			}
		default:
			if foundCommand := findCommand(cmd[0]); foundCommand != "" {
				execCommand(foundCommand, cmd)
			} else {
				fmt.Printf("%s: command not found\n", cmd[0])
			}
		}
	}
}

func findCommand(command string) string {
	pathDirs := strings.Split(os.Getenv("PATH"), ":")
	for _, dir := range pathDirs {
		fullPath := filepath.Join(dir, command)
		if _, err := os.Stat(fullPath); err == nil {
			return fullPath
		}
	}
	return ""
}

func execCommand(command string, args []string) {
	cmd := exec.Command(command, args[1:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
	}
}
