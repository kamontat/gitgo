package client

// Status will exec git commandline with 'status'
func Status() {
	rawOpenCommand("git", "status")
}
