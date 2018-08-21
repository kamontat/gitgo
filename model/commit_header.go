package model

import "fmt"

// CommitHeader is struct of Key and Value, using in list.yaml.
type CommitHeader struct {
	Key   string
	Value string
}

// Format will return format of string.
func (c *CommitHeader) Format() string {
	return fmt.Sprintf("%-10s: %s", c.Key, c.Value)
}

// String will return string that show what is it.
func (c *CommitHeader) String() string {
	return fmt.Sprintf("commit key=%s, value=%s", c.Key, c.Value)
}
