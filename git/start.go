package git

import (
	"os"
)

func Init() {
	rawGitCommand("init")
}

func IsNotInit() bool {
	return !IsInit()
}

func IsInit() bool {
	_, err := os.Stat("./.git")
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return true
}
