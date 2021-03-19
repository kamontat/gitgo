package models

import "fmt"

type CommitSetting struct {
	Key     *CommitTypeSetting
	Scope   *CommitTypeSetting
	Title   *CommitTypeSetting
	Message *CommitTypeSetting
}

func (c *CommitSetting) EnabledMessage() {
	c.Message.Enabled = true
}

func (c *CommitSetting) String() string {
	return fmt.Sprintf("\n    key: %s\n    scope: %s\n    title: %s\n    message: %s\n", c.Key, c.Scope, c.Title, c.Message)
}
