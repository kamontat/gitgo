package logger

import "fmt"

// Logger is logging object
type Logger struct {
	level Level
}

// SetLevel will update current logging level
func (l *Logger) SetLevel(level Level) {
	l.level = level
}

// IsDebug will return true if on debug mode
func (l *Logger) IsDebug() bool {
	return l.CheckLevel(DEBUG)
}

// CheckLevel is input level should be log or not
func (l *Logger) CheckLevel(level Level) bool {
	if level.Code <= l.level.Code {
		return true
	}

	return false
}

func (l *Logger) private(level Level, key int, format string, params ...interface{}) {
	if l.CheckLevel(level) {
		fullFormat := "%s [%03d] " + format

		newParams := make([]interface{}, 0)
		newParams = append(newParams, level.Short)
		newParams = append(newParams, key%1000)
		for _, p := range params {
			newParams = append(newParams, p)
		}

		l.Format(fullFormat, newParams...)
	}
}

// Newline will go to newline
func (l *Logger) Newline() {
	l.Log("")
}

// Debug is logging as debug message
func (l *Logger) Debug(key int, format string, params ...interface{}) {
	l.private(DEBUG, key, format, params...)
}

// Info is logging as info message
func (l *Logger) Info(key int, format string, params ...interface{}) {
	l.private(INFO, key, format, params...)
}

// Warn is logging as warn message
func (l *Logger) Warn(key int, format string, params ...interface{}) {
	l.private(WARN, key, format, params...)
}

// Error is logging as error message
func (l *Logger) Error(key int, err error) {
	l.private(ERROR, key, "%s", err.Error())
}

// Format will do formatting message before log
func (l *Logger) Format(format string, params ...interface{}) {
	l.Log(fmt.Sprintf(format, params...))
}

// Log will logging data without any formatting
func (l *Logger) Log(msg string) {
	if l.level != SILENT {
		fmt.Println(msg)
	}
}
