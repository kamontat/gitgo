package client

import "fmt"

// GitDelete call 'rm -r .git'
func GitDelete() (err error) {
	err = rawCommand("rm", "-r", ".git")
	if err != nil {
		return
	}

	fmt.Println(".git have been deleted")
	return
}
