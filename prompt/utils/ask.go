package utils

import (
	"errors"

	"github.com/AlecAivazis/survey/v2"
	"github.com/kamontat/gitgo/config/models"
)

func Ask(data *string, message string, config *models.CommitTypeSetting) error {
	if !config.Enabled {
		return nil
	}

	validates := make([]survey.AskOpt, 0)

	// Default
	if config.Required {
		validates = append(validates, survey.WithValidator(survey.Required))
	}

	if config.Prompt.Input != nil {
		return askInput(data, message, config, validates)
	} else if config.Prompt.Select != nil {
		return askSelect(data, message, config, validates)
	} else {
		return errors.New("unsupported prompt type")
	}
}
