package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/tenntenn/scrap/okuribito/replacer"
)

func main() {

	var args []string
	for _, arg := range os.Args[2:] {
		if strings.HasSuffix(arg, ".go") {
			newFile, err := replacer.Replace(arg)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(1)
			}
			args = append(args, newFile)
		} else {
			args = append(args, arg)
		}

	}

	// output log
	//os.WriteFile("log.txt", []byte(fmt.Sprintln(os.Args)), 0o644)
	//os.WriteFile("log.txt", []byte(fmt.Sprintln(args)), 0o644)

	cmd := exec.Command(os.Args[1], args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	cmd.Run()
	os.Exit(cmd.ProcessState.ExitCode())
}
