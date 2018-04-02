package client

// GitAddAll is git add command with --all flag
func GitAddAll() {
	rawGitCommand("add", "--all")
}

// GitAdd is git add command with arguments
func GitAdd(arg ...string) {
	data := append([]string{"add"}, arg...)
	rawGitCommand(data...)
}
