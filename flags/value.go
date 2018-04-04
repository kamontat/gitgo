package flag

import (
	"github.com/urfave/cli"
)

var value string

func GetValue() string {
	return value
}

func IsValueExist() bool {
	return value != ""
}

func AddValueFlag(use string) cli.Flag {
	return cli.StringFlag{
		Name:        "value, val, v",
		Usage:       "set value of " + use,
		Destination: &key,
	}
}
