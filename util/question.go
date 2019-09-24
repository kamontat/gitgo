package util

import (
	e "github.com/kamontat/gitgo/exception"
	"github.com/kamontat/gitgo/model"
	"github.com/spf13/viper"
	"gopkg.in/AlecAivazis/survey.v2"
)

// GenerateQuestionViaTypeConfig will generate PromptUI object for prompt input from user
func GenerateQuestionViaTypeConfig(title string, config *model.TypeConfig, listConfig *viper.Viper) survey.Prompt {
	list := model.List{
		Key: config.Key,
	}
	list.Load(listConfig)

	if config.IsList() {
		if !list.IsContain() {
			e.ErrorMessage(e.IsPreCommit, "Cannot prompt commit because list of "+config.Key+" is not exist").ShowMessage().Exit()
		}

		return &survey.Select{
			Message:  title,
			Options:  list.MakeList(),
			PageSize: config.Page(),
		}
	} else if config.IsCustom() {
		return &survey.Input{
			Message: title,
		}
	} else if config.IsMix() {
		// TODO: make mix able to make a list and appendable
		return &survey.Input{
			Message: title,
		}
	}

	// default prompts
	return &survey.Input{
		Message: title,
	}
}
