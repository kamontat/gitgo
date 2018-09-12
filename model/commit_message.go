package model

import (
	"fmt"
	"strings"
)

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

func (c *CommitMessage) GetMessage() string {
	if c.Message == "" {
		return fmt.Sprintf("[%s] %s", c.GetKey(), c.Title)
	}
	return fmt.Sprintf("[%s] %s\n%s", c.GetKey(), c.Title, c.Message)
}
