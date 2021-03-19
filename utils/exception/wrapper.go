package exception

import (
	"os"

	"github.com/kamontat/gitgo/utils/logger"
)

// Wrapper enable more function to manage error
type Wrapper struct {
	err error
}

// Exist return true if error exist
func (e *Wrapper) Exist() bool {
	return e.err != nil
}

// Empty return true if error is empty
func (e *Wrapper) Empty() bool {
	return e.err == nil
}

// OnError will run input function will error is exist
func (e *Wrapper) OnError(fn func(error)) *Wrapper {
	if e.Exist() {
		fn(e.err)
	}
	return e
}

// OnCompleted will run input function will error is empty
func (e *Wrapper) OnCompleted(fn func()) *Wrapper {
	if e.Empty() {
		fn()
	}
	return e
}

// Panic will print panic message to console
func (e *Wrapper) Panic() *Wrapper {
	return e.OnError(func(err error) {
		panic(err)
	})
}

// LogExit will log data first, then exit program
func (e *Wrapper) LogExit(log *logger.Logger, code int) {
	e.OnError(func(err error) {
		log.Error(code, err)
	}).Exit(code)
}

// Exit will exit if error exist
func (e *Wrapper) Exit(code int) {
	e.OnError(func(err error) {
		os.Exit(code)
	})
}
