package client

// GitHasRemote check is git has spectify remote
func GitHasRemote(repoName string) bool {
	_, err := rawGitCommandNoLog("remote", "get-url", repoName)
	return err == nil
}

// GitHaveRemote check is any remote exist?
func GitHaveRemote() bool {
	_, err := rawGitCommandNoLog("ls-remote", "--exit-code")
	return err == nil
}

// GitDontHaveRemote check is any remote not exist?
func GitDontHaveRemote() bool {
	return !GitHaveRemote()
}

// GitAddRemote add new remote, must not exist before
func GitAddRemote(repo string, link string) {
	if !GitHasRemote(repo) {
		rawGitCommand("remote", "add", repo, link)
	}
}

// GitRemoveRemote remove exist remote only
func GitRemoveRemote(repo string) {
	if GitHasRemote(repo) {
		rawGitCommand("remote", "remove", repo)
	}
}

// GitReAddRemote re add exist repository
func GitReAddRemote(repo string, link string) {
	GitRemoveRemote(repo)
	GitAddRemote(repo, link)
}
