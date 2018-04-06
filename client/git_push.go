package client

// GitPush call GitSetupPush with NO set-upstream option
func GitPush(withForce bool, repo string, branch []string) error {
	return GitSetupPush(withForce, false, repo, branch)
}

// GitSetupPush setup for git push, run first time only
//
// call with set upstream (optional)
// and force option (optional)
func GitSetupPush(withForce bool, withUpstream bool, repo string, branch []string) error {
	var arr = []string{"push"}

	// TODO: make repo to optional argument
	if withForce {
		arr = append(arr, "--force")
	}

	if withUpstream {
		arr = append(arr, "--set-upstream")
	}

	arr = append(arr, repo)
	arr = append(arr, branch...)

	return rawGitCommand(arr...)
}
