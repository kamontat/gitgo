package util

import (
	"errors"
	"strconv"

	e "github.com/kamontat/gitgo/exception"
	"github.com/kamontat/gitgo/model"
	"github.com/spf13/viper"
	"gopkg.in/AlecAivazis/survey.v2"
)

// GenerateQuestionViaTypeConfig will generate PromptUI object for prompt input from user
func GenerateQuestionViaTypeConfig(title string, config *model.TypeConfig, listConfig *viper.Viper) (survey.Prompt, []survey.AskOpt) {
	if !config.Enable() {
		return nil, nil
	}

	MaxSize := func(size int) survey.Validator {
		return func(val interface{}) error {
			// since we are validating an Input, the assertion will always succeed
			if str, ok := val.(string); !ok || len(str) > size {
				return errors.New(str + ` is too long (` + strconv.Itoa(len(str)) + `>` + strconv.Itoa(size) + `)`)
			}
			return nil
		}
	}

	validates := []survey.AskOpt{
		survey.WithPageSize(config.Page()),
	}

	if config.Require() {
		validates = append(validates, survey.WithValidator(survey.Required))
	}

	if config.IsList() {
		list := model.List{
			Key: config.Key,
		}

		list.Load(listConfig)

		if !list.IsContain() {
			e.ErrorMessage(e.IsPreCommit, "Cannot prompt commit because list of "+config.Key+" is not exist").ShowMessage().Exit()
		}

		options := list.MakeList()

		if !config.Require() {
			header := model.Header{Type: "", Value: "Empty"}
			options = append(options, header.Format())
		}

		return &survey.Select{
			Message:  title,
			Options:  options,
			PageSize: config.Page(),
		}, validates
	} else if config.IsInput() {
		validates = append(validates, survey.WithValidator(MaxSize(config.Size())))

		return &survey.Input{
			Message: title,
		}, validates
	} else if config.IsMultiline() {
		validates = append(validates, survey.WithValidator(MaxSize(config.Size())))

		return &survey.Multiline{
			Message: title,
		}, validates
	} else if config.IsMix() {
		validates = append(validates, survey.WithValidator(MaxSize(config.Size())))

		// TODO: make mix able to make a list and appendable
		return &survey.Input{
			Message: title,
		}, validates
	}

	// default prompts
	return &survey.Input{
		Message: title,
	}, validates
}
