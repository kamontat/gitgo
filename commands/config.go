package command

import (
	"encoding/json"
	"fmt"
	"gitgo/client"
	flag "gitgo/flags"
	"gitgo/models"
	"reflect"
	"strings"

	"github.com/urfave/cli"
)

func openConfigFile() {
	defaultEditor := models.GetUserConfig().Config.Editor
	client.OpenFile(defaultEditor, models.GetAppLocation().UserLocation)
}

func getConfigByKey(keys string) (res string, err error) {
	var returnable string
	var v reflect.Value
	var result interface{}
	var exist bool
	var key string

	arr := strings.Split(keys, ".")
	kind := reflect.Struct
	val := reflect.ValueOf(models.GetUserConfig())

	for _, e := range arr {
		if kind == reflect.Struct {
			key = e
			// fmt.Println("key: ", key)
			v = reflect.Indirect(val).FieldByName(strings.Title(key))
			result = v.Interface()
			val = reflect.ValueOf(result)
			kind = v.Kind()
		} else {
			break
		}
	}

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
		json, _ := json.MarshalIndent(result, "", "  ")
		res = string(json)
		return
	}

	if !exist {
		err = cli.NewExitError("Config value not exist", 5) // errors.New("Config value not exist")
		return
	}

	res = fmt.Sprintf("%s: %s\n", keys, returnable)
	return
}

func setConfigValue(key string, value string) (result string, err error) {
	result = "completed"
	err = models.GetUserConfig().SetValue(key, value)
	return
}

func writeTo(out []byte, filename string) (err error) {
	return nil
}

// AddConfig add command of setting(s)
func AddConfig() cli.Command {
	return cli.Command{
		Name:      "configuration",
		Aliases:   []string{"config", "g"},
		Category:  "Setting",
		Usage:     "Get config commands",
		UsageText: "gitgo config|g ",
		Flags: []cli.Flag{
			flag.AddValueFlag("configuration"),
			flag.AddKeyFlag("configuration"),
		},
		Subcommands: []cli.Command{
			AddConfigLocation(),
		},
		Action: func(c *cli.Context) (err error) {
			var res string
			if !flag.IsKeyExist() && !flag.IsValueExist() {
				openConfigFile()
				return nil
			} else if flag.IsKeyExist() && !flag.IsValueExist() {
				res, err = getConfigByKey(flag.GetKey())
				if err != nil {
					return
				}
			} else {
				res, err = setConfigValue(flag.GetKey(), flag.GetValue())
				if err != nil {
					return
				}
			}
			fmt.Println(res)
			return nil
		},
	}
}
