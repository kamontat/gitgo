package flag

import (
	"github.com/urfave/cli"
)

var key string

func GetKey() string {
	return key
}

func IsKeyExist() bool {
	return key != ""
}

func AddKeyFlag(use string) cli.Flag {
	return cli.StringFlag{
		Name:        "key, k",
		Usage:       "set key of " + use,
		Destination: &key,
	}
}
