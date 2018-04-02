package client

// func GitInitPush(repo string, branch string) error {
// 	return nil
// }

// func GitPush(repo string, branch []string) error {
// 	return _gitPush(false, false, repo, branch)
// }

// func GitForcePush(repo string, branch []string) error {
// 	return _gitPush(true, false, repo, branch)
// }

func GitPush(withForce bool, repo string, branch []string) error {
	return GitSetupPush(withForce, false, repo, branch)
}

func GitSetupPush(withForce bool, withUpstream bool, repo string, branch []string) error {
	var arr = []string{"push"}

	// TODO: make repo to optional argument
	if withForce {
		arr = append(arr, "--force")
	}

	if withUpstream {
		arr = append(arr, "--set-upstream")
	}

	arr = append(arr, branch...)

	return rawGitCommand(arr...)
}
