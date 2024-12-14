package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	commands := []string{"exit", "echo", "type", "pwd", "cd"}

	for {
		fmt.Fprint(os.Stdout, "$ ")
		if !scanner.Scan() {
			break
		}

		cmd := parseArgs(scanner.Text())
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
			text := strings.Join(cmd[1:], " ")
			if strings.Contains(text, "'") || strings.Contains(text, "\"") {
				text = strings.ReplaceAll(text, "'", "")
				text = strings.ReplaceAll(text, "\"", "")
				fmt.Println(text)
			} else {
				fmt.Println(text)
			}

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

		case "pwd":
			cwd, _ := os.Getwd()
			fmt.Println(cwd)

		case "cd":
			if strings.HasPrefix(cmd[1], "~") {
				if len(cmd[1]) > 1 && cmd[1][:2] == "~/" {
					cmd[1] = filepath.Join(os.Getenv("HOME"), cmd[1][2:])
				} else if cmd[1] == "~" {
					cmd[1] = os.Getenv("HOME")
				}
			}
			err := os.Chdir(cmd[1])
			if err != nil {
				fmt.Printf("cd: %s: No such file or directory\n", cmd[1])
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

func parseArgs(input string) []string {
	re := regexp.MustCompile(`'([^']*)'|([^' ]+)`)
	matches := re.FindAllStringSubmatch(input, -1)

	var result []string
	for _, match := range matches {
		if match[1] != "" {
			result = append(result, match[1])
		} else if match[2] != "" {
			result = append(result, match[2])
		}
	}

	return result
}
