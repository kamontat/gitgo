package client

import (
	"github.com/kamontat/gitgo/models"
	survey "gopkg.in/AlecAivazis/survey.v1"
)

// GenerateKeyPrompt prompt user to input key of the commit
// Use survey Select for prompting
func GenerateKeyPrompt() survey.Prompt {
	return &survey.Select{
		Message:  "Choose commit header:",
		Help:     "Header will represent 'one word' key of the commit",
		Options:  models.GetConfigHelper().CommitAsStringArray(),
		PageSize: models.GetUserConfig().Config.Commit.ShowSize,
	}
}

// GenerateTitleValidator validator of user title commit
// Use survey Validator for validate
func GenerateTitleValidator() survey.Validator {
	return survey.MaxLength(models.GetUserConfig().Config.Commit.Title.Size)
}

// GenerateTitlePrompt prompt user to input title of the commit
// Use survey Input for prompting
func GenerateTitlePrompt() survey.Prompt {
	return &survey.Input{
		Message: "Choose commit title:",
		Help:    "Title will represent one short sentence of the commit",
	}
}

// GenerateMessagePrompt prompt user to input message of the commit
// Use survey Editor for prompting
func GenerateMessagePrompt() survey.Prompt {
	return &survey.Editor{
		Message: "Choose commit message:",
		Help:    "Message will represent everything that commit have done",
	}
}
