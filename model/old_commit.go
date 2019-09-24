// Package model provides the model of repository and commit.
// As long as another model that will be use in gitgo command.
package model

// import (
// 	e "github.com/kamontat/gitgo/exception"

// 	manager "github.com/kamontat/go-error-manager"

// 	om "github.com/kamontat/go-log-manager"
// 	survey "gopkg.in/AlecAivazis/survey.v1"
// )

// // CommitOption is a option of commit
// type CommitOption struct {
// 	Add     bool
// 	Empty   bool
// 	Dry     bool
// 	Message bool
// }

// type CommitSettings struct {
// 	ScopeSize   int
// 	MessageSize int
// }

// // Commit is Commit object of deal with commit things.
// type Commit struct {
// 	throwable *manager.Throwable
// 	KeyList   *List
// 	Settings  *CommitSettings
// }

// // SetSettings will update commit settings
// func (c *Commit) SetSettings(scopeSize int, messageSize int) {
// 	if scopeSize == 0 {
// 		scopeSize = 12
// 	}

// 	if messageSize == 0 {
// 		messageSize = 50
// 	}

// 	c.Settings = &CommitSettings{
// 		ScopeSize:   scopeSize,
// 		MessageSize: messageSize,
// 	}
// }

// func (c *Commit) getQuestion() []*survey.Question {
// 	if !c.KeyList.IsContain() {
// 		e.ShowAndExit(e.ErrorMessage(e.IsInitial, "No key list for commit"))
// 	}

// 	if c.Settings == nil {
// 		c.Settings = &CommitSettings{
// 			ScopeSize:   12,
// 			MessageSize: 50,
// 		}
// 	}

// 	return []*survey.Question{
// 		{
// 			Name: "type",
// 			Prompt: &survey.Select{
// 				Message:  "Select commit header",
// 				Options:  c.KeyList.MakeList(),
// 				Help:     "Header will represent 'one word' key of the commit",
// 				PageSize: 5,
// 				VimMode:  true,
// 			},
// 			Validate: survey.Required,
// 		}, {
// 			Name: "scope",
// 			Prompt: &survey.Input{
// 				Message: "Enter commit type scope",
// 				Help:    "Scope should represent the scope of commit type",
// 			},
// 			Validate: survey.MaxLength(c.Settings.ScopeSize),
// 		}, {
// 			Name: "title",
// 			Prompt: &survey.Input{
// 				Message: "Enter commit title",
// 				Help:    "Title will represent one short sentence of the commit",
// 			},
// 			Validate: func(val interface{}) error {
// 				err := survey.Required(val)
// 				if err == nil {
// 					err = survey.MaxLength(c.Settings.MessageSize)(val)
// 				}
// 				return err
// 			},
// 		}, {
// 			Name: "hasMessage",
// 			Prompt: &survey.Confirm{
// 				Message: "Do you have commit message?",
// 			},
// 		},
// 	}
// }

// // Commit is action for ask the message from user and call CustomCommit.
// func (c *Commit) Commit(key string, option CommitOption) {
// 	// the questions to ask
// 	var qs = c.getQuestion()

// 	if key != "" {
// 		qs = qs[1:]
// 	}

// 	if !option.Message {
// 		qs = qs[:len(qs)-1]
// 	}

// 	om.Log.ToDebug("question list", len(qs))

// 	answers := CommitMessage{
// 		Type: key,
// 	}

// 	manager.StartResultManager().Save("", survey.Ask(qs, &answers)).IfNoError(func() {
// 		// if answers.HasMessage {
// 		// 	messageQuestion := &survey.Multiline{
// 		// 		Message: "Enter commit message",
// 		// 		Help:    "Message will represent everything that commit have done",
// 		// 	}
// 		// 	survey.AskOne(messageQuestion, &answers.Message, nil)
// 		// }

// 		om.Log.ToDebug("commit type", answers.GetType())
// 		om.Log.ToDebug("commit scope", answers.Scope)
// 		om.Log.ToDebug("commit title", answers.Title)
// 		om.Log.ToDebug("commit message", answers.Message)

// 		c.CustomCommit(answers, option)
// 	}).IfError(func(t *manager.Throwable) {
// 		e.ShowAndExit(e.Update(t, e.IsUser))
// 	})
// }

// // CustomCommit will run git commit -m "<message>" with the default format.
// func (c *Commit) CustomCommit(answers CommitMessage, option CommitOption) {

// 	var commitMessage = answers.GetMessage()

// 	var t *manager.Throwable

// 	args := []string{"commit"}

// 	if option.Add {
// 		args = append(args, "-a")
// 		om.Log.ToVerbose("commit", "with -a flag")
// 	}

// 	if option.Empty {
// 		args = append(args, "--allow-empty")
// 		om.Log.ToVerbose("commit", "with --allow-empty flag")
// 	}

// 	args = append(args, "-m")
// 	args = append(args, commitMessage)

// 	om.Log.ToInfo("Commit", commitMessage)
// 	if !option.Dry {
// 		t = Git().Exec(args...).Throw()
// 		e.ShowAndExit(e.Update(t, e.IsCommit))
// 	}
// }
