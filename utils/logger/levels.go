package logger

import (
	"strings"
)

// Level is logging level
type Level struct {
	Code  uint8
	Name  string
	Short string
}

func (l Level) String() string {
	return strings.ToLower(l.Name)
}

var (
	// DEBUG is lowest level in logging
	DEBUG Level = Level{
		Code:  uint8(5),
		Name:  "Debug",
		Short: "DBG",
	}

	// INFO is common message
	INFO Level = Level{
		Code:  uint8(3),
		Name:  "Info",
		Short: "INF",
	}

	// WARN is a calm error message
	WARN Level = Level{
		Code:  uint8(2),
		Name:  "Warn",
		Short: "WRN",
	}

	// ERROR is a critical error message
	ERROR Level = Level{
		Code:  uint8(1),
		Name:  "Error",
		Short: "ERR",
	}

	// SILENT is a special level for mute all logging
	SILENT Level = Level{
		Code:  uint8(0),
		Name:  "Silent",
		Short: "",
	}
)
