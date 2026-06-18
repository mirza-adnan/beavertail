package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Ensures gofmt doesn't remove the "fmt" import in stage 1 (feel free to remove this!)
var _ = fmt.Print

func main() {
	var reader = bufio.NewReader(os.Stdin) 

	for {
		fmt.Print("$ ")

		var cmdline, err = reader.ReadString('\n')
		if (err != nil) {
			fmt.Fprintln(os.Stderr, "Error reading input:", err)
			os.Exit(1)
		}

		cmdline = strings.TrimSpace(cmdline)
		var cmdParts = strings.SplitN(cmdline, " ", 2)
		var cmd = cmdParts[0]
		
		if (strings.ToLower(cmd) == "exit") {
			break
		} else if (cmd == "echo") {
			if (len(cmdParts) > 1) {
				fmt.Println(cmdParts[1])
			}
		} else {
			fmt.Printf("%v: command not found\n", cmd)
		}

	}
}
