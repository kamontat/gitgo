package logger

import "strings"

var logger *Logger = &Logger{
	level: INFO,
}

// Get will return singleton logger object
func Get() *Logger {
	return logger
}

// SetLevel is mirror logger set level
func SetLevel(level Level) {
	logger.SetLevel(level)
}

// SetLevelStr will transform string to level and set
func SetLevelStr(level string) {
	switch strings.ToLower(level) {
	case "debug", "5":
		logger.SetLevel(DEBUG)
	case "info", "3":
		logger.SetLevel(INFO)
	case "warn", "2":
		logger.SetLevel(WARN)
	case "error", "1":
		logger.SetLevel(ERROR)
	case "silent", "0":
		logger.SetLevel(SILENT)
	case "":
		// do mothing
	default:
		logger.Warn(0, "input '%s' is not accepted level", level)
	}
}
