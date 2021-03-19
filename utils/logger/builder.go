package logger

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
	switch level {
	case "debug":
		logger.SetLevel(DEBUG)
	case "info":
		logger.SetLevel(INFO)
	case "warn":
		logger.SetLevel(WARN)
	case "error":
		logger.SetLevel(ERROR)
	case "silent":
		logger.SetLevel(SILENT)
	case "":
		// do mothing
	default:
		logger.Warn(0, "input '%s' is not accepted level", level)
	}
}
