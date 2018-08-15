package model

import (
	"fmt"

	"github.com/kamontat/go-log-manager"
	survey "gopkg.in/AlecAivazis/survey.v1"
)

type Commit struct {
	repo *Repo
}

type CommitQuestion struct {
	Key     string
	Title   string
	Message string
}

func (c *Commit) getQuestion() []*survey.Question {
	return []*survey.Question{
		{
			Name: "key",
			Prompt: &survey.Select{
				Message:  "Select commit header",
				Options:  []string{"a", "b", "c"},
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
			Prompt: &survey.Editor{
				Message: "Enter commit message",
				Help:    "Message will represent everything that commit have done",
			},
		},
	}
}

func (c *Commit) Commit() {
	// the questions to ask
	var qs = c.getQuestion()

	answers := CommitQuestion{}

	// perform the questions
	err := survey.Ask(qs, &answers)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	om.Log().ToDebug("commit key", answers.Key)
	om.Log().ToDebug("commit title", answers.Title)
	om.Log().ToDebug("commit message", answers.Message)

	var commitMessage string
	if answers.Message == "" {
		commitMessage = fmt.Sprintf("[%s] %s", answers.Key, answers.Title)
	} else {
		commitMessage = fmt.Sprintf("[%s] %s\n%s", answers.Key, answers.Title, answers.Message)
	}

	om.Log().ToVerbose("commit full", commitMessage)
	// Git().Exec("commit", "-m", commitMessage)
}
