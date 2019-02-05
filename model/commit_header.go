package model

import "fmt"

// CommitHeader is struct of Key and Value, using in list.yaml.
type CommitHeader struct {
	Type   string
	Value string
}

// Format will return format of string.
func (c *CommitHeader) Format() string {
	return fmt.Sprintf("%-10s: %s", c.Type, c.Value)
}

// String will return string that show what is it.
func (c *CommitHeader) String() string {
	return fmt.Sprintf("commit type=%s, value=%s", c.Type, c.Value)
}
