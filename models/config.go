package models

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"reflect"
	"strconv"
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

// Config every program config value
type Config struct {
	Location     LocationConfig
	AppConfig    AppConfig
	UserConfig   UserConfig
	CommitConfig CommitConfig
}

// LocationConfig location of every settings of the application
type LocationConfig struct {
	DevConfigLocation  string
	ProdConfigLocation string
	AppLocation        string
	UserLocation       string
	CommitDBLocation   string
}

// AppConfig application config, from 'app.yaml'
type AppConfig struct {
	Name        string
	Description string
	Versions    []VersionConfig
	Since       string
	Authors     []cli.Author
	License     string
}

// LatestVersion get latest version of application config
func (appConfig AppConfig) LatestVersion() VersionConfig {
	return appConfig.Versions[0]
}

// VersionConfig version config, contains version number, and description
type VersionConfig struct {
	Version     string
	Description string
}

// ChooseToPrintEveryVersions choose format to print every version that exist in app config
func (appConfig AppConfig) ChooseToPrintEveryVersions(full bool) {
	for _, v := range appConfig.Versions {
		v.ChooseToPrintVersion(full)
	}
}

// PrintEveryVersions print every version that exist in app config
func (appConfig AppConfig) PrintEveryVersions() {
	for _, v := range appConfig.Versions {
		v.PrintVersion()
	}
}

// PrintFullEveryVersions print every version (in full format) that exist in app config
func (appConfig AppConfig) PrintFullEveryVersions() {
	for _, v := range appConfig.Versions {
		v.PrintFullVersion()
	}
}

// PrintVersion print version in application format
func (v VersionConfig) PrintVersion() {
	name := GetAppConfig().Name
	fmt.Printf("%s version %s\n", name, v.Version)
}

// ChooseToPrintVersion choose to print version in application format
func (v VersionConfig) ChooseToPrintVersion(full bool) {
	if full {
		v.PrintFullVersion()
	} else {
		v.PrintVersion()
	}
}

// PrintFullVersion print version in full application format
func (v VersionConfig) PrintFullVersion() {
	name := GetAppConfig().Name
	fmt.Printf("%s version %s: %s\n", name, v.Version, v.Description)
}

// CommitConfig usually this struct will be the commit list database
type CommitConfig struct {
	DB []CommitDB
}

// CommitDB commit struct in yaml files
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

// GetCommitDBByEmojiIcon get commit db struct by icon
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

// GetCommitDBByEmojiName get commit db struct by commit name
func (db CommitConfig) GetCommitDBByEmojiName(key string) (result CommitDB, err error) {
	results := filter(db.DB, func(input CommitDB) bool {
		return strings.Contains(strings.ToLower(input.Name), key)
		// return strings.ToLower(input.Name) == strings.ToLower(key)
	})

	if len(results) == 0 {
		err = errors.New(key + " key not exist!, (get by commit name)")
		return
	}
	result = results[0]
	return
}

// SearchTitleByTextKey get commit title by key text
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

// UserConfig user struct from 'user.yaml'
type UserConfig struct {
	Config struct {
		Commit struct {
			Type     string
			Emoji    string
			Key      InputType
			Title    InputType
			Message  InputType
			ShowSize int
		}
		Editor string
	}
}

// IsEmojiType check is user choose emoji commit
func (user UserConfig) IsEmojiType() bool {
	return user.Config.Commit.Type == "emoji" ||
		user.Config.Commit.Type == "moji" ||
		user.Config.Commit.Type == "e"
}

// IsTextType check is user choose text commit
func (user UserConfig) IsTextType() bool {
	return user.Config.Commit.Type == "text" ||
		user.Config.Commit.Type == "t"
}

// Save save new setting to config file
func (user UserConfig) Save(value reflect.Value) (err error) {
	location := GetAppLocation().UserLocation
	var out []byte
	out, err = yaml.Marshal(user)
	if err != nil {
		return
	}

	if _, err = os.Stat(location); os.IsNotExist(err) {
		return cli.NewExitError("User config not exist, Add to "+location, 2)
	}

	return ioutil.WriteFile(location, out, os.ModePerm)
}

// GetConfigReflectByKey get reflect Kind, Value from keys of yaml file.
//
// Format: separate by '.'
//
// Example: config
//          user.commit.type
func (user UserConfig) GetConfigReflectByKey(keys string) (kind reflect.Kind, val reflect.Value) {
	var key string
	var arr []string
	var valueOf reflect.Value
	var result interface{}

	arr = strings.Split(keys, ".")
	kind = reflect.Struct
	valueOf = reflect.ValueOf(user)

	for _, e := range arr {
		if kind == reflect.Struct {
			key = e
			// fmt.Println("key: ", key)
			val = reflect.Indirect(valueOf).FieldByName(strings.Title(key))
			kind = val.Kind()
			if !val.IsValid() {
				return
			}
			result = val.Interface()
			valueOf = reflect.ValueOf(result)
		} else {
			break
		}
	}
	return
}

// SetValue set user config by key and value inside yaml file
func (user UserConfig) SetValue(key string, value string) (err error) {
	reflectKind, reflectValue := user.GetConfigReflectByKey(key)

	if reflectValue.IsValid() {
		if reflectValue.CanSet() {
			if reflectKind == reflect.Int {
				var res int64
				res, err = strconv.ParseInt(value, 10, 64)
				if err != nil {
					return
				}
				reflectValue.SetInt(res)
			} else if reflectKind == reflect.String {
				reflectValue.SetString(value)
			} else if reflectKind == reflect.Bool {
				var res bool
				res, err = strconv.ParseBool(value)
				if err != nil {
					return
				}
				reflectValue.SetBool(res)
			} else {
				return errors.New("Value is not one of accept type: string, int, bool")
			}
		} else {
			return errors.New("Value cannot set")
		}
	} else {
		return errors.New("Invalid value")
	}

	return user.Save(reflectValue)
}

// InputType input type for key, title, and message of commit message
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
	// fmt.Println("setup app config")
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

func setupLocationConfig(dev bool) (location LocationConfig, err error) {
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

	location = LocationConfig{
		DevConfigLocation:  _setupPath("", ""),
		ProdConfigLocation: _setupPath(defaultLocation, ""),
		AppLocation:        _setupPath(defaultLocation, appfile),
		UserLocation:       _setupPath(defaultLocation, userfile),
		CommitDBLocation:   _setupPath(defaultLocation, commitdbfile),
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

// InstallGitgo will exec install gitgo command
func installGitgo(location LocationConfig) (err error) {
	devLocation := filepath.Clean(location.DevConfigLocation)
	prodLocation := filepath.Clean(location.ProdConfigLocation)
	// already created
	_, err = os.Stat(prodLocation)
	if err == nil {
		return
	}
	_, err = os.Stat(devLocation)
	if err != nil {
		return
		// return cli.NewExitError(devLocation+" cannot be found", 8)
	}
	parent := filepath.Clean(prodLocation + "..")
	_, err = os.Stat(parent)
	if os.IsNotExist(err) {
		err = os.MkdirAll(parent, os.ModePerm)
		if err != nil {
			return
		}
	}
	err = os.Symlink(devLocation, prodLocation)
	return
}

func setupConfig(dev bool) Config {
	location, err := setupLocationConfig(dev)
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	err = installGitgo(location)
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

// Setup setup configuration, must called at the top of main method
func Setup(dev bool) {
	configFile = setupConfig(dev)
}

// GetAppConfig get application config
func GetAppConfig() AppConfig {
	return configFile.AppConfig
}

// GetUserConfig get user config
func GetUserConfig() UserConfig {
	return configFile.UserConfig
}

// GetCommitDBConfig get commit database
func GetCommitDBConfig() CommitConfig {
	return configFile.CommitConfig
}

// GetAppLocation get location of configuration files
func GetAppLocation() LocationConfig {
	return configFile.Location
}
