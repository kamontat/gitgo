package models

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/urfave/cli"
	yaml "gopkg.in/yaml.v2"
)

type Config struct {
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
	db []commitDB
}

type commitDB struct {
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

type UserConfig struct {
	Config struct {
		Commit struct {
			Type string
		}
	}
}

func _setupPath(location string, filename string) string {
	if location == "" && os.Getenv("GOPATH") == "" {
		cli.HandleExitCoder(cli.NewExitError("user location or $GOPATH must be set", 2))
	}
	if location == "" {
		return os.Getenv("GOPATH") + "/src/gitgo/config/" + filename
	}
	return location
}

func setupAppConfig(appLocation string) (appConfig AppConfig) {
	file, e := ioutil.ReadFile(_setupPath(appLocation, "app.yaml"))
	if e != nil {
		fmt.Printf("File error: %v\n", e)
		os.Exit(1)
	}

	yaml.Unmarshal(file, &appConfig)
	return
}

func setupUserConfig(appLocation string) (userConfig UserConfig) {
	file, e := ioutil.ReadFile(_setupPath(appLocation, "user.yaml"))
	if e != nil {
		fmt.Printf("File error: %v\n", e)
		os.Exit(1)
	}

	yaml.Unmarshal(file, &userConfig)
	return
}

func setupCommitDBConfig(location string) (commitConfig CommitConfig) {
	var db []commitDB
	file, e := ioutil.ReadFile(_setupPath(location, "commit_list.yaml"))
	if e != nil {
		fmt.Printf("File error: %v\n", e)
		os.Exit(1)
	}

	yaml.Unmarshal(file, &db)
	return CommitConfig{
		db: db,
	}
}

func setupLocationConfig(dev bool) (locationConfig locationConfig) {
	if dev {
		return locationConfig
	}
	home := os.Getenv("HOME")
	userfile := "/user.yml"
	appfile := "/app.yml"
	commitdbfile := "/commit_list.yml"
	location := ""
	if home == "" {
		cli.HandleExitCoder(cli.NewExitError("$HOME must be set", 2))
		return
	}
	location = fmt.Sprintf("%s/.config/gitgo/config", home)
	if _, err := os.Stat(location + userfile); os.IsNotExist(err) {
		cli.HandleExitCoder(cli.NewExitError("User config not exist, Add to "+location, 2))
	} else {
		locationConfig.UserLocation = location + userfile
	}
	if _, err := os.Stat(location + appfile); os.IsNotExist(err) {
		cli.HandleExitCoder(cli.NewExitError("App config not exist, Add to "+location, 2))
	} else {
		locationConfig.AppLocation = location + appfile
	}
	if _, err := os.Stat(location + commitdbfile); os.IsNotExist(err) {
		cli.HandleExitCoder(cli.NewExitError("Commit config not exist, Add to "+location, 2))
	} else {
		locationConfig.CommitDBLocation = location + commitdbfile
	}
	return
}

func setupConfig(dev bool) Config {
	location := setupLocationConfig(dev)

	return Config{
		AppConfig:    setupAppConfig(location.AppLocation),
		UserConfig:   setupUserConfig(location.UserLocation),
		CommitConfig: setupCommitDBConfig(location.CommitDBLocation),
	}
}

var configFile Config

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
