package main

import (
	"log"
	"os"
	"sort"

	"github.com/urfave/cli"

	"gitgo/commands"
	"gitgo/models"
)

func main() {
	// lv - list-version
	var full, lv bool

	models.Setup(true)

	// config := models.SetupConfig(true)
	appConfig := models.GetAppConfig()

	// fmt.Println(models.GetUserConfig())

	app := cli.NewApp()
	app.Name = appConfig.Name
	app.HelpName = appConfig.Name
	app.Usage = appConfig.Description
	app.Version = appConfig.LatestVersion().Version
	app.Authors = appConfig.Authors
	app.Copyright = appConfig.License

	app.EnableBashCompletion = true

	// app.UsageText = "gitgo [global options] [command] [command options] [subcommand] [subcommand options] [arguments...]"

	app.Commands = []cli.Command{
		command.InitGit(), command.DestroyGit(),
		command.PushGit(), command.PullGit(),

		command.AddGitStatus(), command.AddGit(), command.CommitGit(),
		command.AddConfig(), command.AddVersion(appConfig), command.AddListVersion(appConfig),
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
			appConfig.LatestVersion().PrintFullVersion(appConfig.Name)
		} else {
			appConfig.LatestVersion().PrintVersion(appConfig.Name)
		}
	}

	app.Action = func(c *cli.Context) error {
		if lv {
			if full {
				appConfig.PrintFullEveryVersions()
			} else {
				appConfig.PrintEveryVersions()
			}
		} else {
			cli.ShowAppHelp(c)
		}
		return nil
	}

	sort.Sort(cli.FlagsByName(app.Flags))
	// sort.Sort(cli.CommandsByName(app.Commands))

	runError := app.Run(os.Args)
	if runError != nil {
		log.Fatal(runError)
	}
}
