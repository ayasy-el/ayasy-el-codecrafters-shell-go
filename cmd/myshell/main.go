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
	commands := []string{"exit", "echo", "type", "pwd", "cd"}

	for {
		fmt.Fprint(os.Stdout, "$ ")
		if !scanner.Scan() {
			break
		}

		s := strings.Trim(scanner.Text(), "\r\n")
		cmd, argstr, _ := strings.Cut(s, " ")
		args := parseArgs(argstr)

		if cmd == "" {
			continue
		}

		switch cmd {
		case "exit":
			code := 0
			if len(args) > 0 {
				code, _ = strconv.Atoi(args[0])
			}
			os.Exit(code)

		case "echo":
			fmt.Println(strings.Join(args, " "))

		case "type":
			if len(args) < 1 {
				fmt.Println("type: missing argument")
				continue
			}

			if slices.Contains(commands, args[0]) {
				fmt.Printf("%s is a shell builtin\n", args[0])
			} else if foundCommand := findCommand(args[0]); foundCommand != "" {
				fmt.Printf("%s is %s\n", args[0], foundCommand)
			} else {
				fmt.Printf("%s: not found\n", args[0])
			}

		case "pwd":
			cwd, _ := os.Getwd()
			fmt.Println(cwd)

		case "cd":
			if strings.HasPrefix(args[0], "~") {
				if len(args[0]) > 1 && args[0][:2] == "~/" {
					args[0] = filepath.Join(os.Getenv("HOME"), args[0][2:])
				} else if args[0] == "~" {
					args[0] = os.Getenv("HOME")
				}
			}
			err := os.Chdir(args[0])
			if err != nil {
				fmt.Printf("cd: %s: No such file or directory\n", args[0])
			}

		default:
			if foundCommand := findCommand(cmd); foundCommand != "" {
				execCommand(foundCommand, args)
			} else {
				fmt.Printf("%s: command not found\n", cmd)
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
	cmd := exec.Command(command, args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
	}
}

func parseArgs(argstr string) []string {
	var singleQuote bool
	var doubleQuote bool
	var backslash bool
	var arg string
	var args []string
	for _, r := range argstr {
		switch r {
		case '\'':
			if backslash && doubleQuote {
				arg += "\\"
			}
			if backslash || doubleQuote {
				arg += string(r)
			} else {
				singleQuote = !singleQuote
			}
			backslash = false
		case '"':
			if backslash || singleQuote {
				arg += string(r)
			} else {
				doubleQuote = !doubleQuote
			}
			backslash = false
		case '\\':
			if backslash || singleQuote {
				arg += string(r)
				backslash = false
			} else {
				backslash = true
			}
		case ' ':
			if backslash && doubleQuote {
				arg += "\\"
			}
			if backslash || singleQuote || doubleQuote {
				arg += string(r)
			} else if arg != "" {
				args = append(args, arg)
				arg = ""
			}
			backslash = false
		default:
			if doubleQuote && backslash {
				arg += "\\"
			}
			arg += string(r)
			backslash = false
		}
	}
	if arg != "" {
		args = append(args, arg)
	}
	return args
}
