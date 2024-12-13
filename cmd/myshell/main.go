package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	for true {
		fmt.Fprint(os.Stdout, "$ ")

		cmd, _ := bufio.NewReader(os.Stdin).ReadString('\n')
		cmd = strings.TrimSpace(cmd)
		switch cmd {
		case "exit 0":
			os.Exit(0)
		default:
			fmt.Printf("%s: command not found\n", cmd)
		}
	}
}
