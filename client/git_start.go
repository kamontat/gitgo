package client

import (
	"os"
)

func GitInit() {
	rawGitCommand("init")
}

func GitIsNotInit() bool {
	return !GitIsInit()
}

func GitIsInit() bool {
	_, err := os.Stat("./.git")
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return true
}
