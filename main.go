package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
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
