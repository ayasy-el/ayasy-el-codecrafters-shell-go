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
	var result []string
	// Regex to match single-quoted, double-quoted, or unquoted words
	re := regexp.MustCompile(`'([^']*)'|"([^"]*)"|(\S+)`)
	matches := re.FindAllStringSubmatch(input, -1)

	for _, match := range matches {
		var arg string
		if match[1] != "" { // Single-quoted
			arg = match[1]
		} else if match[2] != "" { // Double-quoted
			arg = match[2]
		} else if match[3] != "" { // Unquoted
			arg = match[3]
			arg = processEscapes(arg)
		}

		result = append(result, arg)
	}

	return result
}

func processEscapes(input string) string {
	var result strings.Builder
	escaped := false

	for i := 0; i < len(input); i++ {
		char := input[i]

		if escaped {
			result.WriteByte(char)
			escaped = false
			continue
		}

		if char == '\\' {
			escaped = true
			continue
		}

		result.WriteByte(char)
	}

	return result.String()
}
