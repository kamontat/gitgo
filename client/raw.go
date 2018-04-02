package client

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
)

// rawGitCommand is interface to exec git commandline
func rawGitCommand(arg ...string) error {
	return rawCommand("git", arg...)
}

// RawLSCommand for exec ls cli
func RawLSCommand(arg ...string) (str string, err error) {
	str, _, err = rawCommandAndReturn("ls", arg...)
	return
}

func rawCommand(name string, arg ...string) (err error) {
	out, stderr, err := rawCommandAndReturn(name, arg...)
	if err == nil {
		fmt.Printf(out)
	} else {
		log.Fatal(stderr)
	}
	return
}

func rawCommandAndReturn(name string, arg ...string) (strout string, strerr string, err error) {
	var stdout, stderr bytes.Buffer

	fmt.Println(name, arg)
	cmd := exec.Command(name, arg...)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err = cmd.Run()
	strout = stdout.String()
	strerr = stderr.String()
	return
}
