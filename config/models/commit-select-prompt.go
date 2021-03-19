package models

import "fmt"

// CommitSelectValue TODO
type CommitSelectValue struct {
	Key  string
	Text string
}

// Format will return format of string.
func (c *CommitSelectValue) Format() string {
	return fmt.Sprintf("%-10s: %s", c.Key, c.Text)
}

// String will return string that show what is it.
func (c *CommitSelectValue) String() string {
	return fmt.Sprintf("commit key=%s, text=%s", c.Key, c.Text)
}

type CommitSelectPromptSetting struct {
	Page       uint8                // for select
	Suggestion bool                 // for select
	Values     []*CommitSelectValue // for select
}

func (c *CommitSelectPromptSetting) List() []string {
	var list []string
	for _, value := range c.Values {
		list = append(list, value.Format())
	}

	return list
}
