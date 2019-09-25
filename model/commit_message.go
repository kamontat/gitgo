package model

import (
	"fmt"
	"strings"
)

// CommitMessage is the commit message for save in commit.
type CommitMessage struct {
	Type    string
	Scope   string
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

// GetScope will try to format the key to right way.
// Otherwise, return normal Key
func (c *CommitMessage) GetScope() string {
	arr := strings.Split(c.Scope, ":")
	if len(arr) > 0 {
		s := arr[0]
		return strings.TrimSpace(s)
	}

	return c.Type
}

// GetMessage will return formatted commit message
func (c *CommitMessage) GetMessage() string {
	scope := c.GetScope()
	if scope != "" {
		scope = "(" + scope + ")"
	}

	// Add message to commit message
	if c.Message != "" {
		return fmt.Sprintf("%s%s: %s\n\n%s", c.GetType(), scope, c.Title, c.Message)
	}

	// generate without message
	return fmt.Sprintf("%s%s: %s", c.GetType(), scope, c.Title)
}
