package model

import (
	"github.com/spf13/viper"
)

// ConfigableConfig is a type of configuration inside commit and branch config
type ConfigableConfig struct {
	Enable  bool
	Require bool
}

// TypeConfig is a type of configuration inside commit and branch config
type TypeConfig struct {
	Key  string
	Type string
	size int
	page int
}

// IsList for check type is list type ?
func (t *TypeConfig) IsList() bool {
	return t.Type == "list"
}

// IsCustom for check type is custom type ?
func (t *TypeConfig) IsCustom() bool {
	return t.Type == "custom"
}

// IsMix for check type is mix type ?
func (t *TypeConfig) IsMix() bool {
	return t.Type == "mix"
}

// Size will return size or default size if not exist
func (t *TypeConfig) Size() int {
	if t.size == 0 {
		return 20
	}

	return t.size
}

// Page will return page size or default page size if not exist
func (t *TypeConfig) Page() int {
	if t.page == 0 {
		return 5
	}

	return t.page
}

// CommitConfig is a structure of commit configuration
type CommitConfig struct {
	Version int
	Key     *TypeConfig
	Scope   *TypeConfig
	Title   *TypeConfig
	Message *TypeConfig
}

// BranchConfig is a structure of branch configuration
type BranchConfig struct {
	Version     int
	Iteration   ConfigableConfig
	Key         *TypeConfig
	Title       *TypeConfig
	Description ConfigableConfig
}

// LoadCommitConfiguration will return Commit config object from yaml config
func LoadCommitConfiguration(vp *viper.Viper) *CommitConfig {
	viper.SetDefault("commit.key.size", 15)
	viper.SetDefault("commit.key.type", "list")
	viper.SetDefault("commit.key.page", 5)

	return &CommitConfig{
		Version: vp.GetInt("version"),
		Key: &TypeConfig{
			Key:  "commit.keys",
			size: vp.GetInt("commit.key.size"),
			Type: vp.GetString("commit.key.type"),
			page: vp.GetInt("commit.key.page"),
		},
		Title: &TypeConfig{
			Key:  "commit.titles",
			size: vp.GetInt("commit.title.size"),
			Type: vp.GetString("commit.title.type"),
			page: vp.GetInt("commit.title.page"),
		},
		Scope: &TypeConfig{
			Key:  "commit.scopes",
			size: vp.GetInt("commit.scope.size"),
			Type: vp.GetString("commit.scope.type"),
			page: vp.GetInt("commit.scope.page"),
		},
		Message: &TypeConfig{
			Key:  "commit.messages",
			size: vp.GetInt("commit.message.size"),
			Type: vp.GetString("commit.message.type"),
			page: vp.GetInt("commit.message.page"),
		},
	}
}

// LoadBranchConfiguration will return Branch config object from yaml config
func LoadBranchConfiguration(vp *viper.Viper) *BranchConfig {
	return &BranchConfig{
		Version: vp.GetInt("version"),
		Iteration: ConfigableConfig{
			Enable:  vp.GetBool("branch.iteration.enable"),
			Require: vp.GetBool("branch.iteration.require"),
		},
		Key: &TypeConfig{
			Key:  "branch.keys",
			size: vp.GetInt("branch.key.size"),
			Type: vp.GetString("branch.key.type"),
			page: vp.GetInt("branch.key.page"),
		},
		Title: &TypeConfig{
			Key:  "branch.titles",
			size: vp.GetInt("branch.title.size"),
			Type: vp.GetString("branch.title.type"),
			page: vp.GetInt("branch.title.page"),
		},
		Description: ConfigableConfig{
			Enable:  vp.GetBool("branch.description.enable"),
			Require: vp.GetBool("branch.description.require"),
		},
	}
}
