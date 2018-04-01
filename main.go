package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/urfave/cli"

	"gitgo/flags"
	"gitgo/git"
)

type appConfig struct {
	Name        string
	Description string
	Versions    []versionConfig
	Since       string
	Authors     []cli.Author
	License     string
}

func (appConfig appConfig) latestVersion() versionConfig {
	return appConfig.Versions[0]
}

type versionConfig struct {
	Version     string
	Description string
}

func main() {
	file, e := ioutil.ReadFile("./config/app.json")
	if e != nil {
		fmt.Printf("File error: %v\n", e)
		os.Exit(1)
	}

	var appConfig appConfig
	json.Unmarshal(file, &appConfig)

	app := cli.NewApp()
	app.Name = appConfig.Name
	app.HelpName = appConfig.Name
	app.Usage = appConfig.Description
	app.Version = appConfig.latestVersion().Version
	app.Authors = appConfig.Authors
	app.Copyright = appConfig.License

	// app.UsageText = "gitgo [global options] [command] [command options] [subcommand] [subcommand options] [arguments...]"

	app.Compiled = time.Now()

	// app.Action = func(ctx *cli.Context) error {
	// 	return cli.NewExitError("it is not in the soup", 86)
	// }

	// Flag !
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "lang, l",
			Value: "english",
			Usage: "Language for the greeting",
		},
		cli.StringFlag{
			Name:  "config, c",
			Usage: "Load configuration from `FILE`",
		},
	}

	global.AddForceAsGlobalFlag(app)
	// flags.AddForceAsGlobalFlag(app)

	// Command !
	app.Commands = []cli.Command{
		{
			Name:    "init",
			Aliases: []string{"i"},
			Usage:   "Inital git",
			Action: func(c *cli.Context) error {
				if git.IsNotInit() || global.IsForce() {
					git.Init()
				} else {
					fmt.Println("Initial already!", "Add --force")
				}
				fmt.Println(global.IsForce())
				return nil
			},
		},
		{
			Name:    "add",
			Aliases: []string{"a"},
			Usage:   "Add every file and folder to git",
			Action: func(c *cli.Context) error {
				git.AddAll()
				// fmt.Println(global.IsForce())
				return nil
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
