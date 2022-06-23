package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	if !(len(os.Args) >= 3 && os.Args[len(os.Args)-1] == "-V=full") {
		fmt.Println("TOOLEXEC_IMPORTPATH", os.Getenv("TOOLEXEC_IMPORTPATH"))
		fmt.Println(strings.Join(os.Args[1:], " "))
	}
	cmd := exec.Command(os.Args[1], os.Args[2:]...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	cmd.Run()
	os.Exit(cmd.ProcessState.ExitCode())
}
