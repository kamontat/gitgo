package client

// GitPull exec git pull with repo (default as origin), branch (default as master)
func GitPull(withForce bool, repo string, branch []string) error {
	var arr = []string{"pull", "--quiet"}

	// TODO: make repo to optional argument
	if withForce {
		arr = append(arr, "--force")
	}

	arr = append(arr, repo)
	arr = append(arr, branch...)

	return rawGitCommand(arr...)
}
