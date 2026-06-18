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
	for {
		fmt.Print("$ ")

		cmd, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if (err != nil) {
			fmt.Fprintln(os.Stderr, "Error reading input:", err)
			os.Exit(1)
		}

		cmd = strings.TrimSuffix(cmd, "\n")
		fmt.Printf("%v: command not found\n", cmd)
	}

	
}
