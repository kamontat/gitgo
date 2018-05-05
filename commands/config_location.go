package command

import (
	"fmt"

	flag "github.com/kamontat/gitgo/flags"
	"github.com/kamontat/gitgo/models"

	"github.com/urfave/cli"
)

// AddConfigLocation add cli command for getting location of config file
func AddConfigLocation() cli.Command {
	return cli.Command{
		Name:    "location",
		Aliases: []string{"l"},
		Flags: []cli.Flag{
			flag.AllFlag(),
		},
		Action: func(c *cli.Context) error {
			dev := models.GetAppLocation().Dev
			prod := models.GetAppLocation().Prod
			if flag.IsAll() {
				fmt.Printf("(Dev)  App    configuration:  %s\n", dev.App)
				fmt.Printf("(Prod) App    configuration:  %s\n", prod.App)
				fmt.Printf("(Dev)  User   configuration:  %s\n", dev.User)
				fmt.Printf("(Prod) User   configuration:  %s\n", prod.User)
				fmt.Printf("(Dev)  Commit configuration:  %s\n", dev.CommitList)
				fmt.Printf("(Prod) Commit configuration:  %s\n", prod.CommitList)
			} else {
				fmt.Printf("(Dev)  User configuration:  %s\n", dev.User)
				fmt.Printf("(Prod) User configuration:  %s\n", prod.User)
			}
			return nil
		},
	}
}
