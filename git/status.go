package git

// Status will exec git commandline with 'status'
func Status() {
	rawGitCommand("status")
}
