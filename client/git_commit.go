package client

import (
	"errors"
	"fmt"
	"strings"

	"github.com/kamontat/gitgo/models"

	"github.com/manifoldco/promptui"
	"github.com/urfave/cli"
)

func _gitCommit(withAdd bool, key string, emoji bool, title string, msg ...string) error {
	var opt string
	if withAdd {
		opt = "-am"
	} else {
		opt = "-m"
	}
	sep := ""
	if title != "" && len(msg) != 0 {
		sep = ","
	}
	str := fmt.Sprintf("[%s]: %s%s \n%s", key, title, sep, strings.Join(msg, " "))
	if emoji {
		str = fmt.Sprintf("%s: %s%s \n%s", key, title, sep, strings.Join(msg, " "))
	}
	return rawGitCommand("commit", opt, str)
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
		Size:      models.GetUserConfig().Config.Commit.ShowSize,
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
		Active:   "{{ .Key.Emoji.Icon }} {{ .Name | cyan }} [{{ .Title | blue }}]",
		Inactive: "  {{ .Name | cyan }} [{{ .Title | blue }}]",
		Selected: "{{ .Key.Emoji.Icon }} {{ .Name | underline | red | cyan }}",
		Details: `
--------- Commit emoji detail ----------
{{ "Name:" | faint }}	{{ .Name }},
{{ "Key:" | faint }}	{{ .Key.Emoji.Icon }} ({{ .Key.Emoji.Name }})
{{ "Title:" | faint }}	{{ .Title }}`,
	}
	index, _, err := _rawSelectPrompt(db, *templates)
	if models.GetUserConfig().Config.Commit.Emoji == "string" ||
		models.GetUserConfig().Config.Commit.Emoji == "str" ||
		models.GetUserConfig().Config.Commit.Emoji == "text" ||
		models.GetUserConfig().Config.Commit.Emoji == "s" ||
		models.GetUserConfig().Config.Commit.Emoji == "t" {
		key = fmt.Sprintf(":%s:", db[index].Key.Emoji.Name)
	} else {
		key = db[index].Key.Emoji.Icon
	}
	title = db[index].Title
	// fmt.Printf("You choose emoji number %d: %s\n", index+1, db[index])
	return
}

func promptTextKey() (key string, title string, err error) {
	db := models.GetCommitDBConfig().DB
	templates := &promptui.SelectTemplates{
		Label:    "{{ . }}?",
		Active:   "{{ .Key.Emoji.Icon }}  {{ .Name | cyan }} [{{ .Title | blue }}]",
		Inactive: "   {{ .Name | cyan }} [{{ .Title | blue }}]",
		Selected: "{{ .Key.Emoji.Icon }} {{ .Name | underline | red | cyan }}",
		Details: `
--------- Commit text detail ----------
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

// BypassCommit will bypass all check, and commit as initial
func BypassCommit(emoji bool, key string, args ...string) error {
	if key == "init" {
		if emoji {
			return _gitCommit(true, "🎉", emoji, "Initial commit")
		}
		return _gitCommit(true, "init", emoji, "Initial commit")
	} else if key == "release" {
		if emoji {
			return _gitCommit(true, "📌", emoji, "Release new version "+args[0])
		}
		return _gitCommit(true, "release", emoji, "Release new version "+args[0])
	}
	return errors.New("wrong key, this exception shouldn't be throwed")
}

func makeGitCommitWith(emoji bool, withAdd bool, key string, title string, message ...string) (err error) {
	keyExist := _isExist(key)
	titleExist := _isExist(title)
	messageExist := _isNotEmpty(message)

	skipKey := !models.GetUserConfig().Config.Commit.Key.Require
	skipTitle := !models.GetUserConfig().Config.Commit.Title.Require
	skipMessage := !models.GetUserConfig().Config.Commit.Message.Require

	if !models.GetUserConfig().Config.Commit.Key.Require &&
		!models.GetUserConfig().Config.Commit.Title.Require {
		return cli.NewExitError("either 'KEY' or 'TITLE' must be required", 9)
	}

	// KEY
	if keyExist && !skipKey { // exist and required
		var commitDB models.CommitDB
		// convert string key -> commit db
		commitDB, err = models.GetCommitDBConfig().GetCommitDBByName(key)
		if err != nil {
			return
		}
		if emoji {
			key = commitDB.Key.Emoji.Icon
		} else {
			key = commitDB.Key.Text
		}

		// update title, if auto is true, no title exist from option
		if models.GetUserConfig().Config.Commit.Title.Auto && !titleExist {
			if emoji {
				title = commitDB.Title
			} else {
				title, err = models.GetCommitDBConfig().SearchTitleByTextKey(key)
				if err != nil {
					return
				}
			}
		}
	} else if !skipKey {
		var t string
		// prompt key and title
		if emoji {
			key, t, err = promptEmojiKey()
		} else {
			key, t, err = promptTextKey()
		}
		if err != nil {
			return
		}
		// only title auto have been set and no title before
		if models.GetUserConfig().Config.Commit.Title.Auto && !titleExist {
			title = t
		}
	}

	// Update skip only not skipped
	if !skipKey {
		keyExist = _isExist(key)
	}
	if !skipTitle {
		titleExist = _isExist(title)
	}

	// log
	fmt.Printf(
		"Key=\"%s\" (%t), Title=\"%s\" (%t)\n",
		key,
		keyExist,
		title,
		titleExist,
	)

	// TITLE
	if !titleExist && !skipTitle {
		if emoji {
			var commitDB models.CommitDB
			commitDB, err = models.GetCommitDBConfig().GetCommitDBByName(key)
			if err != nil {
				return
			}
			if models.GetUserConfig().Config.Commit.Title.Auto {
				title = commitDB.Title
			} else {
				title, err = promptTitle()
				if err != nil {
					return
				}
			}
		} else {
			if models.GetUserConfig().Config.Commit.Title.Auto {
				title, err = models.GetCommitDBConfig().SearchTitleByTextKey(key)
				if err != nil {
					return
				}
			} else {
				title, err = promptTitle()
				if err != nil {
					return
				}
			}
		}
		titleExist = _isExist(title)
	}

	// MESSAGE
	if !messageExist && !skipMessage {
		var m string
		m, err = promptMessage()
		if err != nil {
			return
		}
		if m != "" {
			message = []string{m}
			messageExist = _isNotEmpty(message)
		}
	}

	_gitCommit(withAdd, key, emoji, title, message...)
	return nil
}

// MakeGitCommitWithText create git commit by text format
func MakeGitCommitWithText(withAdd bool, key string, title string, message ...string) error {
	return makeGitCommitWith(false, withAdd, key, title, message...)
}

// MakeGitCommitWithEmoji create git commit by emoji format
func MakeGitCommitWithEmoji(withAdd bool, key string, title string, message ...string) error {
	return makeGitCommitWith(true, withAdd, key, title, message...)
}
