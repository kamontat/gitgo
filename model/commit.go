// Package model provides the model of repository and commit.
// As long as another model that will be use in gitgo command.
package model

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/kamontat/gitgo/exception"

	"github.com/kamontat/go-error-manager"

	"github.com/spf13/viper"

	"github.com/kamontat/go-log-manager"
	survey "gopkg.in/AlecAivazis/survey.v1"
)

// Commit is Commit object of deal with commit things.
type Commit struct {
	throwable *manager.Throwable
	list      []CommitHeader
}

func (c *Commit) listHeaderOptions() *manager.ResultWrapper {
	var list []string
	for _, commits := range c.list {
		list = append(list, commits.Format())
	}
	om.Log.ToVerbose("commit keys size", len(list))
	wrap := manager.WrapNil()
	if len(list) > 0 {
		wrap = manager.Wrap(list)
	}

	return wrap
}

// CanCommit mean this commit contain no errors
func (c *Commit) CanCommit() bool {
	return len(c.list) > 0
}

func (c *Commit) getQuestion() *manager.ResultWrapper {
	return c.listHeaderOptions().UnwrapNext(func(i interface{}) interface{} {
		return []*survey.Question{
			{
				Name: "key",
				Prompt: &survey.Select{
					Message:  "Select commit header",
					Options:  i.([]string),
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
	})
}

// LoadList will initial new list of Header.
func (c *Commit) LoadList(vp *viper.Viper) *Commit {
	if vp == nil {
		om.Log.ToDebug("commit list", "viper is nil, cannot load list")
		return c
	}

	// reset list
	c.list = []CommitHeader{}

	om.Log.ToDebug("commit list", "load commit list from "+vp.ConfigFileUsed())
	return c.MergeList(vp)
}

// MergeList will merge current list to the new ones.
func (c *Commit) MergeList(vp *viper.Viper) *Commit {
	if vp == nil {
		om.Log.ToWarn("Pre commit", "viper is nil, cannot merge commit list")
		return c
	}
	if vp.Get("list") == nil {
		om.Log.ToWarn("Pre commit", "cannot get commit list")
		return c
	}

	if c.list == nil {
		c.list = []CommitHeader{}
	}

	om.Log.ToVerbose("commit list", "merge commit list from "+vp.ConfigFileUsed())
	for i, element := range vp.Get("list").([]interface{}) {
		cm := element.(map[interface{}]interface{})

		if _, ok := cm["key"]; !ok {
			om.Log.ToError("Commit list", "List at "+vp.ConfigFileUsed()+" have invalid key format")
			break
		}

		if _, ok := cm["value"]; !ok {
			om.Log.ToError("Commit list", "List at "+vp.ConfigFileUsed()+"value of key="+cm["key"].(string)+" not exist.")
		}

		commitHeader := CommitHeader{
			Key:   cm["key"].(string),
			Value: cm["value"].(string),
		}

		om.Log.ToVerbose("header "+strconv.Itoa(i), commitHeader.String())
		c.list = append(c.list, commitHeader)
	}
	return c
}

// Commit is action for ask the message from user and call CustomCommit.
func (c *Commit) Commit(add, hasMessage bool) {
	// the questions to ask
	var result = c.getQuestion()

	result.Unwrap(func(i interface{}) {
		qs := i.([]*survey.Question)

		if !hasMessage {
			qs = qs[:len(qs)-1]
		}

		om.Log.ToDebug("question list", len(qs))

		answers := CommitMessage{}
		manager.StartResultManager().Save("", survey.Ask(qs, &answers)).IfNoError(func() {
			om.Log.ToDebug("commit key", answers.GetKey())
			om.Log.ToDebug("commit title", answers.Title)
			om.Log.ToDebug("commit message", answers.Message)

			c.CustomCommit(add, answers)
		}).IfError(func(t *manager.Throwable) {
			e.ShowAndExit(e.Update(t, e.UserError))
		})
	}).Catch(func() error {
		return errors.New("Cannot get list of key commit, maybe not exist")
	}, func(t *manager.Throwable) {
		e.ShowAndExit(e.Update(t, e.PreCommitError))
	})
}

// CustomCommit will run git commit -m "<message>" with the default format.
func (c *Commit) CustomCommit(add bool, answers CommitMessage) {
	var commitMessage string
	if answers.Message == "" {
		commitMessage = fmt.Sprintf("[%s] %s", answers.GetKey(), answers.Title)
	} else {
		commitMessage = fmt.Sprintf("[%s] %s\n%s", answers.GetKey(), answers.Title, answers.Message)
	}

	var t *manager.Throwable

	om.Log.ToDebug("commit full", commitMessage)
	if add {
		om.Log.ToVerbose("commit", "with -a flag")
		t = Git().Exec("commit", "-am", commitMessage).Throw()
	} else {
		t = Git().Exec("commit", "-m", commitMessage).Throw()
	}

	e.ShowAndExit(e.Update(t, e.CommitError))
}
