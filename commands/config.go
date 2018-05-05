package command

import (
	"fmt"

	"github.com/kamontat/gitgo/client"
	flag "github.com/kamontat/gitgo/flags"
	"github.com/kamontat/gitgo/models"

	"github.com/urfave/cli"
)

func openConfigFile() {
	defaultEditor := models.GetUserConfig().Config.Editor
	client.OpenFile(defaultEditor, models.GetAppLocation().Prod.User)
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
				res, err = models.GetUserConfig().GetValue(flag.GetKey())
				if err != nil {
					return
				}
				fmt.Println(res)
			} else {
				err = models.GetUserConfig().SetValue(flag.GetKey(), flag.GetValue())
				if err != nil {
					return
				}
			}
			fmt.Println("completed")
			return nil
		},
	}
}
