package models

import (
	"errors"
	"fmt"
	"strings"
)

type CommitMessage struct {
	Key     string
	Scope   string
	Title   string
	Message string
}

func (c *CommitMessage) Formatted() (string, error) {
	if c.Key == "" && (c.Title == "" || c.Message == "") {
		return "", errors.New("key and title/message cannot both empty")
	}

	if c.Title == "" && c.Message == "" {
		return "", errors.New("required at least either title or message defined")
	}

	key := strings.TrimSpace(c.Key)
	scope := strings.TrimSpace(c.Scope)
	title := strings.TrimSpace(c.Title)
	message := strings.TrimSpace(c.Message)

	var result strings.Builder

	if key != "" {
		if scope != "" {
			result.WriteString(fmt.Sprintf("%s(%s): ", key, scope))
		} else {
			result.WriteString(fmt.Sprintf("%s: ", key))
		}
	}

	if title != "" {
		result.WriteString(title)
	}

	if message != "" {
		result.WriteString(fmt.Sprintf("\n\n%s", message))
	}

	return result.String(), nil
}

func (c *CommitMessage) String() string {
	msg, err := c.Formatted()
	if err != nil {
		return err.Error()
	}

	return msg
}
