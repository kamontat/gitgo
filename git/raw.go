package git

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
)

// rawGitCommand is interface to exec git commandline
func rawGitCommand(arg ...string) {
	var out bytes.Buffer

	cmd := exec.Command("git", arg...)
	cmd.Stdout = &out

	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf(out.String())
}
