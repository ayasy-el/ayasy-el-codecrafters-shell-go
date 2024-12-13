package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	for {
		fmt.Fprint(os.Stdout, "$ ")

		input, _ := bufio.NewReader(os.Stdin).ReadString('\n')
		cmd := strings.Fields(input)

		switch cmd[0] {
		case "exit":
			code, _ := strconv.Atoi(cmd[1])
			os.Exit(code)
		case "echo":
			fmt.Fprintln(os.Stdout, strings.Join(cmd[1:], " "))
		default:
			fmt.Fprintf(os.Stdout, "%s: command not found\n", cmd[0])
		}
	}
}
