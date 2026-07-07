package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"slices"
	"strings"
)

var builtins = []string{"echo", "exit", "type"}

func main() {
	var reader = bufio.NewReader(os.Stdin) 
	var pathList = getPathList()

	for {
		fmt.Print("$ ")

		var cmdline, err = reader.ReadString('\n')  // includes the newline character
		if (err != nil) {
			fmt.Fprintln(os.Stderr, "Error reading input:", err)
			os.Exit(1)
		}

		cmdline = strings.TrimSpace(cmdline)

		if (cmdline == "") {
			continue
		} else if (cmdline == "exit") {
			break
		}

		executeCommand(cmdline, &pathList)
	}
}

func getPathList() []string {
	return strings.Split(os.Getenv("PATH"), ":")
}

func parseCommand(cmdline string) (string, string) {
	cmdline = strings.TrimSpace(cmdline)
	var parts = strings.SplitN(cmdline," ", 2)
	var cmd = parts[0]
	var args string
	if (len(parts) > 1) {
		args = parts[1]
	}

	return cmd, args
}

func executeCommand(cmdline string, pathList *[]string) {
	var cmd, args = parseCommand(cmdline)
	
	switch cmd {
	case "echo":
		echoCommand(args)
	case "type":
		typeCommand(args, pathList)
	default:
		var filePath = findBinary(cmd, pathList)
		if filePath != "" {
			if isExecutable(filePath) {
				var argsParts = strings.Fields(args)  // using Fields instead because split returns [""] for empty string
				
				var cmdExec = exec.Command(cmd, argsParts...)
				var out, _ = cmdExec.CombinedOutput()

				fmt.Printf("%s", out)
			} else {
				fmt.Printf("%v: Permission denied", filePath)
			}
		} else {
			fmt.Printf("%v: command not found\n", cmd)
		}
	}

}

func echoCommand(args string) {
	if (args != "") {
		fmt.Println(args)
	}
}

func typeCommand(args string, pathList *[]string) {
	if (args == "") {
		return 
	}

	var parts = strings.Split(args, " ")
	if (len(parts) > 1) {
		fmt.Println("type: too many arguments")
	}

	if (slices.Contains(builtins, args)) {
		fmt.Printf("%v is a shell builtin\n", args)
	} else {
		var filePath = findBinary(args, pathList)
		if (filePath == "") {
			fmt.Printf("%v: not found\n", args)
		} else {
			fmt.Printf("%v is %v\n", args, filePath)
		}
	}
}

func findBinary(name string, pathList *[]string) string {
	for _, path := range(*pathList) {
		var contents, _ = os.ReadDir(path)
		
		for _, f := range(contents) {
			if (!f.IsDir() && f.Name() == name) {
				var filePath = fmt.Sprintf("%v/%v", path, name)
				if (isExecutable(filePath)) {
					return filePath
				}
			}
		}
	}

	return ""
}

func isExecutable(filePath string) bool {
	var fileInfo, err = os.Stat(filePath)
	if (err != nil) {
		return false
	}

	return fileInfo.Mode()&1 != 0
}