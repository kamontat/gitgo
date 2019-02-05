package model

import "strings"

// CommitMessage is the commit message for save in commit.
type CommitMessage struct {
	Type       string
	Scope      string
	Title      string
	HasMessage bool
	Message    string
}

// GetType will try to format the type to right way.
// Otherwise, return normal Type
func (c *CommitMessage) GetType() string {
	arr := strings.Split(c.Type, ":")
	if len(arr) > 0 {
		s := arr[0]
		return strings.TrimSpace(s)
	}
	return c.Type
}
