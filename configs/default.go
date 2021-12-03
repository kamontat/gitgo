package configs

import "github.com/kamontat/gitgo/configs/models"

func Default() *models.Base {
	return &models.Base{
		Version:       5,
		Settings:      models.DefaultSetting(),
		CommitMessage: models.DefaultCommitMessage(),
		Location:      models.DefaultLocation(),
	}
}
