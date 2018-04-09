package main

import (
	"log"
	"os"
	"sort"

	"github.com/urfave/cli"

	"github.com/kamontat/gitgo/commands"
	flag "github.com/kamontat/gitgo/flags"
	"github.com/kamontat/gitgo/models"
)

func main() {
	models.Setup(false)
	appConfig := models.GetAppConfig()

	app := cli.NewApp()
	app.Name = appConfig.Name
	app.HelpName = appConfig.Name
	app.Usage = appConfig.Description
	app.Version = appConfig.LatestVersion().Version
	app.Authors = appConfig.Authors
	app.Copyright = appConfig.License

	app.EnableBashCompletion = true

	app.Commands = []cli.Command{
		command.AddVersion(appConfig), command.AddListVersion(appConfig),

		command.AddGit(), command.CommitGit(),

		command.InitGit(), command.DestroyGit(),
		command.AddGitStatus(), command.AddConfig(),

		command.PushGit(), command.PullGit(),
	}

	app.Flags = []cli.Flag{
		flag.FullFlag(),
		flag.ListVersionFlag(),
	}

	cli.VersionPrinter = func(c *cli.Context) {
		appConfig.LatestVersion().ChooseToPrintVersion(flag.IsFull())
	}

	app.Action = func(c *cli.Context) error {
		if flag.NeedToListVersion() {
			appConfig.ChooseToPrintEveryVersions(flag.IsFull())
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
