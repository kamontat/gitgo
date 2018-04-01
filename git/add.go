package git

// AddAll is git add command with --all flag
func AddAll() {
	rawGitCommand("add", "--all")
}
