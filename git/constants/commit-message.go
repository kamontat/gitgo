package constants

import "github.com/kamontat/gitgo/git/models"

var InitialCommitMessage = &models.CommitMessage{
	Key:   "feat",
	Scope: "init",
	Title: "initial new project",
}
