package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"
	"strconv"
	"strings"

	"github.com/urfave/cli"
	yaml "gopkg.in/yaml.v2"
)

// ##########################################
// ## Object Representation                ##
// ##########################################

// Location is the location of config files
type Location struct {
	Dir        string
	App        string
	User       string
	CommitList string
}

// LocationConfig location of every settings of the application
type LocationConfig struct {
	Dev  Location
	Prod Location
}

// AppConfig will represent application config file (app.yaml)
type AppConfig struct {
	Name        string
	Description string
	Versions    []Version
	Since       string
	Authors     []cli.Author
	License     string
}

// Version is object represent version in app config (app.yaml)
type Version struct {
	Tag         string
	Description string
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

// InputType of commit input
type InputType struct {
	Require bool
	Auto    bool
	Size    int
}

// Commit is representation of commit structure of commit list file (commit_list.yaml)
type Commit struct {
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

// ConfigurationFile is the config in application
type ConfigurationFile struct {
	Location   LocationConfig
	App        AppConfig
	User       UserConfig
	CommitList []Commit
}

// ##########################################
// ## Raw method                           ##
// ##########################################

func _getDevConfigLocation() (str string, err error) {
	goPath := os.Getenv("GOPATH")
	if goPath == "" {
		return "", cli.NewExitError("$GOPATH must be set", 2)
	}
	return fmt.Sprintf("%s/src/github.com/kamontat/gitgo/config", goPath), nil
}

func _getProdConfigLocation() (str string, err error) {
	home := os.Getenv("HOME")
	if home == "" {
		return "", cli.NewExitError("$HOME must be set", 2)
	}
	return fmt.Sprintf("%s/.config/github.com/kamontat/gitgo/config", home), nil
}

func _appendPath(location string, filename string) (result string, err error) {
	// set default location as dev location
	if location == "" {
		return "", cli.NewExitError("input location not exist", 99)
	}
	// get append path
	result = filepath.Clean(location + "/" + filename)
	return
}

func _setupLocation(home string) (location Location, err error) {
	var appfile = "app.yaml"
	var userfile = "user.yaml"
	var commitlistfile = "commit_list.yaml"

	var a, u, c string

	a, err = _appendPath(home, appfile)
	if err != nil {
		return
	}

	u, err = _appendPath(home, userfile)
	if err != nil {
		return
	}

	c, err = _appendPath(home, commitlistfile)
	if err != nil {
		return
	}

	return Location{
		Dir:        home,
		App:        a,
		User:       u,
		CommitList: c,
	}, nil
}

// ##########################################
// ## Helper Method                        ##
// ##########################################

// ------------------------------------------
// || Setup Configuration                  ||
// ------------------------------------------

func setupLocationConfig() (location LocationConfig, err error) {
	// set dev location
	lo, err := _getDevConfigLocation()
	if err != nil {
		return
	}

	devLoc, err := _setupLocation(filepath.Clean(lo))
	if err != nil {
		return
	}

	// set prod location
	lo, err = _getProdConfigLocation()
	if err != nil {
		return
	}

	prodLoc, err := _setupLocation(filepath.Clean(lo))
	if err != nil {
		return
	}

	location = LocationConfig{
		Dev:  devLoc,
		Prod: prodLoc,
	}
	return
}

func setupAppConfig(location string) (app AppConfig, err error) {
	// fmt.Println("setup app config")
	file, err := ioutil.ReadFile(location)
	if err != nil {
		return
	}

	yaml.Unmarshal(file, &app)
	return
}

func setupUserConfig(location string) (user UserConfig, err error) {
	file, err := ioutil.ReadFile(location)
	if err != nil {
		return
	}

	yaml.Unmarshal(file, &user)
	return
}

func setupCommitConfig(location string) (commit []Commit, err error) {
	file, err := ioutil.ReadFile(location)
	if err != nil {
		return
	}

	yaml.Unmarshal(file, &commit)
	return
}

func setupConfigFile() (config ConfigurationFile, err error) {
	location, err := setupLocationConfig()
	if err != nil {
		return
	}

	err = installGitgo(location)
	if err != nil {
		return
	}

	app, err := setupAppConfig(location.Prod.App)
	if err != nil {
		return
	}

	user, err := setupUserConfig(location.Prod.User)
	if err != nil {
		return
	}

	commit, err := setupCommitConfig(location.Prod.CommitList)
	if err != nil {
		return
	}

	return ConfigurationFile{
		Location:   location,
		App:        app,
		User:       user,
		CommitList: commit,
	}, nil
}

// ------------------------------------------
// || User Configuration                   ||
// ------------------------------------------

// save will save input to config file (user.yaml)
func (user UserConfig) save(value reflect.Value) (err error) {
	location := GetAppLocation().Prod.User
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

// GetReflectByKey get reflect Kind, Value from keys of yaml file.
//
// Format: separate by '.'
//
// Example: config
//          user.commit.type
func (user UserConfig) getReflectByKey(keys string) (kind reflect.Kind, val reflect.Value) {
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

func (user UserConfig) setValue(key string, value string) (err error) {
	reflectKind, reflectValue := user.getReflectByKey(key)

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

	return user.save(reflectValue)
}

func (user UserConfig) getValue(key string) (result string, err error) {
	var returnable string
	var exist bool
	kind, v := user.getReflectByKey(key)

	if kind == reflect.String {
		returnable = v.String()
		exist = true
	} else if kind == reflect.Int {
		r := v.Int()
		returnable = fmt.Sprintf("%v", r)
		exist = true
	} else if kind == reflect.Bool {
		r := v.Bool()
		returnable = fmt.Sprintf("%v", r)
		exist = true
	}

	if !exist && kind == reflect.Struct {
		var res []byte
		res, err = json.MarshalIndent(v.Interface(), "", "  ")
		if err != nil {
			return
		}

		returnable = string(res)
		exist = true
	}

	if !exist {
		err = cli.NewExitError("Config value not exist", 5)
		return
	}

	return fmt.Sprintf("%s: %s\n", key, returnable), nil
}

// ------------------------------------------
// || Version                              ||
// ------------------------------------------

// SSPrint will return shortest version output format
func (v Version) SSPrint() string {
	return fmt.Sprintf("version %s", v.Tag)
}

// LSPrint will return fully version output format
func (v Version) LSPrint() string {
	return fmt.Sprintf("version %s: %s", v.Tag, v.Description)
}

// ------------------------------------------
// || Top Configuration                    ||
// ------------------------------------------

func (config ConfigurationFile) _filterCommit(filter func(Commit) bool) []Commit {
	vsf := make([]Commit, 0)
	for _, v := range config.CommitList {
		if filter(v) {
			vsf = append(vsf, v)
		}
	}
	return vsf
}

func (config ConfigurationFile) receiveOneCommit(key string, filterFunc func(Commit) bool) (commit Commit, err error) {
	if key == "" {
		err = errors.New("Input key is empty")
		return
	}

	results := config._filterCommit(filterFunc)

	if len(results) == 0 {
		err = errors.New(key + " key not exist!, (get by 'Commit.Name')")
		return
	}
	commit = results[0]
	return
}

// ------------------------------------------
// || Gitgo Private helper command         ||
// ------------------------------------------

func installGitgo(location LocationConfig) (err error) {
	devLocation := location.Dev.Dir
	prodLocation := location.Prod.Dir
	_, err = os.Stat(devLocation)
	if err != nil {
		return cli.NewExitError("gitgo config not found", 4)
	}
	// already install
	_, err = os.Stat(prodLocation)
	if err == nil {
		return
	}
	// create production folder
	err = os.MkdirAll(prodLocation, os.ModePerm)
	if err != nil {
		return
	}
	// link files
	err = os.Symlink(devLocation, prodLocation)
	return
}

// ##########################################
// ## Public Method                        ##
// ##########################################

// Setup setup configuration, must called at the top of main method
func Setup() (err error) {
	configFile, err = setupConfigFile()
	if err != nil {
		return err
	}
	return nil
}

// GetConfigHelper will get config helper method
func GetConfigHelper() ConfigurationFile {
	return configFile
}

// GetAppConfig get application config
func GetAppConfig() AppConfig {
	return configFile.App
}

// GetUserConfig get user config
func GetUserConfig() UserConfig {
	return configFile.User
}

// GetCommitListConfig get list of commit
func GetCommitListConfig() []Commit {
	return configFile.CommitList
}

// GetAppLocation will return location config struct
func GetAppLocation() LocationConfig {
	return configFile.Location
}

// CommitAsStringArray will convert []Commit to []string
func (config ConfigurationFile) CommitAsStringArray() []string {
	var arr []string
	for _, c := range config.CommitList {
		arr = append(arr, c.Name)
	}
	return arr
}

// GetCommitByName will filter by 'Commit.Name', ignore-case
func (config ConfigurationFile) GetCommitByName(key string) (result Commit, err error) {
	return config.receiveOneCommit(key, func(input Commit) bool {
		return strings.Contains(strings.ToLower(input.Name), strings.ToLower(key))
	})
}

// GetCommitByTitle will filter by 'Commit.Title', ignore-case
func (config ConfigurationFile) GetCommitByTitle(title string) (result Commit, err error) {
	return config.receiveOneCommit(title, func(input Commit) bool {
		return strings.Contains(strings.ToLower(input.Title), strings.ToLower(title))
	})
}

// LatestVersion will return latest application version
func (app AppConfig) LatestVersion() Version {
	return app.Versions[0]
}

// GetVersionShort is getting version as short format
// index 0 is latest version
func (app AppConfig) GetVersionShort(index int) string {
	return fmt.Sprintf("%s %s", app.Name, app.Versions[index].SSPrint())
}

// PrintAllVersionShort will print short format of every version
func (app AppConfig) PrintAllVersionShort() {
	for i := range app.Versions {
		fmt.Println(app.GetVersionShort(i))
	}
}

// GetVersionLong is getting version as long format
// index 0 is latest version
func (app AppConfig) GetVersionLong(index int) string {
	return fmt.Sprintf("%s %s", app.Name, app.Versions[index].LSPrint())
}

// PrintAllVersionLong will print long format of every version
func (app AppConfig) PrintAllVersionLong() {
	for i := range app.Versions {
		fmt.Println(app.GetVersionLong(i))
	}
}

// IsEmojiCommit check is user choose 'emoji' commit
func (user UserConfig) IsEmojiCommit() bool {
	return user.Config.Commit.Type == "emoji" ||
		user.Config.Commit.Type == "moji" ||
		user.Config.Commit.Type == "e"
}

// IsTextCommit check is user choose 'text' commit
func (user UserConfig) IsTextCommit() bool {
	return user.Config.Commit.Type == "text" ||
		user.Config.Commit.Type == "t"
}

// SetValue set user configuration value to config file (user.yaml)
func (user UserConfig) SetValue(key string, value string) error {
	return user.setValue(key, value)
}

// GetValue get user configuration value from config file (user.yaml)
func (user UserConfig) GetValue(key string) (string, error) {
	return user.getValue(key)
}

// ##########################################
// ## Application logic                    ##
// ##########################################

var configFile ConfigurationFile
