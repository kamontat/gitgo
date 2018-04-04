package client

import (
	"errors"
	"fmt"
	"gitgo/models"
	"strings"

	"github.com/manifoldco/promptui"
	"github.com/urfave/cli"
)

func _gitCommit(withAdd bool, key string, title string, msg ...string) {
	var opt string
	if withAdd {
		opt = "-am"
	} else {
		opt = "-m"
	}
	rawGitCommand("commit", opt, fmt.Sprintf("\"[%s]: %s\n%s\"", key, title, strings.Join(msg, " ")))
}

func _isExist(a string) bool {
	return a != ""
}

func _isNotEmpty(a []string) bool {
	return len(a) != 0
}

func _rawSelectPrompt(db []models.CommitDB, templates promptui.SelectTemplates) (index int, value string, err error) {
	prompt := promptui.Select{
		Label:     "Commit header",
		Items:     db,
		Templates: &templates,
		Size:      models.GetUserConfig().Config.Commit.ListSize,
		Searcher: func(input string, index int) bool {
			commit := db[index]
			name := strings.Replace(strings.ToLower(commit.Name), " ", "", -1)
			input = strings.Replace(strings.ToLower(input), " ", "", -1)
			return strings.Contains(name, input)
		},
	}
	index, value, err = prompt.Run()
	return
}

func _rawPrompt(label string, defaultValue string, validator promptui.ValidateFunc) (result string, err error) {
	prompt := promptui.Prompt{
		Label:    label,
		Validate: validator,
		Default:  defaultValue,
	}

	result, err = prompt.Run()
	return
}

func promptEmojiKey() (key string, title string, err error) {
	db := models.GetCommitDBConfig().DB
	templates := &promptui.SelectTemplates{
		Label:    "{{ . }}?",
		Active:   "{{ .Key.Emoji.Icon }} {{ .Name | cyan }}",
		Inactive: "  {{ .Name | cyan }}",
		Selected: "{{ .Key.Emoji.Icon }} {{ .Name | red | cyan }}",
		Details: `
--------- Commit detail ----------
{{ "Name:" | faint }}	{{ .Name }},
{{ "Key:" | faint }}	{{ .Key.Emoji.Icon }} ({{ .Key.Emoji.Name }})
{{ "Title:" | faint }}	{{ .Title }}`,
	}
	index, _, err := _rawSelectPrompt(db, *templates)
	key = db[index].Key.Emoji.Icon
	title = db[index].Title
	// fmt.Printf("You choose emoji number %d: %s\n", index+1, db[index])
	return
}

func promptTextKey() (key string, title string, err error) {
	db := models.GetCommitDBConfig().DB
	templates := &promptui.SelectTemplates{
		Label:    "{{ . }}?",
		Active:   "{{ .Key.Emoji.Icon }}  {{ .Name | cyan }}",
		Inactive: "   {{ .Name | cyan }}",
		Selected: "[{{ .Key.Emoji.Icon }}] {{ .Name | underline | red | cyan }}",
		Details: `
--------- Commit detail ----------
{{ "Name:" | faint }}	{{ .Name }},
{{ "Key:" | faint }}	{{ .Key.Text }}
{{ "Title:" | faint }}	{{ .Title }}`,
	}
	index, _, err := _rawSelectPrompt(db, *templates)
	key = db[index].Key.Text
	title = db[index].Title
	// fmt.Printf("You choose text number %d: %s\n", index+1, db[index])
	return
}

func promptTitle( /* index int */ ) (t string, err error) {
	t, err = _rawPrompt("Commit title", "", func(input string) error {
		if len(input) > models.GetUserConfig().Config.Commit.Title.Size {
			return errors.New("Commit title shouldn't more than 50 characters")
		}
		return nil
	})
	return
}

func promptMessage() (m string, err error) {
	m, err = _rawPrompt("Commit message", "", func(input string) error {
		return nil
	})
	return
}

func makeGitCommitWith(emoji bool, withAdd bool, key string, title string, message ...string) (err error) {
	skipKey := _isExist(key) || !models.GetUserConfig().Config.Commit.Key.Require
	skipTitle := _isExist(title) || !models.GetUserConfig().Config.Commit.Title.Require
	skipMessage := _isNotEmpty(message) || !models.GetUserConfig().Config.Commit.Message.Require

	if !models.GetUserConfig().Config.Commit.Key.Require &&
		!models.GetUserConfig().Config.Commit.Title.Require {
		return cli.NewExitError("either 'KEY' or 'TITLE' must be required", 9)
	}

	if !skipKey {
		var t string
		if emoji {
			key, t, err = promptEmojiKey()
		} else {
			key, t, err = promptTextKey()
		}
		if err != nil {
			return
		}
		if models.GetUserConfig().Config.Commit.Title.Auto && !skipTitle {
			title = t
		}
		skipKey = _isExist(key)
		skipTitle = _isExist(title)
	}

	if skipKey && emoji {
		key = models.GetCommitDBConfig().GetEmojiByKey(key)
	} else if skipKey && !emoji {
		if models.GetUserConfig().Config.Commit.Title.Auto {
			title = models.GetCommitDBConfig().SearchTitleByTextKey(key)
			skipTitle = _isExist(title)
		}
	}

	if !models.GetUserConfig().Config.Commit.Title.Auto && !skipTitle {
		title, err = promptTitle()
		if err != nil {
			return
		}
		skipTitle = _isExist(title)
	}

	if !skipMessage {
		var m string
		m, err = promptMessage()
		if err != nil {
			return
		}
		message = []string{m}
	}

	if skipKey && skipTitle /* && skipMessage */ {
		_gitCommit(withAdd, key, title, message...)
		return nil
	}
	return cli.NewExitError(fmt.Sprintf("required string no exist, key=%t, title=%t, message=%t", skipKey, skipTitle, skipMessage), 5)
}

func MakeGitCommitWithText(withAdd bool, key string, title string, message ...string) error {
	return makeGitCommitWith(false, withAdd, key, title, message...)
}

func MakeGitCommitWithEmoji(withAdd bool, key string, title string, message ...string) error {
	return makeGitCommitWith(true, withAdd, key, title, message...)
}
