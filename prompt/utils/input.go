package utils

import (
	"errors"
	"strconv"

	"github.com/AlecAivazis/survey/v2"
	"github.com/kamontat/gitgo/config/models"
)

const defaultMaximumLength = 80

func maxLength(size int) survey.Validator {
	return func(val interface{}) error {
		// since we are validating an Input, the assertion will always succeed
		if str, ok := val.(string); !ok || len(str) > size {
			return errors.New(str + ` is too long (` + strconv.Itoa(len(str)) + `>` + strconv.Itoa(size) + `)`)
		}
		return nil
	}
}

func askInput(data *string, message string, config *models.CommitTypeSetting, validates []survey.AskOpt) error {
	var prompt survey.Prompt
	var inputConfig = config.Prompt.Input

	if inputConfig.Max > 0 {
		validates = append(validates, survey.WithValidator(maxLength(int(inputConfig.Max))))
	} else {
		validates = append(validates, survey.WithValidator(maxLength(defaultMaximumLength)))
	}

	if inputConfig.Multiline {
		prompt = &survey.Multiline{
			Message: message,
		}
	} else {
		prompt = &survey.Input{
			Message: message,
		}
	}

	return survey.AskOne(prompt, data, validates...)
}
