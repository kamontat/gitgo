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
	rawGitCommand("commit", opt, fmt.Sprintf("[%s]: %s\n%s", key, title, strings.Join(msg, " ")))
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
	key = db[index].Key.Emoji.Icon
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

func makeGitCommitWith(emoji bool, withAdd bool, key string, title string, message ...string) (err error) {
	keyExist := _isExist(key)
	titleExist := _isExist(title)
	messageExist := _isNotEmpty(message)

	skipKey := !models.GetUserConfig().Config.Commit.Key.Require
	skipTitle := !models.GetUserConfig().Config.Commit.Title.Require
	skipMessage := !models.GetUserConfig().Config.Commit.Message.Require

	// fmt.Printf(
	// 	"Key=\"%s\" (%t),\n",
	// 	key,
	// 	keyExist,
	// )
	// fmt.Printf(
	// 	"Title=\"%s\" (%t),\n",
	// 	title,
	// 	titleExist,
	// )
	// fmt.Printf(
	// 	"Message=\"%s\" (%t)\n",
	// 	strings.Join(message, ", "),
	// 	messageExist,
	// )

	if !models.GetUserConfig().Config.Commit.Key.Require &&
		!models.GetUserConfig().Config.Commit.Title.Require {
		return cli.NewExitError("either 'KEY' or 'TITLE' must be required", 9)
	}

	// KEY
	if keyExist && !skipKey { // exist and required
		var commitDB models.CommitDB
		if emoji {
			// convert string key -> emoji icon
			commitDB, err = models.GetCommitDBConfig().GetCommitDBByEmojiName(key)
			if err != nil {
				return
			}
			key = commitDB.Key.Emoji.Icon
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

	fmt.Printf(
		"Key=\"%s\" (%t),\n",
		key,
		keyExist,
	)
	fmt.Printf(
		"Title=\"%s\" (%t),\n",
		title,
		titleExist,
	)

	// TITLE
	if !titleExist && !skipTitle {
		if emoji {
			var commitDB models.CommitDB
			commitDB, err = models.GetCommitDBConfig().GetCommitDBByEmojiName(key)
			if err != nil {
				return
			}
			title = commitDB.Title
		} else {
			title, err = models.GetCommitDBConfig().SearchTitleByTextKey(key)
			if err != nil {
				return
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
		message = []string{m}
		messageExist = _isNotEmpty(message)
	}

	_gitCommit(withAdd, key, title, message...)
	return nil
	// var keystr, tilstr, msgstr string = "not-required", "not-required", "not-required"
	// if skipKey {
	// 	keystr = "required"
	// }
	// if skipTitle {
	// 	tilstr = "required"
	// }
	// if skipMessage {
	// 	msgstr = "required"
	// }
	// return cli.NewExitError(str, 5)
}

func MakeGitCommitWithText(withAdd bool, key string, title string, message ...string) error {
	return makeGitCommitWith(false, withAdd, key, title, message...)
}

func MakeGitCommitWithEmoji(withAdd bool, key string, title string, message ...string) error {
	return makeGitCommitWith(true, withAdd, key, title, message...)
}
