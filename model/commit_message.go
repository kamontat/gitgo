package model

import (
	"fmt"
	"strings"
)

// CommitMessage is the commit message for save in commit.
type CommitMessage struct {
	Type    string
	Title   string
	Message string
}

// GetType will try to format the key to right way.
// Otherwise, return normal Key
func (c *CommitMessage) GetType() string {
	arr := strings.Split(c.Type, ":")
	if len(arr) > 0 {
		s := arr[0]
		return strings.TrimSpace(s)
	}
	return c.Type
}

func (c *CommitMessage) GetMessage() string {
	if c.Message == "" {
		return fmt.Sprintf("[%s] %s", c.GetType(), c.Title)
	}
	return fmt.Sprintf("[%s] %s\n%s", c.GetType(), c.Title, c.Message)
}
