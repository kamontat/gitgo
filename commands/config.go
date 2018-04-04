package command

import (
	"gitgo/client"
	flag "gitgo/flags"
	"gitgo/models"

	"github.com/urfave/cli"
)

func openConfigFile() {
	defaultEditor := models.GetUserConfig().Config.Editor
	client.OpenFile(defaultEditor, models.GetAppLocation().UserLocation)
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
		Action: func(c *cli.Context) error {
			if !flag.IsKeyExist() && !flag.IsValueExist() {
				openConfigFile()
			}
			return nil
		},
	}
}
