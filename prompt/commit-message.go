package prompt

import (
	"github.com/kamontat/gitgo/config/models"
	git "github.com/kamontat/gitgo/git/models"
	"github.com/kamontat/gitgo/prompt/utils"
)

func CommitMessage(config *models.CommitSetting) (answers git.CommitMessage, err error) {
	err = utils.Ask(&answers.Key, "Please enter commit key", config.Key)
	if err != nil {
		return answers, err
	}

	err = utils.Ask(&answers.Scope, "Please enter commit scope", config.Scope)
	if err != nil {
		return answers, err
	}

	err = utils.Ask(&answers.Title, "Please enter commit title", config.Title)
	if err != nil {
		return answers, err
	}

	err = utils.Ask(&answers.Message, "Please enter commit message", config.Message)
	if err != nil {
		return answers, err
	}

	return answers, nil
}
