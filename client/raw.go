package client

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
)

// rawGitCommand is interface to exec git commandline
func rawGitCommand(arg ...string) {
	rawCommand("git", arg...)
}

// RawLSCommand for exec ls cli
func RawLSCommand(arg ...string) string {
	return rawCommandAndReturn("ls", arg...)
}

func rawCommand(name string, arg ...string) {
	fmt.Printf(rawCommandAndReturn(name, arg...))
}

func rawCommandAndReturn(name string, arg ...string) string {
	var out bytes.Buffer

	cmd := exec.Command(name, arg...)
	cmd.Stdout = &out

	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}

	return out.String()
}
