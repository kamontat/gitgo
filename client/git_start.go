package client

import (
	"os"
)

// GitInit exec 'git init'
func GitInit() {
	rawGitCommand("init")
}

// GitIsNotInit will return true if never init before
func GitIsNotInit() bool {
	return !GitIsInit()
}

// GitIsInit will return true if already init
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
