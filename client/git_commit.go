package client

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/kamontat/gitgo/models"
	"github.com/urfave/cli"
	"gopkg.in/AlecAivazis/survey.v1"
)

// ##########################################
// ## Raw method                           ##
// ##########################################

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

// ##########################################
// ## Private method                       ##
// ##########################################

func setKeyType(withEmoji bool, commit models.Commit) string {
	if withEmoji {
		if models.GetUserConfig().Config.Commit.Emoji == "string" ||
			models.GetUserConfig().Config.Commit.Emoji == "str" ||
			models.GetUserConfig().Config.Commit.Emoji == "text" ||
			models.GetUserConfig().Config.Commit.Emoji == "s" ||
			models.GetUserConfig().Config.Commit.Emoji == "t" {
			return fmt.Sprintf(":%s:", commit.Key.Emoji.Name)
		}
		return commit.Key.Emoji.Icon
	}
	return commit.Key.Text
}

func updateKey(withEmoji bool, k string) (key string, title string, err error) {
	if !_isExist(k) {
		err = survey.AskOne(GenerateKeyPrompt(), &k, nil)
	}

	var commit models.Commit
	commit, err = models.GetConfigHelper().GetCommitByName(k)
	if err != nil {
		return
	}

	key = setKeyType(withEmoji, commit)
	if models.GetUserConfig().Config.Commit.Title.Auto {
		title = commit.Title
	}
	return
}

func updateTitle(t string) (title string, err error) {
	titleExist := _isExist(t)
	if titleExist {
		title = t
		return
	}

	// err = survey.Ask(GeneratetitlePrompt(), &title)
	err = survey.AskOne(GenerateTitlePrompt(), &title, GenerateTitleValidator())
	return
}

func updateMessage(msg ...string) (message []string, err error) {
	var rawMessage string
	messageExist := _isNotEmpty(msg)
	if messageExist {
		message = msg
		return
	}

	// set custom editor to run
	// FIXME: not work!
	os.Setenv("VISUAL", models.GetUserConfig().Config.Editor)
	err = survey.AskOne(GenerateMessagePrompt(), &rawMessage, nil)
	message = []string{rawMessage}
	return
}

func validateConfiguration() error {
	// key and title cannot set not required together
	if !models.GetUserConfig().Config.Commit.Key.Require &&
		!models.GetUserConfig().Config.Commit.Title.Require {
		return cli.NewExitError("either 'KEY' or 'TITLE' must be required", 9)
	}
	return nil
}

func validateCommitMessage(key string, skipK bool, title string, skipT bool, msg []string, skipM bool) error {
	str := ""
	if !_isExist(key) && !skipK {
		str += "Key "
	}
	if !_isExist(title) && !skipT {
		str += "Title "
	}
	if !_isNotEmpty(msg) && !skipM {
		str += "Message "
	}

	if _isExist(str) {
		return cli.NewExitError(str+" is/are not exist!", 1)
	}
	return nil
}

func makeGitCommitWith(emoji bool, withAdd bool, key string, title string, message ...string) (err error) {
	skipKey := !models.GetUserConfig().Config.Commit.Key.Require
	skipTitle := !models.GetUserConfig().Config.Commit.Title.Require
	skipMessage := !models.GetUserConfig().Config.Commit.Message.Require

	var newKey, newTitle string
	var newMessage []string

	err = validateConfiguration()
	if err != nil {
		return err
	}

	// key management
	if !skipKey {
		newKey, newTitle, err = updateKey(emoji, key)
		if err != nil {
			return err
		}
	}

	// title management
	if !skipTitle && newTitle == "" {
		newTitle, err = updateTitle(title)
		if err != nil {
			return err
		}
	}

	// message management
	if !skipMessage {
		newMessage, err = updateMessage(message...)
		if err != nil {
			return err
		}
	}

	err = validateCommitMessage(newKey, skipKey, newTitle, skipTitle, newMessage, skipMessage)
	if err != nil {
		return err
	}

	return _gitCommit(withAdd, newKey, emoji, newTitle, newMessage...)
}

// ##########################################
// ## Public method                        ##
// ##########################################

// MakeGitCommitWithText create git commit by text format
func MakeGitCommitWithText(withAdd bool, key string, title string, message ...string) error {
	return makeGitCommitWith(false, withAdd, key, title, message...)
}

// MakeGitCommitWithEmoji create git commit by emoji format
func MakeGitCommitWithEmoji(withAdd bool, key string, title string, message ...string) error {
	return makeGitCommitWith(true, withAdd, key, title, message...)
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
