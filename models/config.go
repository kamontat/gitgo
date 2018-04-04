package models

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/urfave/cli"
	yaml "gopkg.in/yaml.v2"
)

func filter(vs []CommitDB, f func(CommitDB) bool) []CommitDB {
	vsf := make([]CommitDB, 0)
	for _, v := range vs {
		if f(v) {
			vsf = append(vsf, v)
		}
	}
	return vsf
}

type Config struct {
	Location     locationConfig
	AppConfig    AppConfig
	UserConfig   UserConfig
	CommitConfig CommitConfig
}

type locationConfig struct {
	AppLocation      string
	UserLocation     string
	CommitDBLocation string
}

type AppConfig struct {
	Name        string
	Description string
	Versions    []VersionConfig
	Since       string
	Authors     []cli.Author
	License     string
}

func (appConfig AppConfig) LatestVersion() VersionConfig {
	return appConfig.Versions[0]
}

type VersionConfig struct {
	Version     string
	Description string
}

func (appConfig AppConfig) PrintEveryVersions() {
	for _, v := range appConfig.Versions {
		v.PrintVersion(appConfig.Name)
	}
}

func (appConfig AppConfig) PrintFullEveryVersions() {
	for _, v := range appConfig.Versions {
		v.PrintFullVersion(appConfig.Name)
	}
}

func (v VersionConfig) PrintVersion(name string) {
	fmt.Printf("%s version %s\n", name, v.Version)
}

func (v VersionConfig) PrintFullVersion(name string) {
	fmt.Printf("%s version %s: %s\n", name, v.Version, v.Description)
}

type CommitConfig struct {
	DB []CommitDB
}

type CommitDB struct {
	Name string
	Key  struct {
		Text  string
		Emoji struct {
			Icon string
			Name string
		}
	}
	Title string
}

func (db CommitConfig) GetCommitDBByEmojiIcon(key string) (result CommitDB, err error) {
	results := filter(db.DB, func(input CommitDB) bool {
		// return strings.Contains(input.Key.Emoji.Icon, key)
		return strings.ToLower(input.Key.Emoji.Icon) == strings.ToLower(key)
	})

	if len(results) == 0 {
		err = errors.New(key + " key not exist!, (get by emoji icon)")
		return
	}
	result = results[0]
	return
}

func (db CommitConfig) GetCommitDBByEmojiName(key string) (result CommitDB, err error) {
	results := filter(db.DB, func(input CommitDB) bool {
		return strings.Contains(strings.ToLower(input.Name), key)
		// return strings.ToLower(input.Name) == strings.ToLower(key)
	})

	if len(results) == 0 {
		err = errors.New(key + " key not exist!, (get by emoji name)")
		return
	}
	result = results[0]
	return
}

func (db CommitConfig) SearchTitleByTextKey(key string) (res string, err error) {
	results := filter(db.DB, func(input CommitDB) bool {
		return strings.Contains(input.Key.Text, key)
		// return input.Key.Text == key
	})

	if len(results) == 0 {
		err = errors.New(key + " key not exist!, (get by text)")
		return
	}
	res = results[0].Title
	return
}

type UserConfig struct {
	Config struct {
		Commit struct {
			Type     string
			Key      InputType
			Title    InputType
			Message  InputType
			ShowSize int
		}
		Editor string
	}
}

func (user UserConfig) IsEmojiType() bool {
	return user.Config.Commit.Type == "emoji" ||
		user.Config.Commit.Type == "moji" ||
		user.Config.Commit.Type == "e"
}

func (user UserConfig) IsTextType() bool {
	return user.Config.Commit.Type == "text" ||
		user.Config.Commit.Type == "t"
}

type InputType struct {
	Require bool
	Auto    bool
	Size    int
}

func _setupPath(location string, filename string) string {
	if location == "" && os.Getenv("GOPATH") == "" {
		cli.HandleExitCoder(cli.NewExitError("user location or $GOPATH must be set", 2))
	}
	if location == "" {
		return os.Getenv("GOPATH") + "/src/gitgo/config/" + filename
	}
	return location + "/" + filename
}

func setupAppConfig(location string) (appConfig AppConfig) {
	file, e := ioutil.ReadFile(location)
	if e != nil {
		fmt.Printf("File error: %v\n", e)
		os.Exit(1)
	}

	yaml.Unmarshal(file, &appConfig)
	return
}

func setupUserConfig(location string) (userConfig UserConfig) {
	file, e := ioutil.ReadFile(location)
	if e != nil {
		fmt.Printf("File error: %v\n", e)
		os.Exit(1)
	}

	yaml.Unmarshal(file, &userConfig)
	return
}

func setupCommitDBConfig(location string) (commitConfig CommitConfig) {
	var db []CommitDB
	file, e := ioutil.ReadFile(location)
	if e != nil {
		fmt.Printf("File error: %v\n", e)
		os.Exit(1)
	}

	yaml.Unmarshal(file, &db)
	return CommitConfig{
		DB: db,
	}
}

func setupLocationConfig(dev bool) (location locationConfig, err error) {
	home := os.Getenv("HOME")
	if home == "" {
		err = cli.NewExitError("$HOME must be set", 2)
	}
	var defaultLocation string
	appfile := "app.yaml"
	userfile := "user.yaml"
	commitdbfile := "commit_list.yaml"

	if !dev {
		defaultLocation = home + "/.config/gitgo/config"
	}

	location = locationConfig{
		AppLocation:      _setupPath(defaultLocation, appfile),
		UserLocation:     _setupPath(defaultLocation, userfile),
		CommitDBLocation: _setupPath(defaultLocation, commitdbfile),
	}

	if _, err = os.Stat(location.AppLocation); os.IsNotExist(err) {
		return
		// cli.HandleExitCoder(cli.NewExitError("App config not exist, Add to "+location.AppLocation, 2))
	}
	if _, err = os.Stat(location.CommitDBLocation); os.IsNotExist(err) {
		return
		// cli.HandleExitCoder(cli.NewExitError("Commit database not exist, Add to "+location.CommitDBLocation, 2))
	}
	if _, err = os.Stat(location.UserLocation); os.IsNotExist(err) {
		return
		// cli.HandleExitCoder(cli.NewExitError("User config not exist, Add to "+location.UserLocation, 2))
	}

	return
}

func setupConfig(dev bool) Config {
	location, err := setupLocationConfig(dev)
	if err != nil {
		log.Fatalln(err)
	}
	return Config{
		Location:     location,
		AppConfig:    setupAppConfig(location.AppLocation),
		UserConfig:   setupUserConfig(location.UserLocation),
		CommitConfig: setupCommitDBConfig(location.CommitDBLocation),
	}
}

var configFile Config
var setupError error

func Setup(dev bool) {
	configFile = setupConfig(dev)
}

func GetAppConfig() AppConfig {
	return configFile.AppConfig
}

func GetUserConfig() UserConfig {
	return configFile.UserConfig
}

func GetCommitDBConfig() CommitConfig {
	return configFile.CommitConfig
}

func GetAppLocation() locationConfig {
	return configFile.Location
}
