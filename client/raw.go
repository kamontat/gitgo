package client

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"os/exec"
)

// rawGitCommand is interface to exec git commandline
func rawGitCommand(arg ...string) error {
	return rawCommand("git", arg...)
}

// rawGitCommand is interface to exec git commandline
func rawGitCommandNoLog(arg ...string) (out string, err error) {
	out, _, err = rawCommandAndReturn("git", arg...)
	return
}

// RawLSCommand for exec ls cli
func RawLSCommand(arg ...string) (str string, err error) {
	str, _, err = rawCommandAndReturn("ls", arg...)
	return
}

func RawOpenCommand(editor string, arg ...string) (err error) {
	return rawCommandWithCustomSTD(editor, os.Stdin, os.Stdout, arg...)
}

func RawOpenEditorCommand(arg ...string) (err error) {
	editor := os.Getenv("EDITOR")
	if editor == "" {
		return errors.New("$EDITOR must be set, before open")
	}
	return RawOpenCommand(editor, arg...)
}

func rawCommand(name string, arg ...string) (err error) {
	out, stderr, err := rawCommandAndReturn(name, arg...)
	if err == nil {
		fmt.Printf(out)
	} else {
		fmt.Printf(stderr)
		// log.Fatal()
	}
	return
}

func rawCommandAndReturn(name string, arg ...string) (strout string, strerr string, err error) {
	var stdout, stderr bytes.Buffer
	// fmt.Println(name, strings.Join(arg, " "))
	cmd := exec.Command(name, arg...)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err = cmd.Run()
	strout = stdout.String()
	strerr = stderr.String()
	return
}

func rawCommandWithCustomSTD(name string, in *os.File, out *os.File, arg ...string) (err error) {
	cmd := exec.Command(name, arg...)
	cmd.Stdin = in
	cmd.Stdout = out

	err = cmd.Run()
	return
}
