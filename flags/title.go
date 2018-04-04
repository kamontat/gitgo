package flag

import (
	"github.com/urfave/cli"
)

var title string

func GetTitle() string {
	return title
}

func AddTitleFlag(use string) cli.Flag {
	return cli.StringFlag{
		Name:        "title, t",
		Usage:       "add title of " + use,
		Destination: &title,
	}
}
