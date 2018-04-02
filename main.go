package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sort"

	"github.com/urfave/cli"

	"gitgo/commands"
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

func (appConfig appConfig) printEveryVersions() {
	for _, v := range appConfig.Versions {
		v.printVersion(appConfig.Name)
	}
}

func (appConfig appConfig) printFullEveryVersions() {
	for _, v := range appConfig.Versions {
		v.printFullVersion(appConfig.Name)
	}
}

func (v versionConfig) printVersion(name string) {
	fmt.Printf("%s version %s\n", name, v.Version)
}

func (v versionConfig) printFullVersion(name string) {
	fmt.Printf("%s version %s: %s\n", name, v.Version, v.Description)
}

func addVersion(appConfig appConfig) cli.Command {
	var full bool
	return cli.Command{
		Name:    "version",
		Aliases: []string{"v"},
		Usage:   "show version, same as --version",
		Flags: []cli.Flag{
			cli.BoolFlag{
				Name:        "full, F",
				Usage:       "show full output",
				Destination: &full,
			},
		},
		Action: func(c *cli.Context) error {
			if full {
				appConfig.latestVersion().printFullVersion(appConfig.Name)
			} else {
				appConfig.latestVersion().printVersion(appConfig.Name)
			}
			return nil
		},
	}
}

func addListVersion(appConfig appConfig) cli.Command {
	var full bool
	return cli.Command{
		Name:    "list-version",
		Aliases: []string{"L"},
		Usage:   "list every version, same as --list-version",
		Flags: []cli.Flag{
			cli.BoolFlag{
				Name:        "full, F",
				Usage:       "show full output",
				Destination: &full,
			},
		},
		Action: func(c *cli.Context) error {
			if full {
				appConfig.printFullEveryVersions()
			} else {
				appConfig.printEveryVersions()
			}
			return nil
		},
	}
}

func main() {
	// lv - list-version
	var full, lv bool
	if os.Getenv("GOPATH") == "" {
		cli.HandleExitCoder(cli.NewExitError("$GOPATH must be set", 2))
	}

	file, e := ioutil.ReadFile(os.Getenv("GOPATH") + "/src/gitgo/config/app.json")
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

	app.EnableBashCompletion = true

	// app.UsageText = "gitgo [global options] [command] [command options] [subcommand] [subcommand options] [arguments...]"

	app.Commands = []cli.Command{
		command.InitGit(), command.AddGit(), command.DestroyGit(),
		command.PushGit(), command.PullGit(),
		addVersion(appConfig), addListVersion(appConfig),
	}

	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:        "full, F",
			Usage:       "show full output",
			Destination: &full,
		}, cli.BoolFlag{
			Name:        "list-version, L",
			Usage:       "list all version",
			Destination: &lv,
		},
	}

	cli.VersionPrinter = func(c *cli.Context) {
		if full {
			appConfig.latestVersion().printFullVersion(appConfig.Name)
		} else {
			appConfig.latestVersion().printVersion(appConfig.Name)
		}
	}

	app.Action = func(c *cli.Context) error {
		if lv {
			if full {
				appConfig.printFullEveryVersions()
			} else {
				appConfig.printEveryVersions()
			}
		} else {
			cli.ShowAppHelp(c)
		}
		return nil
	}

	sort.Sort(cli.FlagsByName(app.Flags))
	// sort.Sort(cli.CommandsByName(app.Commands))

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
