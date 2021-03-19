package models

import "fmt"

// CommitTypeSetting TODO
type CommitTypeSetting struct {
	Enabled  bool
	Required bool
	Prompt   *CommitPrompt
}

func (c *CommitTypeSetting) String() string {
	enabled := "disabled"
	if c.Enabled {
		enabled = "enabled"
	}

	required := "optional"
	if c.Required {
		required = "required"
	}

	promptName := "unknown"
	if c.Prompt.Select != nil {
		promptName = "select"
	} else if c.Prompt.Input != nil {
		promptName = "input"
	}

	return fmt.Sprintf("%s (%s, %s)", promptName, enabled, required)
}
