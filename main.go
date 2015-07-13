package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

var version = "0.1.0"
var helpMsg = `NAME:
    ext - A convention for command extensions
USAGE:
    ext commands...
`

func main() {
	defer rescue()

	if len(os.Args) < 2 {
		fmt.Println(helpMsg)
		os.Exit(1)
	}

	switch os.Args[1] {
	case "-h", "--help":
		fmt.Println(helpMsg)
		os.Exit(0)
	case "-v", "--version":
		fmt.Println(version)
		os.Exit(0)
	}

	extArgs, err := LookupExtCmd(os.Args[1:])

	if err != nil {
		fmt.Println(err)
	}

	extCmd := exec.Command(extArgs[0], extArgs[1:]...)
	extCmd.Stdin = os.Stdin
	extCmd.Stdout = os.Stdout
	extCmd.Stderr = os.Stderr
	extCmd.Run()
}

func LookupExtCmd(args []string) ([]string, error) {
	var err error
	for i := len(args); i > 0; i-- {
		extCmd := strings.Join(args[0:i], "-")
		bin, err := exec.LookPath(extCmd)

		if err == nil {
			extArgs := []string{bin}
			extArgs = append(extArgs, args[i:]...)
			return extArgs, nil
		}
	}
	return nil, err
}

func rescue() {
	if err := recover(); err != nil {
		args := strings.Join(os.Args, " ")
		fmt.Printf("Command Failed: %s\n%s\n", args, err)
		os.Exit(1)
	}
}
