// Package model provides the model of repository and commit.
// As long as another model that will be use in gitgo command.
package model

import (
	e "github.com/kamontat/gitgo/exception"
	manager "github.com/kamontat/go-error-manager"
	om "github.com/kamontat/go-log-manager"
)

// CommitOptions is a option to run git commit
type CommitOptions struct {
	Dry   bool
	Empty bool
}

// Commit is Commit object of deal with commit things.
type Commit struct {
	throwable *manager.Throwable
}

// Make is a command execute git commit -m "<message>"
func (c *Commit) Make(answers CommitMessage, options CommitOptions) {
	message := answers.GetMessage()

	args := []string{"commit"}
	if options.Empty {
		args = append(args, "--allow-empty")
	}

	args = append(args, "-m")
	args = append(args, message)

	if options.Dry {
		om.Log.ToInfo("Git commit", message)
	} else {
		c.throwable = Git().Exec(args...).Throw()
		e.ShowAndExit(e.Update(c.throwable, e.IsCommit))
	}
}
