package client

import (
	"fmt"
)

func _gitCommit(withAdd bool, msg string) {
	var opt string
	if withAdd {
		opt = "-am"
	} else {
		opt = "-m"
	}
	rawGitCommand("commit", opt, msg)
}

func MakeGitCommitWithText(key string, title string, message string) {
	fmt.Println("Make git commit with text key")
}

func MakeGitCommitWithEmoji(key string, title string, message string) {
	fmt.Println("Make git commit with text emoji")
}

// GitAdd is git add command with arguments
// func GitCommitAsText(title string, arg ...string) {
// 	data := append([]string{"add"}, arg...)
// 	rawGitCommand(data...)
// }
