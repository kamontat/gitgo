package model

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/kamontat/go-error-manager"

	"github.com/spf13/viper"

	"github.com/kamontat/go-log-manager"
	survey "gopkg.in/AlecAivazis/survey.v1"
)

type CommitHeader struct {
	Key   string
	Value string
}

func (c *CommitHeader) Format() string {
	return fmt.Sprintf("%-10s: %s", c.Key, c.Value)
}

func (c *CommitHeader) String() string {
	return fmt.Sprintf("commit key=%s, value=%s", c.Key, c.Value)
}

type Commit struct {
	repo *Repo
	list []CommitHeader
}

func (c *Commit) ListHeaderOptions() (list []string) {
	for _, commits := range c.list {
		list = append(list, commits.Format())
	}
	return
}

type CommitMessage struct {
	Key     string
	Title   string
	Message string
}

func (c *CommitMessage) GetKey() string {
	s := strings.Split(c.Key, ":")[0]
	return strings.TrimSpace(s)
}

func (c *Commit) getQuestion() []*survey.Question {
	return []*survey.Question{
		{
			Name: "key",
			Prompt: &survey.Select{
				Message:  "Select commit header",
				Options:  c.ListHeaderOptions(),
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

func (c *Commit) LoadList(vp *viper.Viper) *Commit {
	if vp == nil {
		om.Log().ToDebug("commit list", "viper is nil, cannot load list")
		return c
	}

	// reset list
	c.list = []CommitHeader{}

	om.Log().ToDebug("commit list", "load commit list from "+vp.ConfigFileUsed())
	return c.MergeList(vp)
}

func (c *Commit) MergeList(vp *viper.Viper) *Commit {
	if vp == nil {
		om.Log().ToVerbose("commit list", "viper is nil, cannot merge list")
		return c
	}
	if c.list == nil {
		c.list = []CommitHeader{}
	}

	if vp.Get("list") == nil {
		return c
	}

	om.Log().ToVerbose("commit list", "merge commit list from "+vp.ConfigFileUsed())
	for i, element := range vp.Get("list").([]interface{}) {
		cm := element.(map[interface{}]interface{})
		commitHeader := CommitHeader{
			Key:   cm["key"].(string),
			Value: cm["value"].(string),
		}

		om.Log().ToVerbose("header "+strconv.Itoa(i), commitHeader.String())
		c.list = append(c.list, commitHeader)
	}

	return c
}

func (c *Commit) Commit(hasMessage bool) {
	// the questions to ask
	var qs = c.getQuestion()
	if !hasMessage {
		qs = qs[:len(qs)-1]
	}

	om.Log().ToDebug("question list", strconv.Itoa(len(qs)))

	answers := CommitMessage{}

	// perform the questions
	manager.StartNewManageError().E1P(survey.Ask(qs, &answers)).Throw().ShowMessage(nil).Exit()

	om.Log().ToDebug("commit key", answers.GetKey())
	om.Log().ToDebug("commit title", answers.Title)
	om.Log().ToDebug("commit message", answers.Message)

	c.CustomCommit(answers)
}

func (c *Commit) CustomCommit(answers CommitMessage) {
	var commitMessage string
	if answers.Message == "" {
		commitMessage = fmt.Sprintf("[%s] %s", answers.GetKey(), answers.Title)
	} else {
		commitMessage = fmt.Sprintf("[%s] %s\n%s", answers.GetKey(), answers.Title, answers.Message)
	}

	om.Log().ToVerbose("commit full", commitMessage)
	Git().Exec("commit", "-m", commitMessage)
}
