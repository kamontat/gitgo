package models

import (
	"fmt"
	"regexp"

	"github.com/go-git/go-git/v5/plumbing/object"
)

type ChangeLogMessage struct {
	Scope string
	Text  string
	Hash  string
}

func (c ChangeLogMessage) String() string {
	return fmt.Sprintf("**%s**: %s (%s)", c.Scope, c.Text, c.Hash)
}

type ChangeLog struct {
	Tag    string
	Values map[string][]ChangeLogMessage
}

var commitMessageRegex = regexp.MustCompile(`([a-z]+)\(([a-zA-Z]+)\): (.*)`)

func parser(message string) (string, string, string) {
	array := commitMessageRegex.FindStringSubmatch(message)
	return array[1], array[2], array[3]
}

func NewChangelog(version string, cs []*object.Commit) *ChangeLog {
	var changelog = &ChangeLog{
		Tag:    version,
		Values: make(map[string][]ChangeLogMessage),
	}

	for _, c := range cs {
		hash := c.Hash.String()
		shortHash := hash[0:7]
		key, scope, text := parser(c.Message)

		changelog.Values[key] = append(changelog.Values[scope], ChangeLogMessage{
			Scope: scope,
			Text:  text,
			Hash:  shortHash,
		})
	}

	return changelog
}
