package client

func GitDelete() {
	rawCommand("rm", "-r", ".git")
}
