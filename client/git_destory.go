package client

// GitDelete call 'rm -r .git'
func GitDelete() {
	rawCommand("rm", "-r", ".git")
}
