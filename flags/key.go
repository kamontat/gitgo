package flag

import (
	"github.com/urfave/cli"
)

var key string

func GetKey() string {
	return key
}

func AddKeyFlag(use string) cli.Flag {
	return cli.StringFlag{
		Name:        "key, k",
		Usage:       "add key of " + use,
		Destination: &key,
	}
}
