package utils

import (
	"regexp"

	"github.com/AlecAivazis/survey/v2"
	"github.com/kamontat/gitgo/config/models"
)

const defaultSelectPage = 5

func askSelect(data *string, message string, config *models.CommitTypeSetting, validates []survey.AskOpt) error {
	var prompt survey.Prompt

	var pageSize = int(config.Prompt.Select.Page)
	var suggestion = config.Prompt.Select.Suggestion
	var values = config.Prompt.Select.List()

	if pageSize > 0 {
		validates = append(validates, survey.WithPageSize(pageSize))
	} else {
		validates = append(validates, survey.WithPageSize(defaultSelectPage))
	}

	if suggestion {
		prompt = &survey.Input{
			Message: message,
			Suggest: func(str string) (result []string) {
				for _, value := range values {
					if match, err := regexp.MatchString(str, value); match && err == nil {
						result = append(result, value)
					}
				}

				return
			},
		}
	} else {
		if !config.Required {
			optional := &models.CommitSelectValue{
				Key:  "",
				Text: "go to next question",
			}

			values = append(values, optional.Format())
		}

		prompt = &survey.Select{
			Message:  message,
			Options:  values,
			PageSize: pageSize,
		}
	}

	err := survey.AskOne(prompt, data, validates...)
	if err != nil {
		return err
	}

	// update result and remove formatted string on select type
	var result = updateSelectResult(*data)
	*data = result

	return nil
}
