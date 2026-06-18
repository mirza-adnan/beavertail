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
	// TODO: Uncomment the code below to pass the first stage
	fmt.Print("$ ")

	cmd, err := bufio.NewReader(os.Stdin).ReadString('\n')
	cmd = strings.TrimSuffix(cmd, "\n")
	if (err != nil) {
		fmt.Fprintln(os.Stderr, "Error reading input:", err)
		os.Exit(1)
	}

	fmt.Printf("%v: command not found\n", cmd)
}
