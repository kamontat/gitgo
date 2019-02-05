// Package model provides the model of repository and commit.
// As long as another model that will be use in gitgo command.
package model

import (
	"github.com/kamontat/gitgo/exception"

	"github.com/kamontat/go-error-manager"

	"github.com/kamontat/go-log-manager"
	survey "gopkg.in/AlecAivazis/survey.v1"
)

// CommitOption is a option of commit
type CommitOption struct {
	Add     bool
	Empty   bool
	Dry     bool
	Message bool
}

// Commit is Commit object of deal with commit things.
type Commit struct {
	throwable *manager.Throwable
	KeyList   *List
}

func (c *Commit) getQuestion() []*survey.Question {
	if !c.KeyList.IsContain() {
		e.ShowAndExit(e.ErrorMessage(e.IsInitial, "No key list for commit"))
	}

	return []*survey.Question{
		{
			Name: "type",
			Prompt: &survey.Select{
				Message:  "Select commit header",
				Options:  c.KeyList.MakeList(),
				Help:     "Header will represent 'one word' key of the commit",
				PageSize: 5,
				VimMode:  true,
			},
			Validate: survey.Required,
		}, {
			Name: "title",
			Prompt: &survey.Input{
				Message: "Enter commit title",
				Help:    "Title will represent one short sentence of the commit",
			},
			Validate: func(val interface{}) error {
				err := survey.Required(val)
				if err == nil {
					err = survey.MaxLength(50)(val)
				}
				return err
			},
		}, {
			Name: "message",
			Prompt: &survey.Multiline{
				Message: "Enter commit message",
				Help:    "Message will represent everything that commit have done",
			},
		},
	}
}

// Commit is action for ask the message from user and call CustomCommit.
func (c *Commit) Commit(key string, option CommitOption) {
	// the questions to ask
	var qs = c.getQuestion()

	if key != "" {
		qs = qs[1:]
	}

	if !option.Message {
		qs = qs[:len(qs)-1]
	}

	om.Log.ToDebug("question list", len(qs))

	answers := CommitMessage{
		Type: key,
	}

	manager.StartResultManager().Save("", survey.Ask(qs, &answers)).IfNoError(func() {
		om.Log.ToDebug("commit type", answers.GetType())
		om.Log.ToDebug("commit title", answers.Title)
		om.Log.ToDebug("commit message", answers.Message)

		c.CustomCommit(answers, option)
	}).IfError(func(t *manager.Throwable) {
		e.ShowAndExit(e.Update(t, e.IsUser))
	})
}

// CustomCommit will run git commit -m "<message>" with the default format.
func (c *Commit) CustomCommit(answers CommitMessage, option CommitOption) {

	var commitMessage = answers.GetMessage()

	var t *manager.Throwable

	args := []string{"commit"}

	if option.Add {
		args = append(args, "-a")
		om.Log.ToVerbose("commit", "with -a flag")
	}

	if option.Empty {
		args = append(args, "--allow-empty")
		om.Log.ToVerbose("commit", "with --allow-empty flag")
	}

	args = append(args, "-m")
	args = append(args, commitMessage)

	om.Log.ToInfo("Commit", commitMessage)
	if !option.Dry {
		t = Git().Exec(args...).Throw()
		e.ShowAndExit(e.Update(t, e.IsCommit))
	}
}
