package phase

import (
	"github.com/kamontat/gitgo/utils/exception"
	"github.com/kamontat/gitgo/utils/logger"
)

var log *logger.Logger = logger.Get()
var currentPhase *phase

func setPhase(phaseID *phase) {
	currentPhase = phaseID
}

// Debug will log data as debug
func Debug(format string, params ...interface{}) {
	log.Debug(currentPhase.ID, format, params...)
}

// Info will log data as infomation
func Info(format string, params ...interface{}) {
	log.Info(currentPhase.ID, format, params...)
}

// Warn will log error as warning message
func Warn(err error) {
	exception.When(err).OnError(func(err error) {
		log.Warn(currentPhase.ID, "%s", err.Error())
	})
}

// Error will throw error if exist with current phase
func Error(err error) {
	exception.When(err).LogExit(logger.Get(), currentPhase.ID)
}

// Log will print data without formatting anything
func Log(message string) {
	log.Log(message)
}

func Format(format string, params ...interface{}) {
	log.Format(format, params...)
}
