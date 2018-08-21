package model

import "strings"

// CommitMessage is the commit message for save in commit.
type CommitMessage struct {
	Key     string
	Title   string
	Message string
}

// GetKey will try to format the key to right way.
// Otherwise, return normal Key
func (c *CommitMessage) GetKey() string {
	arr := strings.Split(c.Key, ":")
	if len(arr) > 0 {
		s := arr[0]
		return strings.TrimSpace(s)
	}
	return c.Key
}
