package command

import (
	"fmt"

	flag "github.com/kamontat/gitgo/flags"
	"github.com/kamontat/gitgo/models"

	"github.com/urfave/cli"
)

func AddConfigLocation() cli.Command {
	return cli.Command{
		Name:    "location",
		Aliases: []string{"l"},
		Flags: []cli.Flag{
			flag.AllFlag(),
		},
		Action: func(c *cli.Context) error {
			if flag.IsAll() {
				fmt.Printf("App configuration:  %s\n", models.GetAppLocation().AppLocation)
				fmt.Printf("User configuration: %s\n", models.GetAppLocation().UserLocation)
				fmt.Printf("Commit database:    %s\n", models.GetAppLocation().CommitDBLocation)
			} else {
				fmt.Printf("User configuration: %s\n", models.GetAppLocation().UserLocation)
			}
			return nil
		},
	}
}
