package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

func main() {
	var reader = bufio.NewReader(os.Stdin) 
	var builtins = []string{"echo", "exit", "type"}
	var PATH = os.Getenv("PATH")
	var pathList = strings.Split(PATH, ":")

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
		
		if (cmd == "exit") {
			break
		} else if (cmd == "echo") {
			if (len(cmdParts) > 1) {
				fmt.Println(cmdParts[1])
			}
		} else if (cmd == "type") {
			if (len(cmdParts) > 1) {
				if (slices.Contains(builtins, cmdParts[1])) {
					fmt.Printf("%v is a shell builtin\n", cmdParts[1])
				} else {
					var found = false
					for _, path := range pathList {
						var dirFiles, _ = os.ReadDir(path)

						for _, dirFile := range dirFiles {
							if (!dirFile.IsDir() && dirFile.Name() == cmdParts[1]) {
								var fileInfo, _ = os.Stat(fmt.Sprintf("%v/%v", path, dirFile.Name()))
								if (fileInfo.Mode() & 1 > 0) {
									fmt.Printf("%v is %v/%v\n", cmdParts[1], path, cmdParts[1])
									found = true
									break
								}
							}
						}
						if (found) {
							break
						}
					}
					if (!found) {
						fmt.Printf("%v: not found\n", cmdParts[1])
					}
				}
			}
		} else {
			fmt.Printf("%v: command not found\n", cmd)
		}
	}
}
