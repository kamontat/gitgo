package client

// ListTag list every tag in stdout
func ListTag() {
	rawGitCommand("git", "tag")
}

// SetTag set input as git tag
func SetTag(tag string) error {
	return rawGitCommand("git", "tag", tag)
}
